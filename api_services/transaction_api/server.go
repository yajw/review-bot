package transaction_api

import (
	"log"
	"net/http"
	"sync"

	"github.com/yajw/review-bot/api_services/transaction_api/create_txn"
	"github.com/yajw/review-bot/api_services/transaction_api/review_form"
	"github.com/yajw/review-bot/api_services/transaction_api/submit_txn_review"
	"github.com/yajw/review-bot/api_services/transaction_api/transaction_list"
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

		log.Fatal(http.ListenAndServe(":8080", nil))
	})
}
