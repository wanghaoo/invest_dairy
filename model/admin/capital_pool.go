package admin

import "invest_dairy/common"

type CapitalPool struct {
	Id         int `gorm:"primaryKey"`
	Money      int
	CreateTime int64
}

func (p *CapitalPool) Insert() error {
	return common.MySQL.Create(&p).Error
}

func QueryPositionStockMeony() (int64, error) {
	var money int64
	err := common.MySQL.Select("sum(in_price * number)").Table("i_stock_pool").
	Where("out_price = 0").Find(&money).Error
	return money, err
}

