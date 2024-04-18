package create_txn

type Request struct {
	UID     int64 `json:"uid,omitempty"`
	OrderID int64 `json:"order_id,omitempty"`
}
