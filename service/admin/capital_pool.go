package admin

import (
	"invest_dairy/common"
	"invest_dairy/model/admin"
	"time"
)

func AddMoney(money int) *common.ResponseData {
	pool := new(admin.CapitalPool)
	pool.Money = money
	pool.CreateTime = time.Now().Unix()
	err := pool.Insert()
	if err != nil {
		common.Mlog.Errorf("insert capital pool errir: %s", err.Error())
		return common.CommonError()
	}
	return common.CommonSuccess()
}

type CapitailPoolDetailVo struct {
	TotalMoeny    int64
	PositionMoney int64
	IncomeMoeny   int64
	RiskMoeny     int64
	Pool          []admin.CapitalPool
}

func GetCapitalPoolDetail() *common.ResponseData {
	result := CapitailPoolDetailVo{}
	totalMoney, err := admin.QueryTotalMoney()
	if err != nil {
		common.Mlog.Errorf("query total money error: %s", err.Error())
		common.CommonError()
	}
	positionMoney, err := admin.QueryPositionStockMeony()
	if err != nil {
		common.Mlog.Errorf("query position stock money error: %s", err.Error())
		return common.CommonError()
	}
	incomeMoeny, riskMoney, err := admin.QueryIncomeMoney()
	if err != nil {
		common.Mlog.Errorf("query income money error: %s", err.Error())
		return common.CommonError()
	}
	pool, err := admin.QueryCapitalPool()
	if err != nil {
		common.Mlog.Errorf("query capital pool error: %s", err.Error())
		return common.CommonError()
	}
	result.TotalMoeny = totalMoney
	result.PositionMoney = positionMoney
	result.IncomeMoeny = incomeMoeny
	result.RiskMoeny = riskMoney
	result.Pool = pool
	return common.SetData(result)
}
