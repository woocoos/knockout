package status

import "errors"

var (
	ErrMismatchPWD         = errors.New("password not match")
	ErrUserOrPWD           = errors.New("username or password error")
	ErrLoginFailUpperLimit = errors.New("login fail upper limit,and account is locked")
	ErrCaptchaNotMatch     = errors.New("captcha not match")
)
