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
	router.POST("/login/auth", wrapAuth(si))
	router.POST("/logout", wrapLogout(si))
	router.POST("/login/verify-factor", wrapVerifyFactor(si))
}

func wrapAuth(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.AuthRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.Status(http.StatusBadRequest)
			c.Error(err)
			return
		}
		resp, err := si.Auth(c, &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
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
			c.Status(http.StatusInternalServerError)
			c.Error(err)
			return
		}
	}
}

func wrapVerifyFactor(si oas.Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req oas.VerifyFactorRequest
		if err := c.ShouldBind(&req.Body); err != nil {
			c.Status(http.StatusBadRequest)
			c.Error(err)
			return
		}
		resp, err := si.VerifyFactor(c, &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.Error(err)
			return
		}
		handler.NegotiateResponse(c, http.StatusOK, resp, []string{"application/json"})
	}
}
