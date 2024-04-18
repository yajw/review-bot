package create_txn

import (
	"encoding/json"
	"net/http"

	"github.com/yajw/review-bot/api_services/common"
	"github.com/yajw/review-bot/order_services/order_service"
)

func CreateTxn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req Request
	err := decoder.Decode(&req)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	err = order_service.CreateOrder(req.UID, &order_service.Order{
		ID:  req.UID,
		UID: req.OrderID,
	})

	common.HandleError(err, w)
	return
}
