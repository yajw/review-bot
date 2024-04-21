package transaction_api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/yajw/review-bot/api_services/common"
	"github.com/yajw/review-bot/api_services/transaction_api/create_txn"
	"github.com/yajw/review-bot/api_services/transaction_api/review_form"
	"github.com/yajw/review-bot/api_services/transaction_api/submit_txn_review"
	"github.com/yajw/review-bot/api_services/transaction_api/transaction_list"
	"github.com/yajw/review-bot/review_services/review_bot"
	"github.com/yajw/review-bot/review_services/review_bot/command_handlers"
)

var (
	once sync.Once
)

func Start() {
	once.Do(func() {
		http.HandleFunc("/transaction/list", transaction_list.TransactionList)
		http.HandleFunc("/transaction/review/submit", submit_txn_review.SubmitTxnReviewHandler)
		http.HandleFunc("/transaction/review/form", review_form.GetReviewForm)

		// create order mock
		http.HandleFunc("/transaction/create", create_txn.CreateTxn)
		// order callback mock
		http.HandleFunc("/product/receive", mockReceiveProduct)

		log.Fatal(http.ListenAndServe(":8080", nil))
	})
}

func mockReceiveProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	type Req struct {
		ChatID   int64  `json:"chat_id"`
		SceneID  int64  `json:"scene_id"`
		UserName string `json:"user_name"`
	}
	var req Req
	err := decoder.Decode(&req)
	if err != nil {
		common.HandleError(err, w)
		return
	}

	command_handlers.ReceiveProduct(review_bot.Bot, req.ChatID, req.UserName, "chat.order", req.SceneID)

	common.JsonResp(w, map[string]interface{}{"result": "success"})
}
