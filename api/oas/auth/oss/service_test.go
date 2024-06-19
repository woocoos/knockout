package oss

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/filesource"
	"testing"
	"time"
)

var (
	minioFileSource = ent.FileSource{
		Kind:            filesource.KindMinio,
		AccessKeyID:     "test1",
		AccessKeySecret: "test1234",
		Endpoint:        "localhost:9000",
		StsEndpoint:     "http://localhost:9000",
		Region:          "cn-east-1",
		RoleArn:         "arn:aws:s3:::*",
		Policy:          "",
		DurationSeconds: 3600,
		Bucket:          "test2",
		BucketUrl:       "http://localhost:9000/test2",
	}
	aliFileSource = ent.FileSource{
		Kind:            filesource.KindAliOSS,
		AccessKeyID:     "LTAI5tDStcqxb8q7MkJXo54M",
		AccessKeySecret: "xxx",
		Endpoint:        "https://oss-cn-shenzhen.aliyuncs.com",
		StsEndpoint:     "sts.cn-shenzhen.aliyuncs.com",
		RoleArn:         "acs:ram::5755321561100682:role/devossrwrole",
		Bucket:          "qldevtest",
		BucketUrl:       "https://qldevtest.oss-cn-shenzhen.aliyuncs.com",
		Region:          "oss-cn-shenzhen",
		Policy:          "",
		DurationSeconds: 3600,
	}
)

func TestMinioSTS(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&minioFileSource)
	assert.NoError(t, err)
	resp, err := provider.GetSTS("")
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestAliSTS(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&aliFileSource)
	assert.NoError(t, err)
	resp, err := provider.GetSTS("test")
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestMinioPreSignedUrl(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&minioFileSource)
	assert.NoError(t, err)
	u, err := provider.GetPreSignedURL("test2", "63f91559865b894cbe885e9a49d92a29.jpeg", time.Hour)
	assert.NoError(t, err)
	fmt.Println(u)
}

func TestAliOSSPreSignedUrl(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&aliFileSource)
	assert.NoError(t, err)
	u, err := provider.GetPreSignedURL("qldevtest", "cust/159ecc5f964dfe00", time.Hour)
	assert.NoError(t, err)
	fmt.Println(u)
}

func TestAliUrlParse(t *testing.T) {
	aliUrl := "https://qldevtest.oss-cn-shenzhen.aliyuncs.com/cust%2F159ecc5f964dfe00?Expires=1718624031&OSSAccessKeyId=LTAI5tDStcqxb8q7MkJXo54M&Signature=bvSneABFjFZiYz75N8wX%2F2HgxbY%3D"
	bucketUrl, path, err := ParseAliOSSUrl(aliUrl)
	assert.NoError(t, err)
	fmt.Println(bucketUrl, path)
}

func TestMinioUrlParse(t *testing.T) {
	minioUrl := "http://localhost:9000/test2/63f91559865b894cbe885e9a49d92a29.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=test1%2F20240618%2Fcn-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240618T085259Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=cf0ac630cc5f39b106721a200ff4d178b53860d5c64d0e57fd59769b0825b035"
	bucketUrl, path, err := ParseMinioUrl(minioUrl)
	assert.NoError(t, err)
	fmt.Println(bucketUrl, path)
}
