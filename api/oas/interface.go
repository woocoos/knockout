// Code generated by woco, DO NOT EDIT.

package oas

import (
	"context"
	"fmt"
)

// Server is the server API for  service.
type Server interface {
	// (GET /captcha)
	Captcha(context.Context, *CaptchaRequest) ([]byte, error)
	// (POST /login/auth)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// (POST /logout)
	Logout(context.Context) error
	// (POST /login/reset-password)
	ResetPassword(context.Context, *ResetPasswordRequest) (*LoginResponse, error)
	// (POST /login/verify-factor)
	VerifyFactor(context.Context, *VerifyFactorRequest) (*LoginResponse, error)
}

type UnimplementedServer struct {
}

func (UnimplementedServer) Captcha(ctx context.Context, req *CaptchaRequest) (_ []byte, err error) {
	err = fmt.Errorf("method Captcha not implemented")
	return
}

func (UnimplementedServer) Login(ctx context.Context, req *LoginRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method Login not implemented")
	return
}

func (UnimplementedServer) Logout(ctx context.Context) (err error) {
	err = fmt.Errorf("method Logout not implemented")
	return
}

func (UnimplementedServer) ResetPassword(ctx context.Context, req *ResetPasswordRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method ResetPassword not implemented")
	return
}

func (UnimplementedServer) VerifyFactor(ctx context.Context, req *VerifyFactorRequest) (_ *LoginResponse, err error) {
	err = fmt.Errorf("method VerifyFactor not implemented")
	return
}
