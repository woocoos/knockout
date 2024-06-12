package oss

import (
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"strconv"
	"time"
)

type Provider interface {
	GetSTS(roleSessionName string) (*STSResponse, error)
}

type STSResponse struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

type Service struct {
	cnf       *conf.Configuration
	providers map[int]Provider
}

func NewService(cnf *conf.Configuration) (*Service, error) {
	if cnf == nil {
		return nil, fmt.Errorf("unknow config")
	}
	svc := &Service{
		cnf:       cnf,
		providers: make(map[int]Provider),
	}
	return svc, nil
}

func (svc *Service) GetProvider(tid int) (Provider, error) {
	v, ok := svc.providers[tid]
	if ok {
		return v, nil
	}
	subCnf := svc.cnf.Sub(strconv.Itoa(tid))
	kind := subCnf.String("kind")
	switch kind {
	case "minio":
		return NewMinio(subCnf.Sub("config")), nil
	case "aliOSS":
		return NewAliOSS(subCnf.Sub("config")), nil
	default:
		return nil, fmt.Errorf("service type: %s is not supported", kind)
	}
}
