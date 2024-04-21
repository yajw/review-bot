package submit_txn_review

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yajw/review-bot/api_services/common"
	"github.com/yajw/review-bot/review_services/customer_review_service"
	"github.com/yajw/review-bot/review_services/customer_review_service/service"
)

func SubmitTxnReviewHandler(w http.ResponseWriter, r *http.Request) {
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

	var resp *service.SubmitReviewResponse
	resp, err = customer_review_service.GetReviewService().SubmitReview(&service.SubmitReviewRequest{
		SceneKey:      req.SceneKey,
		SceneID:       req.SceneID,
		UserID:        req.UID,
		ReviewContent: req.ReviewContent,
		SubmitTime:    time.Now(),
		ExtraAttrs:    req.ExtraAttrs,
	})

	if err != nil {
		common.HandleError(err, w)
		return
	}

	common.JsonResp(w, resp)
}
