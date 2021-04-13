package admin

import (
	"database/sql"
	"invest_dairy/common"
)

type CapitalPool struct {
	Id         int `gorm:"primaryKey"`
	Money      int
	CreateTime int64
}

func (p *CapitalPool) Insert() error {
	return common.MySQL.Create(&p).Error
}

func QueryPositionStockMeony() (int64, error) {
	var money sql.NullInt64
	err := common.MySQL.Select("sum(in_price * number)").Table("i_stock_pool").
	Where("out_price = 0").Find(&money).Error
	return money.Int64, err
}

func QueryIncomeMoney() (int64, int64, error) {
	var result int64
	var dangerMoeny int64
	rows, err := common.MySQL.Select("number, in_price, out_price, danger_price, (select price from i_stock_dairy where stock_code = sp.id) newst_price").
	Table("i_stock_pool sp").Rows()
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		common.Mlog.Errorf("query income money erro: %s", err.Error())
		return result, dangerMoeny, err
	}
	for rows.Next() {
		var number float64
		var inPrice, outPrice, newstPrice, dangerPrice float64
		if err := rows.Scan(&number, &inPrice, &outPrice, &dangerPrice, &newstPrice); err != nil {
			common.Mlog.Errorf("scan income money error: %s", err.Error())
			return result, dangerMoeny, err
		}
		if outPrice > 0 {
			result += int64(number * (outPrice - inPrice))
		} else if (newstPrice > 0) {
			result += int64(number * (newstPrice - inPrice))
		} else {
			result += int64(number * inPrice)
		}
		if outPrice <= 0 {
			dangerMoeny += int64(number * inPrice) - int64(number * dangerPrice)
		}
	}
	return result, dangerMoeny, nil
}

func QueryCapitalPool() ([]CapitalPool, error) {
	result := make([]CapitalPool, 0)
	err := common.MySQL.Find(&result).Error
	if err != nil {
		common.Mlog.Errorf("find capital pool error: %s", err.Error())
	}
	return result, err
}