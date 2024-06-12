package oss

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tsingsun/woocoo/pkg/conf"
	"time"
)

type AliOSS struct {
	config    *AliOSSConfig
	stsClient *sts20150401.Client
}

type AliOSSConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
	RoleArn         string `json:"roleArn"`
	Policy          string `json:"policy"`
	DurationSeconds int    `json:"durationSeconds"`
}

func NewAliOSS(cnf *conf.Configuration) *AliOSS {
	svc := &AliOSS{}
	config := &AliOSSConfig{}
	if err := cnf.Unmarshal(config); err != nil {
		panic(err)
	}
	svc.config = config

	cfg := &openapi.Config{
		AccessKeyId:     tea.String(svc.config.AccessKeyId),
		AccessKeySecret: tea.String(svc.config.SecretAccessKey),
	}
	cfg.Endpoint = tea.String(svc.config.Endpoint)
	stsClient, err := sts20150401.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	svc.stsClient = stsClient
	return svc
}

func (svc *AliOSS) GetSTS(roleSessionName string) (*STSResponse, error) {
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleSessionName: tea.String(roleSessionName),
		RoleArn:         tea.String(svc.config.RoleArn),
	}
	if svc.config.Policy != "" {
		assumeRoleRequest.Policy = tea.String(svc.config.Policy)
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
