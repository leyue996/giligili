package api

import (
	"giligili/app/http/pkg/e"
	"giligili/app/http/pkg/util"
	"giligili/app/http/service"
	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context) {
	createVideoService := service.CreateVideoService{}
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&createVideoService); err != nil {
		c.JSON(e.Error, ErrorResponse(err))
	} else {
		res := createVideoService.Create(c.Request.Context(), claims.Id)
		c.JSON(e.Success, res)
	}
}

func ShowVideo(c *gin.Context) {
	showVideoService := service.ShowVideoService{}
	if err := c.ShouldBind(&showVideoService); err != nil {
		c.JSON(e.Error, ErrorResponse(err))
	} else {
		res := showVideoService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(e.Success, res)
	}
}
func ListVideo(c *gin.Context) {
	listVideosService := service.ListVideosService{}
	if err := c.ShouldBind(&listVideosService); err != nil {
		c.JSON(e.Error, ErrorResponse(err))
	} else {
		res := listVideosService.List(c.Request.Context())
		c.JSON(e.Success, res)
	}
}
func UpdateVideo(c *gin.Context) {
	updateVideoService := service.UpdateVideoService{}
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&updateVideoService); err != nil {
		c.JSON(e.Error, ErrorResponse(err))
	} else {
		res := updateVideoService.Update(c.Request.Context(), claims.Id, c.Param("id"))
		c.JSON(e.Success, res)
	}
}
func DeleteVideo(c *gin.Context) {
	deleteVideoService := service.DeleteVideoService{}
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&deleteVideoService); err != nil {
		c.JSON(e.Error, ErrorResponse(err))
	} else {
		res := deleteVideoService.Delete(c.Request.Context(), claims.Id, c.Param("id"))
		c.JSON(e.Success, res)
	}
}
