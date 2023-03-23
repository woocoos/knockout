package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlersManual(router *gin.Engine, service *Service) {
	router.GET("/login/mfaqr.png", func(c *gin.Context) {
		var req struct {
			UserID int `form:"userId" binding:"required"`
		}
		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if req.UserID < 0 {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		bs, err := service.MfaQRCode(c, req.UserID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "image/png", bs)
	})
}
