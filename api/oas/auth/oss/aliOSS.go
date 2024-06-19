package oss

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/woocoos/knockout/ent"
	"net/url"
	"strings"
	"time"
)

type AliOSS struct {
	stsClient  *sts20150401.Client
	ossClient  *oss.Client
	fileSource *ent.FileSource
}

func NewAliOSS(fileSource *ent.FileSource) (*AliOSS, error) {
	svc := &AliOSS{
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
	return svc, nil
}

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

func (svc *AliOSS) initAliOSS() error {
	client, err := oss.New(svc.fileSource.Endpoint, svc.fileSource.AccessKeyID, svc.fileSource.AccessKeySecret)
	if err != nil {
		return err
	}
	svc.ossClient = client
	return nil
}

func (svc *AliOSS) GetSTS(roleSessionName string) (*STSResponse, error) {
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleSessionName: tea.String(roleSessionName),
		RoleArn:         tea.String(svc.fileSource.RoleArn),
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
		Endpoint:        svc.fileSource.Endpoint,
		Bucket:          svc.fileSource.Bucket,
		Region:          svc.fileSource.Region,
		Kind:            svc.fileSource.Kind.String(),
	}, nil
}

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

func ParseAliOSSUrl(ossUrl string) (bucketUrl, path string, err error) {
	u, err := url.Parse(ossUrl)
	if err != nil {
		return
	}
	bucketUrl = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	path = u.Path
	return
}
