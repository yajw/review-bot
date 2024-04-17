package transaction_api

import (
	"log"
	"net/http"
	"sync"
)

var (
	once sync.Once
)

func Start() {
	once.Do(func() {
		http.HandleFunc("/user/transaction/list", TransactionList)
		http.HandleFunc("/user/transaction/review", TransactionReview)

		log.Fatal(http.ListenAndServe(":8080", nil))
	})
}
