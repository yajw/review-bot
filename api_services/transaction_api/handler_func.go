package transaction_api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yajw/review-bot/review_services/customer_review_service"
	"github.com/yajw/review-bot/review_services/customer_review_service/service"
)

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
