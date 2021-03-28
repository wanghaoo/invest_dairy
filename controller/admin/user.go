package admin

import (
	"invest_dairy/common"
	"invest_dairy/model/admin"
	ms "invest_dairy/service/admin"
	"invest_dairy/util"

	"github.com/yb7/echoswg"
)

func init() {
	e := echoswg.NewApiGroup(util.EchoInst, "Admin/系统管理", "/admin/user")
	e.POST("/login", login, "登录")
	e.POST("/logout", logout, "登出")
	e.GET("", validateSysUser, findUsers, "查询系统用户")
	e.GET("/:Id", validateSysUser, findUserDetail, "用户详情")
	e.POST("", validateSysUser, AddOrUpdateSysUser, "添加/修改系统用户")
	e.PUT("", validateSysUser, updatePassword, "修改密码")
}

func login(req *struct{ Body admin.LoginBo }) *common.ResponseData {
	return ms.Login(req.Body)
}

func logout(req *struct{ Token string }) *common.ResponseData {
	return ms.Logout(req.Token)
}

func validateSysUser(req *struct{ Token string }) (*ms.UserInfoVo, error) {
	return ms.ValidateSysUser(req.Token)
}

func findUsers(req *struct{ admin.SysUserFilter }) *common.ResponseData {
	return ms.FindSysUser(req.SysUserFilter)
}

func AddOrUpdateSysUser(req *struct{ Body ms.SysUserBo }) *common.ResponseData {
	return ms.AddOrUpdateSysUser(req.Body)
}

func findUserDetail(req *struct{ Id int }) *common.ResponseData {
	return ms.FindUserDetail(req.Id)
}

func updatePassword(req *struct{ Body ms.ModifySysUserBo }, vo *ms.UserInfoVo) *common.ResponseData {
	return ms.UpdatePassword(req.Body, vo)
}
