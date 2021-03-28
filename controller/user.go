package controller

import (
	"invest_dairy/common"
	"invest_dairy/service"
	"invest_dairy/util"

	"github.com/labstack/echo/v4"
	"github.com/yb7/echoswg"
)

func init() {
	e := echoswg.NewApiGroup(util.EchoInst, "用户API", "/api/user")
	e.POST("/login", login, "登录")
	e.POST("/loginOut", validUserToken, loginOut, "退出")
}

func login(req *struct{ Body service.UserLoginBo }) *common.ResponseData {
	return service.Login(req.Body)
}

func loginOut(user *service.UserInfoVo) *common.ResponseData {
	return service.LoginOut(user)
}

func validUserToken(ctx echo.Context) (*service.UserInfoVo, error) {
	return service.VerifyUserToken(ctx)
}
