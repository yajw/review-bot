package transaction_list

import (
	"time"
)

type TxnData struct {
	OrderInfo  *OrderInfo  `json:"order_info,omitempty"`
	ReviewInfo *ReviewInfo `json:"review_info,omitempty"`
}

type OrderInfo struct {
	OrderID int64 `json:"order_id,omitempty"`
}

type ReviewInfo struct {
	ID            int64     `json:"id,omitempty"`
	UserID        int64     `json:"user_id,omitempty"`
	ReviewContent string    `json:"review_content,omitempty"`
	SubmitTime    time.Time `json:"submit_time"`
}

type TxnListResponse struct {
	TxnList []*TxnData `json:"txn_list,omitempty"`
}
