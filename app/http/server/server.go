package server

import (
	"giligili/app"
	"giligili/app/http/routers"
	"github.com/gin-gonic/gin"
)

func NewServer() error {
	r := gin.Default()

	routers.NewRouter(r)

	return r.Run(app.Config.HttpService.Lister)
}
