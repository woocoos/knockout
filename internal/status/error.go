package status

import "errors"

const (
	ErrUserNameOrPassword     = 10000
	ErrUserHasLocked          = 10001
	ErrUserCanNotLogin        = 10002
	ErrInvalidUser            = 10003
	ErrPasswordNotMatch       = 10004
	ErrPasswordExpired        = 10005
	ErrPasswordDuplicate      = 10006
	ErrOldPasswordNotMatch    = 10007
	ErrCaptchaNotMatch        = 10008
	ErrMfaInvalidCode         = 10009
	ErrMfaNotActive           = 10010
	ErrMfaDisable             = 10011
	ErrInvalidToken           = 10012
	ErrEmailEmpty             = 10013
	ErrEmailVerify            = 10014
	ErrOrgNotFound            = 10015
	ErrInvalidPermission      = 10016
	ErrNotLoginPermission     = 10017
	ErrInvalidSpm             = 10018
	ErrClientIdOrClientSecret = 10019
	ErrFileIdentityIsNull     = 10020
	ErrUnsupportedVerify      = 10021
)

var (
	ErrMismatchPWD = errors.New("password not match")
	ErrorCodeMap   = map[int]string{
		ErrUserNameOrPassword:     "用户名或密码错误",
		ErrPasswordNotMatch:       "密码错误",
		ErrOldPasswordNotMatch:    "旧密码错误",
		ErrUserHasLocked:          "账号已锁定，请联系客服修改密码解除锁定",
		ErrPasswordExpired:        "密码已过期，请重置密码或联系客服修改密码恢复",
		ErrCaptchaNotMatch:        "验证码错误",
		ErrMfaInvalidCode:         "身份验证码错误",
		ErrUserCanNotLogin:        "账号禁止登录",
		ErrInvalidToken:           "无效的令牌",
		ErrPasswordDuplicate:      "新密码不能与最近一次使用的密码相同",
		ErrMfaNotActive:           "MFA未激活",
		ErrMfaDisable:             "MFA未启用",
		ErrEmailEmpty:             "邮箱不存在",
		ErrEmailVerify:            "未找到邮箱，请确认邮箱是否正确",
		ErrInvalidPermission:      "无效的权限",
		ErrOrgNotFound:            "组织未找到",
		ErrNotLoginPermission:     "没有登录权限",
		ErrInvalidSpm:             "无效的SPM",
		ErrClientIdOrClientSecret: "the clientID or clientSecret is incorrect or the status is not active",
		ErrFileIdentityIsNull:     "文件凭证不存在",
		ErrUnsupportedVerify:      "不支持的验证方式",
		ErrInvalidUser:            "账号验证错误",
	}
)
