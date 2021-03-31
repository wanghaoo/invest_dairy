package admin

import (
	"invest_dairy/common"
	"invest_dairy/service/admin"
	"invest_dairy/util"

	"github.com/yb7/echoswg"
)

func init() {
	e := echoswg.NewApiGroup(util.EchoInst, "资金池管理", "/admin/capital")
	e.GET("", validateSysUser, getCapitalPoolDetail, "获取资金池详细信息")
	e.POST("", validateSysUser, addMoney, "添加金额")
}

func getCapitalPoolDetail() *common.ResponseData {
	return admin.GetCapitalPoolDetail()
}

func addMoney(req *struct{Body struct{Money int}}) *common.ResponseData {
	return admin.AddMoney(req.Body.Money)
}
