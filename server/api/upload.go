package api

import (
	"github.com/gin-gonic/gin"
	"wiliwili/service"
	"wiliwili/util"
)

// UploadToken 上传授权
func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UploadVideo 上传图片
func UploadImage(c *gin.Context) {
	service := service.UploadImageService{}
	var err error
	if service.File, err = c.FormFile("file"); err != nil {
		c.JSON(200, ErrorResponse(err))
	}

	c.JSON(200, service.Post(c))
}

// UploadVideo 上传文件
func UploadVideo(c *gin.Context) {
	util.LogD("UploadVideo")
	service := service.UploadVideoService{}
	var err error
	if service.File, err = c.FormFile("file"); err != nil {
		c.JSON(200, ErrorResponse(err))
	}

	c.JSON(200, service.Post(c))
}
