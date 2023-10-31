package file

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout-go/pkg/snowflake"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/filesource"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"
)

var _ FileServer = (*ServerImpl)(nil)

type ServerImpl struct {
	baseDir  string
	endpoint string
	db       *ent.Client

	webServer *web.Server
}

func NewServer(app *woocoo.App) *ServerImpl {
	cnf := app.AppConfiguration()
	si := &ServerImpl{
		baseDir:  cnf.Abs(cnf.String("files.local.baseDir")),
		endpoint: cnf.String("files.local.endpoint"),
	}
	ents := koapp.BuildEntComponents(cnf)
	if cnf.Development {
		si.db = ent.NewClient(ent.Driver(ents["portal"]), ent.Debug())
	} else {
		si.db = ent.NewClient(ent.Driver(ents["portal"]))
	}
	si.buildWebServer(app)

	app.RegisterServer(si.webServer)
	return si
}

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (si *ServerImpl) Start(ctx context.Context) error {
	return nil
}

func (si *ServerImpl) Stop(ctx context.Context) error {
	return si.db.Close()
}

func (si *ServerImpl) buildWebServer(app *woocoo.App) *web.Server {
	si.webServer = web.New(web.WithConfiguration(app.AppConfiguration().Sub("web")),
		web.WithGracefulStop(),
		otelweb.RegisterMiddleware(),
	)
	// default group is '/'
	RegisterFileHandlers(si.webServer.Router().FindGroup("/").Group, si)
	return si.webServer
}

func (si *ServerImpl) UploadFileInfo(c *gin.Context, r *UploadFileInfoRequest) (string, error) {
	tid, err := si.tryGetTenantID(c)
	if err != nil {
		return "", err
	}
	fs := r.FileSource
	fsID, err := si.db.FileSource.Query().Where(
		filesource.KindEQ(filesource.Kind(fs.Kind)),
		filesource.Endpoint(fs.Endpoint),
		filesource.Bucket(fs.Bucket),
		filesource.Region(fs.Region),
	).Select(filesource.FieldID).Int(c)
	if err != nil {
		return "", fmt.Errorf("invalid filesource")
	}
	fileInput := r.File
	fi, err := si.db.File.Create().SetTenantID(tid).SetName(fileInput.Name).SetSourceID(fsID).
		SetPath(fileInput.Path).SetSize(fileInput.Size).SetMineType(fileInput.MineType).
		Save(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(fi.ID), nil
}

func (si *ServerImpl) GetFileRaw(c *gin.Context, r *GetFileRawRequest) ([]byte, error) {
	tid, err := si.tryGetTenantID(c)
	if err != nil {
		return nil, err
	}
	fid, err := strconv.Atoi(r.FileId)
	if err != nil {
		return nil, err
	}

	fi, err := si.db.File.Query().Where(file.TenantID(tid), file.ID(fid)).First(c)
	if err != nil {
		return nil, err
	}
	c.Header("Content-Type", fi.MineType)

	fn, err := si.getStorePath(c, fi.Path)
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

func (si *ServerImpl) DeleteFile(c *gin.Context, r *DeleteFileRequest) error {
	tid, err := si.tryGetTenantID(c)
	if err != nil {
		return err
	}
	fid, err := strconv.Atoi(r.FileId)
	if err != nil {
		return err
	}

	fi, err := si.db.File.Query().Where(file.TenantID(tid), file.ID(fid)).First(c)
	if err != nil {
		return err
	}
	fn, err := si.getStorePath(c, fi.Path)
	if err != nil {
		return err
	}
	if err := os.Remove(fn); err != nil {
		return err
	}
	return si.db.File.DeleteOne(fi).Exec(c)
}

func (si *ServerImpl) GetFile(c *gin.Context, r *GetFileRequest) (*FileInfo, error) {
	tid, err := si.tryGetTenantID(c)
	if err != nil {
		return nil, err
	}
	fid, err := strconv.Atoi(r.FileId)
	if err != nil {
		return nil, err
	}
	fi, err := si.db.File.Query().Where(file.TenantID(tid), file.ID(fid)).WithSource().First(c)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		ID:        r.FileId,
		Name:      fi.Name,
		Size:      fi.Size,
		Path:      fi.Path,
		CreatedAt: fi.CreatedAt,
		FileSource: &FileSource{
			ID:       fi.Edges.Source.ID,
			Endpoint: fi.Edges.Source.Endpoint,
			Bucket:   fi.Edges.Source.Bucket,
			Region:   fi.Edges.Source.Region,
			Kind:     fi.Edges.Source.Kind.String(),
		},
	}, nil
}

func (si *ServerImpl) tryGetTenantID(c *gin.Context) (tid int, err error) {
	if str := c.GetHeader("X-Tenant-ID"); str != "" {
		if tid, err = strconv.Atoi(str); err != nil {
			return 0, err
		}
	}
	return
}

func (si *ServerImpl) UploadFile(c *gin.Context, r *UploadFileRequest) (fid string, err error) {
	var tid int
	if tid, err = si.tryGetTenantID(c); err != nil {
		return "", err
	}
	// 获取source
	fs, err := si.db.FileSource.Query().Where(
		filesource.KindEQ(filesource.KindLocal),
		filesource.Bucket(r.Bucket),
		filesource.Endpoint(si.endpoint),
	).Only(c)
	if err != nil {
		return "", err
	}
	if fs == nil {
		return "", fmt.Errorf("invalid bucket:%s", r.Bucket)
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	if header.Size == 0 {
		return "", fmt.Errorf("file size is zero")
	}

	refname, err := si.getStorePath(c, r.Key)
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
	fi, err := si.db.File.Create().SetID(int(id)).SetTenantID(tid).SetName(header.Filename).SetSourceID(fs.ID).
		SetPath(r.Key).SetSize(int(size)).SetMd5(md5Sum).SetMineType(mine).
		Save(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(fi.ID), nil
}

func (si *ServerImpl) getStorePath(ctx context.Context, key string) (string, error) {
	return filepath.Join(si.baseDir, key), nil
}

func (si *ServerImpl) ReportRefCount(ctx *gin.Context, r *ReportRefCountRequest) (bool, error) {
	tid, err := si.tryGetTenantID(ctx)
	if err != nil {
		return false, err
	}
	err = clientx.WithTx(ctx, func(ctx context.Context) (clientx.Transactor, error) {
		return si.db.Tx(ctx)
	}, func(itx clientx.Transactor) error {
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
