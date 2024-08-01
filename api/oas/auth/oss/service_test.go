package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	minioFileSource = FileSource{
		Kind:              KindMinio,
		AccessKeyID:       "test",
		AccessKeySecret:   "test1234",
		Endpoint:          "http://192.168.0.17:32650",
		EndpointImmutable: false,
		StsEndpoint:       "http://192.168.0.17:32650",
		Region:            "minio",
		RoleArn:           "arn:aws:s3:::*",
		Policy:            "",
		DurationSeconds:   3600,
		Bucket:            "woocootest",
		BucketUrl:         "http://192.168.0.17:32650/woocootest",
	}
	aliFileSource = FileSource{
		Kind:              KindAliOSS,
		AccessKeyID:       "LTAI5tDStcqxb8q7MkJXo54M",
		AccessKeySecret:   "xxx",
		Endpoint:          "https://oss-cn-shenzhen.aliyuncs.com",
		EndpointImmutable: false,
		StsEndpoint:       "sts.cn-shenzhen.aliyuncs.com",
		RoleArn:           "acs:ram::5755321561100682:role/devossrwrole",
		Bucket:            "qldevtest",
		BucketUrl:         "https://qldevtest.oss-cn-shenzhen.aliyuncs.com",
		Region:            "oss-cn-shenzhen",
		Policy:            "",
		DurationSeconds:   3600,
	}
)

func TestMinioSTS(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &minioFileSource)
	assert.NoError(t, err)
	resp, err := provider.GetSTS("")
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestMinioPreSignedUrl(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &minioFileSource)
	assert.NoError(t, err)
	u, err := provider.GetPreSignedURL("woocootest", "3a9809ba339ec87f1636c7878685f616.jpeg", time.Hour)
	assert.NoError(t, err)
	fmt.Println(u)
}

func TestMinioS3GetObject(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &minioFileSource)
	assert.NoError(t, err)
	s3Client := provider.GetS3Client()
	out, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{Bucket: aws.String("woocootest"), Key: aws.String("/3a9809ba339ec87f1636c7878685f616.jpeg")})
	assert.NoError(t, err)
	defer out.Body.Close()
}

func TestAliSTS(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &aliFileSource)
	assert.NoError(t, err)
	resp, err := provider.GetSTS("test")
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestAliOSSPreSignedUrl(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &aliFileSource)
	assert.NoError(t, err)
	u, err := provider.GetPreSignedURL("qldevtest", "cust/159ecc5f964dfe00", time.Hour)
	assert.NoError(t, err)
	fmt.Println(u)
}

func TestAliS3GetObject(t *testing.T) {
	oss, err := NewService()
	assert.NoError(t, err)
	provider, err := oss.GetProvider(context.TODO(), &aliFileSource)
	assert.NoError(t, err)
	s3Client := provider.GetS3Client()
	out, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{Bucket: aws.String("qldevtest"), Key: aws.String("cust/159ecc5f964dfe00")})
	assert.NoError(t, err)
	defer out.Body.Close()
}
