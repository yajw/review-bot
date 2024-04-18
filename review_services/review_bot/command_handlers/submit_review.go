package command_handlers

import (
	"math/rand"
	"time"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/review_services/customer_review_service"
	"github.com/yajw/review-bot/review_services/customer_review_service/service"
)

func SubmitReview(bot *botapi.BotAPI, command string, chatID int64, userName string, msgID int) error {
	req := &service.SubmitReviewRequest{
		SceneKey:      "chat.order",
		SceneID:       rand.Int63(),
		UserID:        rand.Int63(),
		ReviewContent: "mock review content",
		SubmitTime:    time.Now(),
	}
	_, err := customer_review_service.GetReviewService().SubmitReview(req)
	if err != nil {
		return err
	}

	msg := botapi.NewMessage(chatID, "Thank you")
	msg.ReplyToMessageID = msgID

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}
