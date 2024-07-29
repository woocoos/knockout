package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	awsCredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"strings"
	"time"
)

type Minio struct {
	ctx        context.Context
	stsClient  *sts.Client
	s3Client   *s3.Client
	fileSource *FileSource
}

func NewMinio(fileSource *FileSource) (*Minio, error) {
	svc := &Minio{
		ctx:        context.TODO(),
		fileSource: fileSource,
	}

	err := svc.initAwsClient()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (svc *Minio) initAwsClient() error {
	creds := awsCredentials.NewStaticCredentialsProvider(svc.fileSource.AccessKeyID, svc.fileSource.AccessKeySecret, "")
	cfg, err := awsCfg.LoadDefaultConfig(svc.ctx, awsCfg.WithCredentialsProvider(creds))
	if err != nil {
		return err
	}

	stsClient := sts.NewFromConfig(cfg, func(options *sts.Options) {
		options.Region = svc.fileSource.Region
		options.BaseEndpoint = aws.String(svc.fileSource.StsEndpoint)
	})
	svc.stsClient = stsClient

	s3Client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.Region = svc.fileSource.Region
		options.EndpointResolverV2 = &EndpointResolverV2{
			EndpointImmutable: svc.fileSource.EndpointImmutable,
			endpoint:          svc.fileSource.Endpoint,
		}
		options.BaseEndpoint = aws.String(svc.fileSource.Endpoint)
		options.UsePathStyle = true
	})
	svc.s3Client = s3Client
	return nil
}

func (svc *Minio) GetSTS(roleSessionName string) (*STSResponse, error) {
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(svc.fileSource.RoleArn),
		Policy:          aws.String(svc.fileSource.Policy),
		RoleSessionName: aws.String(roleSessionName),
		DurationSeconds: aws.Int32(int32(svc.fileSource.DurationSeconds)),
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
	pClient := s3.NewPresignClient(svc.s3Client)
	resp, err := pClient.PresignGetObject(svc.ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(strings.TrimLeft(path, "/")),
	}, func(options *s3.PresignOptions) {
		options.Expires = expires
	})
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("signUrl response is nil")
	}
	return resp.URL, nil
}

func (svc *Minio) GetS3Client() *s3.Client {
	return svc.s3Client
}
