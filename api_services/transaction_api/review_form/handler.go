package review_form

import (
	"fmt"
	"net/http"

	"github.com/yajw/review-bot/api_services/common"
	"github.com/yajw/review-bot/review_services/customer_review_service"
)

// GetReviewForm
// for showcase only
func GetReviewForm(w http.ResponseWriter, r *http.Request) {
	sceneKey := r.URL.Query().Get("scene")

	resp, err := customer_review_service.GetReviewService().GetReviewTemplate(sceneKey)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	resp.Template = fmt.Sprintf(resp.Template, r.URL.Query().Get("user_name"))
	common.JsonResp(w, resp)
}
