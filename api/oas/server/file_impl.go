package server

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/filesource"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"
)

var _ oas.FileServer = (*FileService)(nil)

type FileService struct {
	BaseDir  string
	Endpoint string
	DB       *ent.Client
}

func (f *FileService) UploadFileInfo(c *gin.Context, r *oas.UploadFileInfoRequest) (string, error) {
	tid, err := f.tryGetTenantID(c)
	if err != nil {
		return "", err
	}
	fs := r.Body.FileSource
	fsID, err := f.DB.FileSource.Query().Where(
		filesource.KindEQ(filesource.Kind(fs.Kind)),
		filesource.Endpoint(fs.Endpoint),
		filesource.Bucket(fs.Bucket),
		filesource.Region(fs.Region),
	).Select(filesource.FieldID).Int(c)
	if err != nil {
		return "", fmt.Errorf("invalid filesource")
	}
	fileInput := r.Body.File
	fi, err := f.DB.File.Create().SetTenantID(tid).SetName(fileInput.Name).SetSourceID(fsID).
		SetPath(fileInput.Path).SetSize(fileInput.Size).SetMineType(fileInput.MineType).
		Save(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(fi.ID), nil
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

	fn, err := f.getStorePath(c, fi.Path)
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
	fn, err := f.getStorePath(c, fi.Path)
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
	fi, err := f.DB.File.Query().Where(file.TenantID(tid), file.ID(fid)).WithSource().First(c)
	if err != nil {
		return nil, err
	}
	return &oas.FileInfo{
		ID:        r.UriParams.FileId,
		Name:      fi.Name,
		Size:      fi.Size,
		Path:      fi.Path,
		CreatedAt: fi.CreatedAt,
		FileSource: &oas.FileSource{
			ID:       fi.Edges.Source.ID,
			Endpoint: fi.Edges.Source.Endpoint,
			Bucket:   fi.Edges.Source.Bucket,
			Region:   fi.Edges.Source.Region,
			Kind:     fi.Edges.Source.Kind.String(),
		},
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
	// 获取source
	fs, err := f.DB.FileSource.Query().Where(
		filesource.KindEQ(filesource.KindLocal),
		filesource.Bucket(r.Body.Bucket),
		filesource.Endpoint(f.Endpoint),
	).Only(c)
	if err != nil {
		return "", err
	}
	if fs == nil {
		return "", fmt.Errorf("invalid bucket:%s", r.Body.Bucket)
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	if header.Size == 0 {
		return "", fmt.Errorf("file size is zero")
	}

	refname, err := f.getStorePath(c, r.Body.Key)
	if err != nil {
		return "", err
	}
	fingerprint := md5.New()
	// 落盘
	if err = os.MkdirAll(filepath.Dir(refname), os.ModePerm); err != nil {
		return "", err
	}
	// 上传文件只能是新建,不允许覆盖
	out, err := os.OpenFile(refname, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	mw := io.MultiWriter(out, fingerprint)

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
	fi, err := f.DB.File.Create().SetID(int(id)).SetTenantID(tid).SetName(header.Filename).SetSourceID(fs.ID).
		SetPath(r.Body.Key).SetSize(int(size)).SetMd5(md5Sum).SetMineType(mine).
		Save(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(fi.ID), nil
}

func (f *FileService) getStorePath(ctx context.Context, key string) (string, error) {
	return filepath.Join(f.BaseDir, key), nil
}

func (f *FileService) ReportRefCount(ctx *gin.Context, r *oas.ReportRefCountRequest) (bool, error) {
	tid, err := f.tryGetTenantID(ctx)
	if err != nil {
		return false, err
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return f.DB.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		for _, v := range r.Inputs {
			update := tx.File.UpdateOneID(v.FileId).Where(file.TenantID(tid))
			if v.OpType == "plus" {
				// 加
				update.AddRefCount(1)
			} else if v.OpType == "minus" {
				// 减
				update.AddRefCount(-1)
			} else {
				return fmt.Errorf("invalid opType")
			}
			err = update.Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
