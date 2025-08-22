package resource

import (
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
)

var (
	// TODO 添加注释
	defaultPwdPolicy = PasswordPolicy{
		Length:               6,
		IncludeElement:       3,
		IncludeChar:          4,
		AllowIncludeUserName: false,
		InvalidDay:           30,
		InvalidLoginLimit:    false,
		Retry:                5,
		CaptchaTimes:         3,
	}
)

type Option func(*Service)

// PasswordPolicy 密码策略
type PasswordPolicy struct {
	// 密码最短长度，长度应在6-32位之间
	Length int32 `json:"length"`
	// 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号
	IncludeElement int32 `json:"includeElement"`
	// 最少包含的不同字符数，最多8个，0代表不限制
	IncludeChar int32 `json:"includeChar"`
	// 是否允许包含用户名
	AllowIncludeUserName bool `json:"allowIncludeUserName"`
	// 有效天数，最大1095天，0代表不过期
	InvalidDay int32 `json:"invalidDay"`
	// 过期后是否限制登录
	InvalidLoginLimit bool `json:"invalidLoginLimit"`
	// 一小时内密码错误最多尝试次数，最大32次，0代表不限次数
	Retry int32 `json:"retry"`
	// 密码错误多少次出现验证码，最大5次，0代表不出现验证码
	CaptchaTimes int32 `json:"captchaTimes"`
}

// JwtConfig 用于JWT,在ko-proxy移除后, 应该删除
type JwtConfig struct {
	SigningMethod string
	SigningKey    string
}

type ClearLoginTokens struct {
	Exclude []int
}

// Service 企业目录服务管理
type Service struct {
	Client      *ent.Client
	redisClient *redisx.Client
	KOSDK       *api.SDK
	cnf         *conf.AppConfiguration
	// 已经暴露一个密码策略, 这边不需要再暴露了
	passwordPolicy   PasswordPolicy
	jwtConfig        JwtConfig
	clearLoginTokens ClearLoginTokens
}

func WithClient(client *ent.Client) Option {
	return func(s *Service) {
		s.Client = client
	}
}

func WithRedis(redisClient *redisx.Client) Option {
	return func(s *Service) {
		s.redisClient = redisClient
	}
}

func WithKOSDK(sdk *api.SDK) Option {
	return func(s *Service) {
		s.KOSDK = sdk
	}
}

func WithCfg(cnf *conf.AppConfiguration) Option {
	return func(s *Service) {
		s.cnf = cnf
	}
}

func NewService(opt ...Option) *Service {
	r := &Service{
		cnf: conf.Global(),
		clearLoginTokens: ClearLoginTokens{
			Exclude: make([]int, 0),
		},
	}
	for _, option := range opt {
		option(r)
	}
	pp := defaultPwdPolicy
	if r.cnf.IsSet("auth.pwdPolicy") {
		err := r.cnf.Sub("auth.pwdPolicy").Unmarshal(&pp)
		if err != nil {
			panic(err)
		}
	}
	r.passwordPolicy = pp
	if r.cnf.IsSet("auth.clearLoginTokens") {
		err := r.cnf.Sub("auth.clearLoginTokens").Unmarshal(&r.clearLoginTokens)
		if err != nil {
			panic(err)
		}
	}
	if err := r.cnf.Sub("jwt").Unmarshal(&r.jwtConfig); err != nil {
		panic(err)
	}
	return r
}
