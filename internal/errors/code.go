package errors

const (
	ErrUserNameOrPassword  = 10000
	ErrUserHasLocked       = 10001
	ErrUserCanNotLogin     = 10002
	ErrInvalidUser         = 10003
	ErrPasswordNotMatch    = 10004
	ErrPasswordExpired     = 10005
	ErrPasswordDuplicate   = 10006
	ErrOldPasswordNotMatch = 10007
	ErrCaptchaNotMatch     = 10008
	ErrMfaInvalidCode      = 10009
	ErrMfaNotActive        = 10010
	ErrMfaDisable          = 10011
	ErrInvalidToken        = 10012
	ErrEmailEmpty          = 10013
	ErrEmailVerify         = 10014
	ErrOrgNotFound         = 10015
	ErrInvalidPermission   = 10016
	ErrNotLoginPermission  = 10017
	ErrInvalidSpm          = 10018
	ErrFileIdentityIsNull  = 10019
	ErrUnsupportedVerify   = 10020
	ErrCaptchaInvalid      = 10021
	// ErrPasswordRetry 密码错误，您还可以尝试%d次
	ErrPasswordRetry = 10022
	// ErrUserDeviceLimit 登录设备超过%d台限制，请前往旧设备删除登录设备后登录
	ErrUserDeviceLimit      = 10023
	ErrUserIdentityNotFound = 10024
)
