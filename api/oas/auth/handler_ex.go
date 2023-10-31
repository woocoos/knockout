package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlersManual(router *gin.RouterGroup, service *ServerImpl) {
	router.GET("/login/mfaqr.png", func(c *gin.Context) {
		var req struct {
			UserID int    `form:"userId" binding:"required"`
			Secret string `form:"secret"`
		}
		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if req.UserID < 0 {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		bs, err := service.MfaQRCode(c, req.UserID, req.Secret)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "image/png", bs)
	})
}
