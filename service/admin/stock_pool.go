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
	err := dairy.StockDairy.Insert()
	if err != nil {
		common.Mlog.Errorf("insert stock dairy error: %s", err.Error())
		return common.CommonError()
	}
	if dairy.IsSell {
		pool := new(admin.StockPool)
		pool.Id = dairy.StockId
		pool.OutPrice = dairy.Price
		pool.OutLogic = dairy.Logic
		err := pool.Update()
		if err != nil {
			common.Mlog.Errorf("update stock pool error: %s", err.Error)
			common.CommonError()
		}
	}
	return common.CommonSuccess()
}

func GetStockPoolDetail(filter admin.StockDetailCountFilter) *common.ResponseData {
	stocks, err := admin.GetStockDetail(filter.StockCode)
	if err != nil {
		common.Mlog.Errorf("get stock detail error: %s", err.Error())
		common.CommonError()
	}
	stockCount, err := admin.GetStockPriceChart(filter)
	if err != nil {
		common.Mlog.Errorf("get stock price error: %s", err.Error())
		common.CommonError()
	}
	
}

func QueryAllStock() *common.ResponseData {
	stock, err := admin.QueryAllStock()
	if err != nil {
		common.Mlog.Errorf("query all stock error: %s", err.Error())
		return common.CommonError()
	}
	return common.SetData(stock)
}