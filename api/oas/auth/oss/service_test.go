package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/knockout/ent/filesource"
	"testing"
	"time"
)

var (
	minioFileSource = FileSource{
		Kind:            filesource.KindMinio,
		AccessKeyID:     "test",
		AccessKeySecret: "test1234",
		Endpoint:        "http://192.168.0.17:32650",
		StsEndpoint:     "http://192.168.0.17:32650",
		Region:          "minio",
		RoleArn:         "arn:aws:s3:::*",
		Policy:          "",
		DurationSeconds: 3600,
		Bucket:          "woocootest",
		BucketUrl:       "http://192.168.0.17:32650/woocootest",
	}
	aliFileSource = FileSource{
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

func TestMinioPreSignedUrl(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&minioFileSource)
	assert.NoError(t, err)
	u, err := provider.GetPreSignedURL("woocootest", "3a9809ba339ec87f1636c7878685f616.jpeg", time.Hour)
	assert.NoError(t, err)
	fmt.Println(u)
}

func TestMinioUrlParse(t *testing.T) {
	minioUrl := "http://192.168.0.17:32650/woocootest/3a9809ba339ec87f1636c7878685f616.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=test%2F20240708%2Fminio%2Fs3%2Faws4_request&X-Amz-Date=20240708T024958Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=9b8120f422dddb6a084e4c7b2cb5a74df9dc4b780fff488aef6c94e628af7bd7"
	bucketUrl, path, err := ParseMinioUrl(minioUrl)
	assert.NoError(t, err)
	fmt.Println(bucketUrl, path)
}

func TestMinioS3GetObject(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&minioFileSource)
	assert.NoError(t, err)
	s3Client := provider.GetS3Client()
	out, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{Bucket: aws.String("woocootest"), Key: aws.String("/3a9809ba339ec87f1636c7878685f616.jpeg")})
	assert.NoError(t, err)
	defer out.Body.Close()
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

func TestAliS3GetObject(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(&aliFileSource)
	assert.NoError(t, err)
	s3Client := provider.GetS3Client()
	out, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{Bucket: aws.String("qldevtest"), Key: aws.String("cust/159ecc5f964dfe00")})
	assert.NoError(t, err)
	defer out.Body.Close()
}
