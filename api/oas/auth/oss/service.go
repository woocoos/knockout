package oss

import (
	"fmt"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/filesource"
	"time"
)

type Provider interface {
	GetSTS(roleSessionName string) (*STSResponse, error)
	GetPreSignedURL(bucket, path string, expires time.Duration) (string, error)
}

type STSResponse struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
	Endpoint        string
	Bucket          string
	Region          string
	Kind            string
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

func (svc *Service) GetProvider(fs *ent.FileSource) (Provider, error) {
	v, ok := svc.providers[getProviderKey(fs)]
	if ok {
		return v, nil
	}
	switch fs.Kind {
	case filesource.KindMinio:
		provider, err := NewMinio(fs)
		if err != nil {
			return nil, err
		}
		svc.providers[getProviderKey(fs)] = provider
		return provider, nil
	case filesource.KindAliOSS:
		provide, err := NewAliOSS(fs)
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
func (svc *Service) CleanProvider(fs *ent.FileSource) {
	delete(svc.providers, getProviderKey(fs))
}

func getProviderKey(fs *ent.FileSource) string {
	return fmt.Sprintf("%d:%s:%s:%s", fs.TenantID, fs.Endpoint, fs.Bucket, fs.Kind)
}
