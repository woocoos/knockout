package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

type Provider interface {
	GetSTS(roleSessionName string) (*STSResponse, error)
	GetPreSignedURL(bucket, path string, expires time.Duration) (string, error)
	GetS3Client() *s3.Client
}

type Kind string

// Kind values.
const (
	KindMinio  Kind = "minio"
	KindAliOSS Kind = "aliOSS"
	KindAwsS3  Kind = "awsS3"
)

type FileSource struct {
	// 租户ID
	TenantID int `json:"tenant_id,omitempty"`
	// 文件来源
	Kind Kind `json:"kind,omitempty"`
	// accesskey id
	AccessKeyID string `json:"access_key_id,omitempty"`
	// accesskey secret
	AccessKeySecret string `json:"access_key_secret,omitempty"`
	// 对外服务的访问域名
	Endpoint string `json:"endpoint,omitempty"`
	// endpoint是否不可变,如果是自定义域名则为true
	EndpointImmutable bool `json:"endpointImmutable"`
	// sts服务的访问域名
	StsEndpoint string `json:"sts_endpoint,omitempty"`
	// 地域，数据存储的物理位置
	Region string `json:"region,omitempty"`
	// 文件存储空间
	Bucket string `json:"bucket,omitempty"`
	// 文件存储空间地址，用于匹配url
	BucketUrl string `json:"bucketUrl,omitempty"`
	// 角色的资源名称(ARN)，用于STS
	RoleArn string `json:"role_arn,omitempty"`
	// 指定返回的STS令牌的权限的策略
	Policy string `json:"policy,omitempty"`
	// STS令牌的有效期，默认3600s
	DurationSeconds int `json:"duration_seconds,omitempty"`
}

type STSResponse struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

type Service struct {
	providers map[string]Provider
}

func NewService() (*Service, error) {
	svc := &Service{
		providers: make(map[string]Provider),
	}
	return svc, nil
}

func (svc *Service) GetProvider(ctx context.Context, fs *FileSource) (Provider, error) {
	v, ok := svc.providers[getProviderKey(fs)]
	if ok {
		return v, nil
	}
	switch fs.Kind {
	case KindMinio, KindAwsS3:
		provider, err := NewAwsS3(ctx, fs)
		if err != nil {
			return nil, err
		}
		svc.providers[getProviderKey(fs)] = provider
		return provider, nil
	case KindAliOSS:
		provide, err := NewAliOSS(ctx, fs)
		if err != nil {
			return nil, err
		}
		svc.providers[getProviderKey(fs)] = provide
		return provide, nil
	default:
		return nil, fmt.Errorf("service type: %s is not supported", fs.Kind)
	}
}

// CleanProvider 清除缓存中的provider
func (svc *Service) CleanProvider(fs *FileSource) {
	delete(svc.providers, getProviderKey(fs))
}

func getProviderKey(fs *FileSource) string {
	return fmt.Sprintf("%d:%s:%s:%s", fs.TenantID, fs.Endpoint, fs.Bucket, fs.Kind)
}
