package service

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"wiliwili/conf"
	"wiliwili/serializer"
	"wiliwili/util"

	"github.com/google/uuid"
)

// UploadImageService 获得上传图片
type UploadImageService struct {
	NewFileName string
	File        *multipart.FileHeader
}

// Post 创建token
func (service *UploadImageService) Post(c *gin.Context) serializer.Response {
	if service.File == nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "请上传图片",
			Error:  "no file",
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.File.Filename)
	service.NewFileName = uuid.Must(uuid.NewRandom()).String() + ext
	util.LogD("newFileName", service.NewFileName)

	// 上传文件至指定目录
	updatePath := conf.Data.UpdatePath + "image/" + service.NewFileName
	util.LogD("updatePath", updatePath)

	if err := c.SaveUploadedFile(service.File, updatePath); err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "保存文件失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"file_url": updatePath,
		},
		Msg: "上传成功",
	}
}
