// Code generated by woco, DO NOT EDIT.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/knockout/api/oas"
)

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.RouterGroup, si oas.Server) {
	router.GET("/captcha", wrapCaptcha(si))
	router.POST("/login/auth", wrapLogin(si))
	router.POST("/logout", wrapLogout(si))
	router.POST("/login/reset-password", wrapResetPassword(si))
	router.POST("/login/verify-factor", wrapVerifyFactor(si))
}

func wrapCaptcha(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.CaptchaRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.Captcha(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		c.Data(http.StatusOK, "image/png", resp)
	}
}

func wrapLogin(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.LoginRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.Login(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapLogout(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := si.Logout(c)
		if err != nil {
			c.Error(err)
			return
		}
	}
}

func wrapResetPassword(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ResetPasswordRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ResetPassword(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapVerifyFactor(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.VerifyFactorRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.VerifyFactor(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}