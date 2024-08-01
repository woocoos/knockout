package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	awsCredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"net/url"
	"strings"
	"time"
)

type AwsS3 struct {
	ctx        context.Context
	stsClient  *sts.Client
	s3Client   *s3.Client
	fileSource *FileSource
}

func NewAwsS3(ctx context.Context, fileSource *FileSource) (*AwsS3, error) {
	svc := &AwsS3{
		ctx:        ctx,
		fileSource: fileSource,
	}
	err := svc.initAwsClient()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

// initAwsClient 初始化aws client
func (svc *AwsS3) initAwsClient() error {
	// 初始化sts
	stsClient, err := initAwsSTS(svc.ctx, svc.fileSource)
	if err != nil {
		return err
	}
	svc.stsClient = stsClient
	// 初始化s3Client
	s3Client, err := initAwsClient(svc.ctx, svc.fileSource)
	if err != nil {
		return err
	}
	svc.s3Client = s3Client
	return nil
}

// initAwsSTS 初始化sts
func initAwsSTS(ctx context.Context, fs *FileSource) (*sts.Client, error) {
	creds := awsCredentials.NewStaticCredentialsProvider(fs.AccessKeyID, fs.AccessKeySecret, "")
	cfg, err := awsCfg.LoadDefaultConfig(ctx, awsCfg.WithCredentialsProvider(creds))
	if err != nil {
		return nil, err
	}

	stsClient := sts.NewFromConfig(cfg, func(options *sts.Options) {
		options.Region = fs.Region
		options.BaseEndpoint = aws.String(fs.StsEndpoint)
	})
	return stsClient, nil
}

// initAwsClient 初始化s3Client
func initAwsClient(ctx context.Context, fs *FileSource) (*s3.Client, error) {
	creds := awsCredentials.NewStaticCredentialsProvider(fs.AccessKeyID, fs.AccessKeySecret, "")
	cfg, err := awsCfg.LoadDefaultConfig(ctx, awsCfg.WithCredentialsProvider(creds))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.Region = fs.Region
		// 自定义endpoint resolver
		options.EndpointResolverV2 = &EndpointResolverV2{
			EndpointImmutable: fs.EndpointImmutable, // 如果为true，则启用自定义域名
			endpoint:          fs.Endpoint,
		}
		options.BaseEndpoint = aws.String(fs.Endpoint)
		// 如果是minio，则使用path style
		if fs.Kind == KindMinio {
			options.UsePathStyle = true
		}
	})
	return s3Client, nil
}

// GetSTS 获取sts
func (svc *AwsS3) GetSTS(roleSessionName string) (*STSResponse, error) {
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
	}, nil
}

// GetPreSignedURL 获取预签名url
func (svc *AwsS3) GetPreSignedURL(bucket, path string, expires time.Duration) (string, error) {
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

// GetS3Client 获取s3Client
func (svc *AwsS3) GetS3Client() *s3.Client {
	return svc.s3Client
}

// EndpointResolverV2 自定义endpoint resolver
type EndpointResolverV2 struct {
	endpoint          string
	EndpointImmutable bool
}

// ResolveEndpoint resolve endpoint
func (r *EndpointResolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	// if the endpoint is immutable, return the configured endpoint
	if r.EndpointImmutable {
		u, err := url.Parse(r.endpoint)
		if err != nil {
			return smithyendpoints.Endpoint{}, err
		}
		return smithyendpoints.Endpoint{
			URI: *u,
		}, nil
	}
	// delegate back to the default v2 resolver otherwise
	return s3.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
}
