package server

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/file"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"
)

var _ oas.FileServer = (*FileService)(nil)

type FileService struct {
	baseDir string
	DB      *ent.Client
}

func (f *FileService) GetFileRaw(c *gin.Context, r *oas.GetFileRawRequest) ([]byte, error) {
	tid, err := f.tryGetTenantID(c)
	if err != nil {
		return nil, err
	}
	fid, err := strconv.Atoi(r.UriParams.FileId)
	if err != nil {
		return nil, err
	}

	fi, err := f.DB.File.Query().Where(file.TenantID(tid), file.ID(fid)).First(c)
	if err != nil {
		return nil, err
	}
	c.Header("Content-Type", fi.MineType)

	fn, err := f.getStorePath(c, tid, fi.Path)
	if err != nil {
		return nil, err
	}
	fs, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	_, err = io.Copy(c.Writer, fs)
	return nil, err
}

func (f *FileService) DeleteFile(c *gin.Context, r *oas.DeleteFileRequest) error {
	tid, err := f.tryGetTenantID(c)
	if err != nil {
		return err
	}
	fid, err := strconv.Atoi(r.UriParams.FileId)
	if err != nil {
		return err
	}

	fi, err := f.DB.File.Query().Where(file.TenantID(tid), file.ID(fid)).First(c)
	if err != nil {
		return err
	}
	fn, err := f.getStorePath(c, tid, fi.Path)
	if err != nil {
		return err
	}
	if err := os.Remove(fn); err != nil {
		return err
	}
	return f.DB.File.DeleteOne(fi).Exec(c)
}

func (f *FileService) GetFile(c *gin.Context, r *oas.GetFileRequest) (*oas.FileInfo, error) {
	tid, err := f.tryGetTenantID(c)
	if err != nil {
		return nil, err
	}
	fid, err := strconv.Atoi(r.UriParams.FileId)
	if err != nil {
		return nil, err
	}
	fi, err := f.DB.File.Query().Where(file.TenantID(tid), file.ID(fid)).First(c)
	if err != nil {
		return nil, err
	}
	return &oas.FileInfo{
		ID:        r.UriParams.FileId,
		Name:      fi.Name,
		Size:      fi.Size,
		CreatedAt: fi.CreatedAt,
	}, nil
}

func (f *FileService) tryGetTenantID(c *gin.Context) (tid int, err error) {
	if str := c.GetHeader("X-Tenant-ID"); str != "" {
		if tid, err = strconv.Atoi(str); err != nil {
			return 0, err
		}
	}
	return
}

func (f *FileService) UploadFile(c *gin.Context, r *oas.UploadFileRequest) (fid string, err error) {
	var tid int
	if tid, err = f.tryGetTenantID(c); err != nil {
		return "", err
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	if header.Size == 0 {
		return "", fmt.Errorf("file size is zero")
	}

	refname, err := f.getStorePath(c, tid, r.Body.Key)
	if err != nil {
		return "", err
	}
	fingerprint := md5.New()
	// 上传文件只能是新建,不允许覆盖
	out, err := os.OpenFile(refname, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	mw := io.MultiWriter(out, fingerprint)
	// 落盘
	if err = os.MkdirAll(filepath.Dir(refname), os.ModePerm); err != nil {
		return "", err
	}
	size, err := io.Copy(mw, file)
	if err != nil {
		return "", err
	}
	// 后续如果存在数据库操作错误等,则需要删除文件存储
	defer func() {
		out.Close()

		if perr := recover(); perr != nil {
			if os.Remove(refname) != nil {
				log.Warn(fmt.Sprintf("delete file %s failed", refname))
			}
			panic(perr)
		}

		if err != nil {
			if os.Remove(refname) != nil {
				log.Warn(fmt.Sprintf("delete file %s failed", refname))
			}
		}
	}()

	// 计算上传文件的Md5值
	md5Sum := hex.EncodeToString(fingerprint.Sum(nil))
	// size kb
	size = header.Size >> 10
	id := snowflake.New().Int64()
	mine := mime.TypeByExtension(filepath.Ext(header.Filename))
	fi, err := f.DB.File.Create().SetID(int(id)).SetTenantID(tid).SetName(header.Filename).SetSourceID(0).
		SetPath(r.Body.Key).SetSize(int(size)).SetMd5(md5Sum).SetMineType(mine).
		Save(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(fi.ID), nil
}

func (f *FileService) getStorePath(ctx context.Context, tid int, key string) (string, error) {
	return filepath.Join(f.baseDir, strconv.Itoa(tid), key), nil
}
