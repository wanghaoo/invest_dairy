package controller

import (
	"invest_dairy/common"
	"invest_dairy/service"
	"invest_dairy/util"

	"github.com/labstack/echo/v4"
	"github.com/yb7/echoswg"
)

func init() {
	g := echoswg.NewApiGroup(util.EchoInst, "File", "/api/files")
	g.SetDescription("api for upload and download file")
	g.POST("/upload/image", validUserToken, uploadImage, "上传文件")
}

func uploadImage(c echo.Context, user *service.UserInfoVo) *common.ResponseData {
	return service.UploadImage(c)
}
