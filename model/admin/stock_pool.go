package admin

import "invest_dairy/common"

type StockPool struct {
	Id          int `gorm:"primaryKey"`
	StockName   string
	StockCode   string
	InPrice     float64
	Number      int
	InLogic     string
	DangerPrice float64
	OutPrice    float64
	OutLogic    string
	CreateTime  int64
}

func (p *StockPool) Insert() error {
	return common.MySQL.Create(&p).Error
}

type StockDairy struct {
	Id         int `gorm:"primaryKey"`
	StockId    int
	Price      float64
	Logic      string
	CreateTime int64
}

func (p *StockDairy) Insert() error {
	return common.MySQL.Create(&p).Error
}

func QueryTotalMoney() (int64, error) {
	var total int64
	err := common.MySQL.Select("sum(money)").Table("i_capital_pool").Find(&total).Error
	return total, err
}
