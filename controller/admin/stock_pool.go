func package

func init() {
	e := echoswg.NewApiGroup(util.EchoInst, "股票池管理", "/admin/stock")
	e.GET("", validateSysUser, getStockPoolDetail, "获取资金池详细信息")
	e.POST("", validateSysUser, addStock, "添加股票")
	e.POST("/dairy", validateSysUser, addStockDairy, "添加股票日记")
}

func getStockPoolDetail() *common.ResponseData {

}

func addStock(req *struct{Body *admin.StockPool}) *common.ResponseData {
	return admin.AddStockPool(req.Body)
}

func addStockDairy(req *struct{Body *admin.StockDairyBo}) *common.ResponseData {
	return admin.AddStockDairy(req.Body)
}