package admin

import (
	"fmt"
	"invest_dairy/common"
	"invest_dairy/model/admin"
	"sort"
	"strconv"
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
	pool := new(admin.StockPool)
	pool.StockCode = dairy.StockCode
	err = pool.LoadLastByStockCode()
	if err != nil {
		common.Mlog.Errorf("load stock pool error: %s", err.Error())
		return common.CommonError()
	}
	if int64(pool.Number) < dairy.Number {
		return common.CommonError()
	}
	if dairy.Number != 0 {
		pool.Number += int(dairy.Number)
		pool.InPrice = (pool.InPrice * float64(pool.Number)) - (((dairy.Price - pool.InPrice) * float64(dairy.Number + int64(pool.Number))))
		err = pool.Update()
		if err != nil {
			common.Mlog.Errorf("update stock pool error: %s", err.Error)
			common.CommonError()
		}
	}
	return common.CommonSuccess()
}

type StockDetailVo struct {
	StockName   string
	StockCode   string
	Price       float64
	Number      int64
	Moeny       int64  
	MoneyRate   float64
	IncomeMoeny int64  
	DangerPrice []StockDetailCountVo
	StockCount  []StockDetailCountVo
	Dairy       []admin.StockDairy
}

type StockDetailCountVo struct {
	Day   string
	Value float64
}

func GetStockPoolDetail(filter admin.StockDetailCountFilter) *common.ResponseData {
	result := StockDetailVo{}
	stocks, err := admin.GetStockDetail(filter.StockCode)
	if err != nil {
		common.Mlog.Errorf("get stock detail error: %s", err.Error())
		common.CommonError()
	}
	dairy, err := admin.GetStockDairy(filter.StockCode)
	stockCount, err := admin.GetStockPriceChart(filter)
	if err != nil {
		common.Mlog.Errorf("get stock price error: %s", err.Error())
		common.CommonError()
	}
	stockChart := make([]admin.StockDetailCountVo, 0)
	var income int64
	if len(stockCount) > 0 {
		if len(stocks) <= 1 {
			stockChart = append(stockChart, admin.StockDetailCountVo{Day: stocks[0].CreateTime, Value: stocks[0].InPrice})
			stockChart = append(stockChart, stockCount...)
		} else if len(stocks) == 2 {
			stockChart = append(stockChart, admin.StockDetailCountVo{Day: stocks[0].CreateTime, Value: stocks[0].InPrice})
			stockChart = append(stockChart, stockCount...)
			stockChart = append(stockChart, admin.StockDetailCountVo{Day: stocks[0].CreateTime, Value: stocks[0].InPrice})
		} else {
			stockChart = append(stockChart, stockCount...)
			for _, s := range stocks {
				stockChart = append(stockChart, admin.StockDetailCountVo{Day: s.CreateTime, Value: s.InPrice})
			}
			sort.Sort(StockDetailSlice(stockChart))
		}
		income = int64((stockChart[len(stockChart)-1].Value - stocks[len(stocks)-1].InPrice) * float64(stocks[len(stocks)-1].Number))
	} else {
		income = 0
	}
	total, err := admin.QueryTotalMoney()
	if err != nil {
		common.Mlog.Errorf("query total moeny error: %s", err.Error())
		return common.CommonError()
	}
	stocksChartsData := make([]StockDetailCountVo, 0)
	for _, s := range stockChart {
		day := time.Unix(s.Day, 0).Format("2006-01-02")
		stocksChartsData = append(stocksChartsData, StockDetailCountVo{Day: day, Value: s.Value})
	}
	dangerPriceChartsData := make([]StockDetailCountVo, 0)
	for _, s := range stockChart {
		day := time.Unix(s.Day, 0).Format("2006-01-02")
		dangerPriceChartsData = append(dangerPriceChartsData, StockDetailCountVo{Day: day, Value: stocks[len(stocks)-1].DangerPrice})
	}
	money := int64(stocks[len(stocks)-1].InPrice * float64(stocks[len(stocks)-1].Number))
	result.StockName = stocks[len(stocks)-1].StockName
	result.StockCode = stocks[len(stocks)-1].StockCode
	result.Price = stocks[len(stocks)-1].InPrice
	result.Number = int64(stocks[len(stocks)-1].Number)
	result.Moeny = money
	result.IncomeMoeny = income
	result.MoneyRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(money) / float64(total) * float64(100)), 64)
	result.DangerPrice = dangerPriceChartsData
	result.StockCount = stocksChartsData
	result.Dairy = dairy
	return common.SetData(result)
}

type StockDetailSlice []admin.StockDetailCountVo

func (a StockDetailSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a StockDetailSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a StockDetailSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Day > a[i].Day
}

func QueryAllStock() *common.ResponseData {
	stock, err := admin.QueryAllStock()
	if err != nil {
		common.Mlog.Errorf("query all stock error: %s", err.Error())
		return common.CommonError()
	}
	return common.SetData(stock)
}
