package oss

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tsingsun/woocoo/pkg/conf"
	"testing"
)

func TestMinioSTS(t *testing.T) {
	b := []byte(`
oss:
  1:
    kind: "minio"
    config:
      endpoint: "http://localhost:9000"
      accessKeyId: "test1"
      secretAccessKey: "test1234"
      region: "cn-east-1"
      roleArn: "arn:aws:s3:::*"
      policy: ""
      durationSeconds: 3600
`)
	cfg := conf.NewFromBytes(b)
	oss, err := NewService(cfg.Sub("oss"))
	assert.NoError(t, err)
	srv, err := oss.GetProvider(1)
	assert.NoError(t, err)
	resp, err := srv.GetSTS("")
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestAliSTS(t *testing.T) {
	b := []byte(`
oss:
  1:
    kind: "aliOSS"
    config:
      endpoint: "sts.cn-shenzhen.aliyuncs.com"
      accessKeyId: "xxx"
      secretAccessKey: "xxxxx"
      region: "oss-cn-shenzhen"
      roleArn: "xxx"
      policy: ""
      durationSeconds: 3600
`)
	cfg := conf.NewFromBytes(b)
	oss, err := NewService(cfg.Sub("oss"))
	assert.NoError(t, err)
	srv, err := oss.GetProvider(1)
	assert.NoError(t, err)
	resp, err := srv.GetSTS("test")
	assert.NoError(t, err)
	fmt.Println(resp)
}
