package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	awsCredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/minio/minio-go/v7"
	minioCredentials "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/woocoos/knockout/ent"
	"net/url"
	"strings"
	"time"
)

type Minio struct {
	ctx         context.Context
	stsClient   *sts.Client
	minioClient *minio.Client
	fileSource  *ent.FileSource
}

func NewMinio(fileSource *ent.FileSource) (*Minio, error) {
	svc := &Minio{
		ctx:        context.TODO(),
		fileSource: fileSource,
	}

	err := svc.initAwsClient()
	if err != nil {
		return nil, err
	}
	err = svc.initMinioClient()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (svc *Minio) initAwsClient() error {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: svc.fileSource.StsEndpoint,
		}, nil
	})
	creds := awsCredentials.NewStaticCredentialsProvider(svc.fileSource.AccessKeyID, svc.fileSource.AccessKeySecret, "")
	cfg, err := awsCfg.LoadDefaultConfig(svc.ctx, awsCfg.WithCredentialsProvider(creds), awsCfg.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return err
	}

	stsClient := sts.NewFromConfig(cfg, func(options *sts.Options) {
		options.Region = svc.fileSource.Region
	})
	svc.stsClient = stsClient
	return nil
}

func (svc *Minio) initMinioClient() error {
	minioClient, err := minio.New(svc.fileSource.Endpoint, &minio.Options{
		Creds:  minioCredentials.NewStaticV4(svc.fileSource.AccessKeyID, svc.fileSource.AccessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	svc.minioClient = minioClient
	return nil
}

func (svc *Minio) GetSTS(roleSessionName string) (*STSResponse, error) {
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(svc.fileSource.RoleArn),
		Policy:          aws.String(svc.fileSource.Policy),
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
		Endpoint:        svc.fileSource.Endpoint,
		Bucket:          svc.fileSource.Bucket,
		Region:          svc.fileSource.Region,
		Kind:            svc.fileSource.Kind.String(),
	}, nil
}

func (svc *Minio) GetPreSignedURL(bucket, path string, expires time.Duration) (string, error) {
	preSignedURL, err := svc.minioClient.PresignedGetObject(svc.ctx, bucket, strings.TrimLeft(path, "/"), expires, make(url.Values))
	if err != nil {
		panic(err)
	}
	return preSignedURL.String(), nil
}

func ParseMinioUrl(ossUrl string) (bucketUrl, path string, err error) {
	u, err := url.Parse(ossUrl)
	if err != nil {
		return
	}
	bucketPath := ""
	if u.Path != "" {
		trimPath := strings.TrimLeft(u.Path[1:], "/")
		index := strings.Index(trimPath, "/")
		if index != -1 {
			bucketPath = "/" + trimPath[:index]
			path = strings.TrimPrefix(u.Path, bucketPath)
		}
	}
	bucketUrl = fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, bucketPath)
	return
}
