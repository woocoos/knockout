package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tsingsun/woocoo/pkg/conf"
	"time"
)

type Minio struct {
	ctx       context.Context
	config    *MinioConfig
	stsClient *sts.Client
	s3Client  *s3.Client
}

type MinioConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
	RoleArn         string `json:"roleArn"`
	Policy          string `json:"policy"`
	DurationSeconds int    `json:"durationSeconds"`
}

func NewMinio(cnf *conf.Configuration) *Minio {
	svc := &Minio{
		ctx: context.TODO(),
	}
	config := &MinioConfig{}
	if err := cnf.Unmarshal(config); err != nil {
		panic(err)
	}
	svc.config = config

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: config.Endpoint,
		}, nil
	})
	creds := credentials.NewStaticCredentialsProvider(config.AccessKeyId, config.SecretAccessKey, "")
	cfg, err := awsCfg.LoadDefaultConfig(svc.ctx, awsCfg.WithCredentialsProvider(creds), awsCfg.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(cfg, func(options *sts.Options) {
		options.Region = config.Region
	})
	svc.stsClient = stsClient

	s3Client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.Region = config.Region
	})
	svc.s3Client = s3Client

	bucket := "test2"
	presignClient := s3.NewPresignClient(svc.s3Client)
	req, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    aws.String("63f91559865b894cbe885e9a49d92a29.jpeg"),
	}, func(o *s3.PresignOptions) {
		o.Expires = 20 * time.Minute
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(req.URL)
	return svc
}

func (svc *Minio) GetSTS(roleSessionName string) (*STSResponse, error) {
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(svc.config.RoleArn),
		Policy:          aws.String(svc.config.Policy),
		RoleSessionName: aws.String(roleSessionName),
	}
	out, err := svc.stsClient.AssumeRole(svc.ctx, input)
	if err != nil {
		return nil, err
	}
	return &STSResponse{
		AccessKeyID:     *out.Credentials.AccessKeyId,
		SecretAccessKey: *out.Credentials.SecretAccessKey,
		SessionToken:    *out.Credentials.SessionToken,
		Expiration:      *out.Credentials.Expiration,
	}, nil
}
