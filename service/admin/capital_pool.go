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