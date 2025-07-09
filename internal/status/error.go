package status

import (
	"errors"
	"github.com/gin-gonic/gin"
)

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
	ErrCaptchaInvalid         = 10022
)

var (
	ErrUndefined = errors.New("undefined error")
)

func CodeError(errCode gin.ErrorType) error {
	return &gin.Error{
		Type: errCode,
		Err:  ErrUndefined,
	}
}
