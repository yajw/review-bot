package order_service

type Order struct {
	ID  int64
	UID int64
}

func GetOrders(uid int64, limit int) ([]*Order, error) {
	panic("implement me")
}

func CreateOrder(uid int64, order *Order) error {
	panic("implement me")
}
