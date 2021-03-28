package admin

import (
	"invest_dairy/common"
	"invest_dairy/model/admin"
	"time"
)

func AddStockPool(pool *admin.StockPool) *common.ResponseData {
	pool.CreateTime = time.Now().Unix()
	err := pool.Insert()
	if err != nil {
		common.Mlog.Errorf("insert stock pool error: %s", err.Error())
		return common.CommonError()
	}
	return common.CommonSuccess()
}

type StockDairyBo struct {
	admin.StockDairy
	IsSell bool
}

func AddStockDairy(dairy *StockDairyBo) *common.ResponseData {
	dairy.CreateTime = time.Now().Unix()
	err := dairy.Insert()
	if err != nil {
		common.Mlog.Errorf("insert stock dairy error: %s", err.Error())
		return common.CommonError()
	}
	return common.CommonSuccess()
}