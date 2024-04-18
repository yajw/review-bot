package transaction_list

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/yajw/review-bot/api_services/common"
	"github.com/yajw/review-bot/order_services/order_service"
	"github.com/yajw/review-bot/review_services/customer_review_service"
)

const (
	sceneKey = "transaction_list"
)

func TransactionList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uid, err := strconv.ParseInt(r.URL.Query().Get("uid"), 10, 64)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	if uid <= 0 {
		common.HandleError(errors.New("invalid uid"), w)
		return
	}

	orders, err := order_service.GetOrders(uid, 1000)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	orderIDs := make([]int64, 0, len(orders))
	for _, order := range orders {
		if order == nil {
			continue
		}
		orderIDs = append(orderIDs, order.ID)
	}

	reviewMap, err := getOrderReviews(uid, orderIDs)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	txnList := make([]*TxnData, 0, len(orderIDs))
	for _, order := range orders {
		if order == nil {
			continue
		}

		txnList = append(txnList, &TxnData{
			OrderInfo:  convert(order),
			ReviewInfo: reviewMap[order.ID],
		})
	}

	resp := &TxnListResponse{TxnList: txnList}
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

func convert(order *order_service.Order) *OrderInfo {
	if order == nil {
		return nil
	}

	return &OrderInfo{
		OrderID: order.ID,
	}
}

func getOrderReviews(uid int64, orderIDs []int64) (map[int64]*ReviewInfo, error) {
	if len(orderIDs) == 0 {
		return nil, nil
	}

	reviews, err := customer_review_service.GetReviewService().GetUserReviews(uid, sceneKey, orderIDs)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]*ReviewInfo, len(orderIDs))
	for _, review := range reviews {
		res[review.SceneID] = &ReviewInfo{
			ID:            review.ID,
			UserID:        review.UserID,
			ReviewContent: review.ReviewContent,
			SubmitTime:    review.SubmitTime,
		}
	}
	return res, nil
}
