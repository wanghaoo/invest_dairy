package model

import (
	"invest_dairy/common"
	"time"
)

type Order struct {
	Id           int `gorm:"primaryKey"`
	Uid          int
	Price        float64
	OrderId      string
	QueryOrderId string
	Status       string
	PayMethod    string
	PayChannel   string
	MerchantKey  string
	CreateTime   int64
	ModifyTime   int64
	Remark       string
}

func (order *Order) GetByOrderId() error {
	err := common.MySQL.Last(&order, "order_id = ?", order.OrderId).Error
	if err != nil {
		common.Mlog.Errorf("query order by id error: %s", err.Error())
	}
	return err
}

func (order *Order) Insert() error {
	err := common.MySQL.Create(order).Error
	if err != nil {
		common.Mlog.Errorf("insert order error: %s", err.Error())
	}
	return err
}

func (order *Order) Update() error {
	err := common.MySQL.Save(order).Error
	if err != nil {
		common.Mlog.Errorf("update order error: %s", err.Error())
	}
	return err
}

func QueryAllPendingOrder() ([]Order, error) {
	result := make([]Order, 0)
	createTime := time.Now().Add(-24 * time.Hour).Unix()
	err := common.MySQL.Where("pay_method = 'Payment' and status = ? and create_time > ?", common.ORDER_STATUS_PENDING, createTime).Find(&result).Error
	if err != nil {
		common.Mlog.Errorf("query order errro: %s", err.Error())
	}
	return result, err
}
