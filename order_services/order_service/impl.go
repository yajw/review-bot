package order_service

import (
	"time"

	"github.com/yajw/review-bot/infra/mysql"
)

type UserOrder struct {
	ID         int64
	UID        int64
	CreateTime time.Time
}

func (*UserOrder) TableName() string {
	return "user_order"
}

func GetOrders(uid int64, limit int) ([]*UserOrder, error) {
	db, err := mysql.GetConnection()
	if err != nil {
		return nil, err
	}

	var orders []*UserOrder
	err = db.Table("user_order").Where("uid = ?", uid).Limit(limit).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func CreateOrder(uid int64, order *UserOrder) (int64, error) {
	db, err := mysql.GetConnection()
	if err != nil {
		return 0, err
	}

	if err := db.Create(order).Error; err != nil {
		return 0, err
	}

	return order.ID, nil
}
