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
	router.POST("/mfa/bind", wrapBindMfa(si))
	router.POST("/mfa/bind-prepare", wrapBindMfaPrepare(si))
	router.GET("/captcha", wrapCaptcha(si))
	router.POST("/forget-pwd/begin", wrapForgetPwdBegin(si))
	router.POST("/forget-pwd/reset", wrapForgetPwdReset(si))
	router.POST("/forget-pwd/send-email", wrapForgetPwdSendEmail(si))
	router.POST("/forget-pwd/verify-email", wrapForgetPwdVerifyEmail(si))
	router.POST("/forget-pwd/verify-mfa", wrapForgetPwdVerifyMfa(si))
	router.POST("/login/auth", wrapLogin(si))
	router.POST("/logout", wrapLogout(si))
	router.POST("/login/reset-password", wrapResetPassword(si))
	router.POST("/mfa/unbind", wrapUnBindMfa(si))
	router.POST("/login/verify-factor", wrapVerifyFactor(si))
}

func wrapBindMfa(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.BindMfaRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.BindMfa(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapBindMfaPrepare(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		resp, err := si.BindMfaPrepare(c)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
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
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapForgetPwdBegin(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ForgetPwdBeginRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ForgetPwdBegin(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapForgetPwdReset(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ForgetPwdResetRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ForgetPwdReset(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapForgetPwdSendEmail(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ForgetPwdSendEmailRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ForgetPwdSendEmail(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapForgetPwdVerifyEmail(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ForgetPwdVerifyEmailRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ForgetPwdVerifyEmail(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}

func wrapForgetPwdVerifyMfa(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.ForgetPwdVerifyMfaRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.ForgetPwdVerifyMfa(c, &req)
		if err != nil {
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
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

func wrapUnBindMfa(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.UnBindMfaRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := si.UnBindMfa(c, &req)
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
