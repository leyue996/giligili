package server

import (
	"giligili/app"
	"giligili/app/http/routers"
	"github.com/gin-gonic/gin"
)

func NewServer() error {
	r := routers.NewRouter()
	if app.Config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return r.Run(app.Config.HttpService.Lister)
}
