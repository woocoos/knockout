package oss

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
	"time"
)

type AliOSS struct {
	ctx        context.Context
	stsClient  *sts20150401.Client
	ossClient  *oss.Client
	s3Client   *s3.Client
	fileSource *FileSource
}

func NewAliOSS(ctx context.Context, fileSource *FileSource) (*AliOSS, error) {
	svc := &AliOSS{
		ctx:        ctx,
		fileSource: fileSource,
	}
	err := svc.initAliSTS()
	if err != nil {
		return nil, err
	}
	err = svc.initAliOSS()
	if err != nil {
		return nil, err
	}
	err = svc.initAwsClient()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

// initAliSTS 初始化阿里云STS客户端
func (svc *AliOSS) initAliSTS() error {
	cfg := &openapi.Config{
		AccessKeyId:     tea.String(svc.fileSource.AccessKeyID),
		AccessKeySecret: tea.String(svc.fileSource.AccessKeySecret),
		RegionId:        tea.String(svc.fileSource.Region),
	}
	cfg.Endpoint = tea.String(svc.fileSource.StsEndpoint)
	stsClient, err := sts20150401.NewClient(cfg)
	if err != nil {
		return err
	}
	svc.stsClient = stsClient
	return nil
}

// initAliOSS 初始化阿里云OSS客户端
func (svc *AliOSS) initAliOSS() error {
	useCname := false
	// 判断是否自定义域名
	if svc.fileSource.EndpointImmutable {
		useCname = true
	}
	client, err := oss.New(svc.fileSource.Endpoint, svc.fileSource.AccessKeyID, svc.fileSource.AccessKeySecret, oss.UseCname(useCname))
	if err != nil {
		return err
	}
	svc.ossClient = client
	return nil
}

// initAwsClient 初始化AWS客户端
func (svc *AliOSS) initAwsClient() error {
	s3Client, err := initAwsClient(svc.ctx, svc.fileSource)
	if err != nil {
		return err
	}
	svc.s3Client = s3Client
	return nil
}

// GetSTS 获取阿里云STS
func (svc *AliOSS) GetSTS(roleSessionName string) (*STSResponse, error) {
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleSessionName: tea.String(roleSessionName),
		RoleArn:         tea.String(svc.fileSource.RoleArn),
		DurationSeconds: tea.Int64(int64(svc.fileSource.DurationSeconds)),
	}
	if svc.fileSource.Policy != "" {
		assumeRoleRequest.Policy = tea.String(svc.fileSource.Policy)
	}

	resp, err := svc.stsClient.AssumeRoleWithOptions(assumeRoleRequest, &service.RuntimeOptions{})
	if err != nil {
		return nil, err
	}

	expiration, err := time.Parse(time.RFC3339, *resp.Body.Credentials.Expiration)
	if err != nil {
		return nil, err
	}
	return &STSResponse{
		AccessKeyID:     *resp.Body.Credentials.AccessKeyId,
		SecretAccessKey: *resp.Body.Credentials.AccessKeySecret,
		SessionToken:    *resp.Body.Credentials.SecurityToken,
		Expiration:      expiration,
	}, nil
}

// GetPreSignedURL 获取阿里云预签名URL
func (svc *AliOSS) GetPreSignedURL(bucket, path string, expires time.Duration) (string, error) {
	bk, err := svc.ossClient.Bucket(bucket)
	if err != nil {
		return "", err
	}
	signedURL, err := bk.SignURL(strings.TrimLeft(path, "/"), oss.HTTPGet, int64(expires.Seconds()))
	if err != nil {
		return "", err
	}
	return signedURL, nil
}

// GetS3Client 获取AWS S3客户端
func (svc *AliOSS) GetS3Client() *s3.Client {
	return svc.s3Client
}
