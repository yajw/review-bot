package transaction_api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yajw/review-bot/rpc_services/customer_review_service"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/service"
)

func TransactionList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uid, _ := strconv.ParseInt(r.URL.Query().Get("uid"), 10, 64)
	itemIDStrs, _ := r.URL.Query()["item_ids"]

	itemIDs := make([]int64, 0, len(itemIDStrs))
	for _, s := range itemIDStrs {
		itemID, _ := strconv.ParseInt(s, 10, 64)
		itemIDs = append(itemIDs, itemID)
	}

	resp, err := customer_review_service.GetReviewService().GetUserRates(
		&service.GetUserRatesRequest{
			UserID:     uid,
			ItemIDList: itemIDs,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

func TransactionReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uid, _ := strconv.ParseInt(r.FormValue("uid"), 10, 64)
	itemID, _ := strconv.ParseInt(r.FormValue("item_id"), 10, 64)
	comment := r.FormValue("comment")
	star, _ := strconv.ParseInt(r.FormValue("star"), 10, 64)

	resp, err := customer_review_service.GetReviewService().Rate(
		&service.RateRequest{
			UserID:  uid,
			ItemID:  itemID,
			Star:    int(star),
			Comment: comment,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
