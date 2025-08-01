package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Join wrapper errors.join
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// As wrapper errors.As
func As(err error, target any) bool {
	return errors.As(err, &target)
}

var (
	ErrUndefined = errors.New("undefined error")
)

func Code(errCode gin.ErrorType) *gin.Error {
	return &gin.Error{
		Type: errCode,
		Err:  ErrUndefined,
	}
}

// Codel create a new error by code and args. args set to Meta will use int key.
func Codel(code gin.ErrorType, a ...any) *gin.Error {
	e := Code(code)
	if a != nil {
		e.Meta = a
	}
	return e
}
