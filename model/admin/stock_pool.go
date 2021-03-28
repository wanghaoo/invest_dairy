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

func (p *StockPool) Update() error {
	err := common.MySQL.Update("out_price", p.OutPrice).Update("out_logic", p.OutLogic).
		Table("i_stock_pool").Where("id = ?", p.Id).Error
	return err
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

type Stock struct {
	StockName string
	StockCode string
}

func QueryAllStock() ([]Stock, error) {
	result := make([]Stock, 0)
	err := common.MySQL.Select("stock_name, stock_code").Model("i_stock_pool").Order("create_time desc").Find(&result).Error
	return result, err
}

type StockDetailCountFilter struct {
	StockCode string
	BeginDate int64
	EndDate int64
}

type StockDetailCountVo struct {
	Day   int64
	Value float64
}

func GetStockPriceChart(filter StockDetailCountFilter) ([]StockDetailCountVo, error) {
	result := make([]StockDetailCountVo, 0)
	err := common.MySQL.Select("create_time as day, price as value").
	Table("i_stock_dairy").Where("stock_code = ?", filter.StockCode).Order("create_time asc").Find(&result).Error
	return result, err
}

func GetStockDetail(stockCode string) ([]StockPool, error) {
	result := make([]StockPool, 0)
	err := common.MySQL.Table("i_stock_pool").Where("stock_code = ?", stockCode).
	Order("create_time asc").Find(&result).Error
	return result, err
} 

