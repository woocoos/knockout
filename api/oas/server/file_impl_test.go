package server

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/woocoos/knockout/test"
	"google.golang.org/grpc/testdata"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type fileServiceSuite struct {
	suite.Suite
	*ServiceSuite
}

// loginFlowSuite runs all the tests in the suite.
func TestFileService(t *testing.T) {
	service := ServiceSuite{}
	service.SetupSuite(t)
	suite.Run(t, &fileServiceSuite{
		ServiceSuite: &service,
	})
}

func (t *fileServiceSuite) SetupSuite() {
	sf, err := os.Open(test.Path("testdata/etc/app.yaml"))
	t.Require().NoError(err)
	t.Require().NoError(os.MkdirAll(test.Path("testdata/tmp/1000/tmp"), os.ModePerm))
	defer sf.Close()
	tr, err := os.Create(test.Path("tmp/1000/tmp/suite.yaml"))
	t.Require().NoError(err)
	_, err = io.Copy(tr, sf)
	defer tr.Close()
	t.Require().NoError(err)

	ctx := context.Background()
	db := t.FileService.DB
	db.File.Create().SetID(1).SetTenantID(1000).SetSourceID(0).SetCreatedBy(1).
		SetName("suite.yaml").SetPath("tmp/suite.yaml").SetMineType("text/yaml").SetSize(0).SaveX(ctx)
}

func (t *fileServiceSuite) Test_Upload() {
	t.Run("normal", func() {
		testFile := testdata.Path("testdata/etc/app.yaml")
		// 使用bufio创建上传文件数据
		payload := &bytes.Buffer{}
		mrw := multipart.NewWriter(payload)
		t.NoError(mrw.WriteField("key", "tmp/app.yaml"))
		t.NoError(mrw.WriteField("bucket", "test-bucket"))

		formFile, _ := mrw.CreateFormFile("file", filepath.Base(testFile))
		fh, err := os.Open(test.Path("testdata/etc/app.yaml"))
		t.Require().NoError(err)
		_, _ = io.Copy(formFile, fh)
		fh.Close()
		_ = mrw.Close()

		// 创建http请求
		req := httptest.NewRequest("POST", "/files", payload)
		req.Header.Set("Authorization", "Bearer "+adminToken)
		req.Header.Set("Content-Type", mrw.FormDataContentType())
		req.Header.Set("X-Tenant-ID", "1000")
		// 发起请求并获取响应
		w := httptest.NewRecorder()
		t.server.Router().ServeHTTP(w, req)
		resp := w.Result()

		// 断言响应状态码和文件信息
		if !t.Equal(200, resp.StatusCode) {
			print(w.Body.String())
		}
	})
}

func (t *fileServiceSuite) Test_GetFileRaw() {
	// 创建http请求
	req := httptest.NewRequest("GET", "/files/1/raw", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1000")
	// 发起请求并获取响应
	w := httptest.NewRecorder()
	t.server.Router().ServeHTTP(w, req)
	resp := w.Result()

	t.Equal(200, resp.StatusCode)
	t.Equal("text/yaml", resp.Header.Get("Content-Type"))
}

func (t *fileServiceSuite) Test_GetFile() {
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

func (t *fileServiceSuite) Test_Delete() {
	req := httptest.NewRequest("DELETE", "/files/1", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1000")
	w := httptest.NewRecorder()
	t.server.Router().ServeHTTP(w, req)
	resp := w.Result()

	t.Equal(200, resp.StatusCode)
	t.NoFileExists(test.Path("tmp/1000/tmp/suite.yaml"))
}
