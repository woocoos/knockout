package file

import (
	"bytes"
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/test"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	adminToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI2N2E4NzQ4MmU5MWY0ZjJlOTIyMGY1MTM3NjE4NWI3ZSIsInN1YiI6IjEiLCJuYW1lIjoiYWRtaW4iLCJpYXQiOjE1MTYyMzkwMjJ9.MzfOsaGAtHj0IIoVDgpfUD1GMtgLTNbY7D_CkEoR9CQ"
	adminTokenJTI = "67a87482e91f4f2e9220f51376185b7e"
)

type fileSuite struct {
	suite.Suite
	FileServer  *ServerImpl
	redisServer *miniredis.Miniredis

	server *web.Server
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, &fileSuite{})
}

func (t *fileSuite) SetupSuite() {
	t.redisServer = miniredis.RunT(t.T())
	t.Require().NoError(t.redisServer.Set(adminTokenJTI, "1"))

	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	t.Require().NoError(err)

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()
	cnf.Parser().Set("cache.redis.addrs", []string{t.redisServer.Addr()})
	app := koapp.New(woocoo.WithAppConfiguration(cnf))
	appCnf := app.AppConfiguration()
	appCnf.Parser().Set("store.portal", appCnf.Get("store.test"))
	t.FileServer = NewServer(app)

	t.server = t.FileServer.webServer

	// testdata
	sf, err := os.Open(test.Path("testdata/etc/app.yaml"))
	t.Require().NoError(err)
	dir := test.Path("/tmp/1000/tmp")
	if fi, _ := os.Stat(dir); fi != nil {
		os.RemoveAll(dir)
	}
	t.Require().NoError(os.MkdirAll(dir, os.ModePerm))
	defer sf.Close()
	tr, err := os.Create(test.Path("/tmp/1000/tmp/suite.yaml"))
	t.Require().NoError(err)
	_, err = io.Copy(tr, sf)
	defer tr.Close()
	t.Require().NoError(err)

	ctx := context.Background()
	db := t.FileServer.db
	err = db.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	t.NoError(err)

	db.FileSource.Create().SetID(1).SetKind(filesource.KindLocal).SetComments("本地存储bucket").
		SetEndpoint("http://localhost:9000").SetBucket("test-bucket").SetCreatedBy(1).SaveX(ctx)
	db.File.Create().SetID(1).SetTenantID(1000).SetSourceID(1).SetCreatedBy(1).
		SetName("suite.yaml").SetPath("1000/tmp/suite.yaml").SetMineType("text/yaml").SetSize(0).SaveX(ctx)
}

func (t *fileSuite) Test_Upload() {
	t.Run("normal", func() {
		testFile := test.Path("testdata/etc/app.yaml")
		// 使用bufio创建上传文件数据
		payload := &bytes.Buffer{}
		mrw := multipart.NewWriter(payload)
		t.NoError(mrw.WriteField("key", "1000/tmp/app.yaml"))
		t.NoError(mrw.WriteField("bucket", "test-bucket"))

		formFile, _ := mrw.CreateFormFile("file", filepath.Base(testFile))
		fh, err := os.Open(test.Path("testdata/etc/app.yaml"))
		t.Require().NoError(err)
		_, _ = io.Copy(formFile, fh)
		fh.Close()
		_ = mrw.Close()

		req := httptest.NewRequest("POST", "/files", payload)
		req.Header.Set("Authorization", "Bearer "+adminToken)
		req.Header.Set("Content-Type", mrw.FormDataContentType())
		req.Header.Set("X-Tenant-ID", "1000")
		w := httptest.NewRecorder()
		t.server.Router().ServeHTTP(w, req)
		resp := w.Result()

		if !t.Equal(200, resp.StatusCode) {
			print(w.Body.String())
		}
	})
}

func (t *fileSuite) TestGetFileRaw() {
	req := httptest.NewRequest("GET", "/files/1/raw", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1000")
	w := httptest.NewRecorder()
	t.server.Router().ServeHTTP(w, req)
	resp := w.Result()

	t.Equal(200, resp.StatusCode)
	t.Equal("text/yaml", resp.Header.Get("Content-Type"))
}

func (t *fileSuite) TestGetFile() {
	req := httptest.NewRequest("GET", "/files/1", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1000")
	w := httptest.NewRecorder()
	t.server.Router().ServeHTTP(w, req)
	resp := w.Result()

	t.Equal(200, resp.StatusCode)
	bd, err := io.ReadAll(resp.Body)
	t.NoError(err)
	t.Contains(string(bd), "suite.yaml")
}

func (t *fileSuite) Test_Delete() {
	req := httptest.NewRequest("DELETE", "/files/1", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1000")
	w := httptest.NewRecorder()
	t.server.Router().ServeHTTP(w, req)
	resp := w.Result()

	t.Equal(200, resp.StatusCode)
	t.NoFileExists(test.Path("tmp/1000/tmp/suite.yaml"))
}

func TestPathMatch(t *testing.T) {
	c, router := gin.CreateTestContext(httptest.NewRecorder())
	router.GET("/:bucket/*path", func(c *gin.Context) {
		t.Log(c.Param("path"))
		c.JSON(200, "ok")
	})
	r := httptest.NewRequest("GET", "/abc/bcd", nil)
	c.Request = r
	router.ServeHTTP(c.Writer, c.Request)
	assert.Equal(t, 200, c.Writer.Status())
}
