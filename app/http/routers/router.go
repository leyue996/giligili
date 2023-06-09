package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(r *gin.Engine) {
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
}
