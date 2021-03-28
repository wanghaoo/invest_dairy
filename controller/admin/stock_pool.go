package admin

import (
	"invest_dairy/common"
	"invest_dairy/util"
	"invest_dairy/service/admin"
	ma "invest_dairy/model/admin"
	"github.com/yb7/echoswg"
)

func init() {
	e := echoswg.NewApiGroup(util.EchoInst, "股票池管理", "/admin/stock")
	e.GET("/:StockCode", validateSysUser, getStockPoolDetail, "获取资金池详细信息")
	e.GET("", validateSysUser, queryAllStock, "查询所有股票")
	e.POST("", validateSysUser, addStock, "添加股票")
	e.POST("/dairy", validateSysUser, addStockDairy, "添加股票日记")
}

func getStockPoolDetail(req *struct{ma.StockDetailCountFilter}) *common.ResponseData {
	return admin.GetStockPoolDetail(req.StockDetailCountFilter)
}

func queryAllStock() *common.ResponseData {
	return admin.QueryAllStock()
}

func addStock(req *struct{Body *ma.StockPool}) *common.ResponseData {
	return admin.AddStockPool(req.Body)
}

func addStockDairy(req *struct{Body *admin.StockDairyBo}) *common.ResponseData {
	return admin.AddStockDairy(req.Body)
}