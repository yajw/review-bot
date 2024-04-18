package main

import (
	"github.com/yajw/review-bot/api_services/transaction_api"
	"github.com/yajw/review-bot/infra/mysql"
	"github.com/yajw/review-bot/review_services/customer_review_service"
	"github.com/yajw/review-bot/review_services/review_bot"
)

func main() {
	mysql.StartServer()
	customer_review_service.Start()
	transaction_api.Start()
	review_bot.Start()
}
