package resource

import (
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
)

type Option func(*Service)

type PwdPolicy struct {
	// 密码最短长度，长度应在6-32位之间
	Length               int32 `json:"length"`
	// 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号
	IncludeElement       int32 `json:"includeElement"`
	// 最少包含的不同字符数，最多8个，0代表不限制
	IncludeChar          int32 `json:"includeChar"`
	// 是否允许包含用户名
	AllowIncludeUserName bool  `json:"allowIncludeUserName"`
	// 有效天数，最大1095天，0代表不过期
	InvalidDay           int32 `json:"invalidDay"`
	// 过期后是否限制登录
	InvalidLoginLimit    bool  `json:"invalidLoginLimit"`
	// 一小时内密码错误最多尝试次数，最大32次，0代表不限次数
	Retry                int32 `json:"retry"`
	// 密码错误多少次出现验证码，最大5次，0代表不出现验证码
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

func NewService(opt ...Option) *Service {
	r := &Service{}
	for _, option := range opt {
		option(r)
	}
	pp := PwdPolicy{
		Length:               6,
		IncludeElement:       3,
		IncludeChar:          4,
		AllowIncludeUserName: false,
		InvalidDay:           30,
		InvalidLoginLimit:    false,
		Retry:                5,
		CaptchaTimes:         3,
	}
	if r.Cfg != nil && r.Cfg.IsSet("adminx.pwdPolicy") {
		err := r.Cfg.Sub("adminx.pwdPolicy").Unmarshal(&pp)
		if err != nil {
			panic(err)
		}
	}
	r.PwdPolicy = &pp
	return r
}
