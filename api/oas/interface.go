// Code generated by woco, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server is the server API for  service.
type Server interface {
	// (GET /captcha)
	Captcha(*gin.Context, *CaptchaRequest) ([]byte, error)
	// (POST /login/auth)
	Login(*gin.Context, *LoginRequest) (*LoginResponse, error)
	// (POST /logout)
	Logout(*gin.Context) error
	// (POST /login/reset-password)
	ResetPassword(*gin.Context, *ResetPasswordRequest) (*LoginResponse, error)
	// (POST /login/verify-factor)
	VerifyFactor(*gin.Context, *VerifyFactorRequest) (*LoginResponse, error)
}

type UnimplementedServer struct {
}

func (UnimplementedServer) Captcha(c *gin.Context, req *CaptchaRequest) (_ []byte, err error) {
	err = fmt.Errorf("method Captcha not implemented")
	return
}

func (UnimplementedServer) Login(c *gin.Context, req *LoginRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method Login not implemented")
	return
}

func (UnimplementedServer) Logout(c *gin.Context) (err error) {
	err = fmt.Errorf("method Logout not implemented")
	return
}

func (UnimplementedServer) ResetPassword(c *gin.Context, req *ResetPasswordRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method ResetPassword not implemented")
	return
}

func (UnimplementedServer) VerifyFactor(c *gin.Context, req *VerifyFactorRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method VerifyFactor not implemented")
	return
}
