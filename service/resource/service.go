package resource

import (
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
)

type Option func(*Service)

type PwdPolicy struct {
	Length               int32 `json:"length"`
	IncludeElement       int32 `json:"includeElement"`
	IncludeChar          int32 `json:"includeChar"`
	AllowIncludeUserName bool  `json:"allowIncludeUserName"`
	InvalidDay           int32 `json:"invalidDay"`
	InvalidLoginLimit    bool  `json:"invalidLoginLimit"`
	Retry                int32 `json:"retry"`
	CaptchaTimes         int32 `json:"captchaTimes"`
}

// Service 企业目录服务管理
type Service struct {
	Client    *ent.Client
	KOSDK     *api.SDK
	Cfg       *conf.AppConfiguration
	PwdPolicy *PwdPolicy
}

func WithClient(client *ent.Client) Option {
	return func(s *Service) {
		s.Client = client
	}
}

func WithKOSDK(sdk *api.SDK) Option {
	return func(s *Service) {
		s.KOSDK = sdk
	}
}

func WithCfg(cnf *conf.AppConfiguration) Option {
	return func(s *Service) {
		s.Cfg = cnf
	}
}

func WithPwdPolicy(pp *PwdPolicy) Option {
	return func(s *Service) {
		s.PwdPolicy = pp
	}
}

func NewService(opt ...Option) *Service {
	r := &Service{}
	for _, option := range opt {
		option(r)
	}
	return r
}
