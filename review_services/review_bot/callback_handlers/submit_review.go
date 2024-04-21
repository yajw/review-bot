package callback_handlers

import (
	"net/url"
	"strconv"
	"time"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/review_services/customer_review_service"
	"github.com/yajw/review-bot/review_services/customer_review_service/service"
	"github.com/yajw/review-bot/review_services/review_bot/mock_user_mapping"
)

func SubmitReview(bot *botapi.BotAPI, msg *botapi.CallbackQuery, q *url.URL) error {
	content := q.Query().Get("c")
	sceneID, _ := strconv.ParseInt(q.Query().Get("sid"), 10, 64)

	println(content, sceneID)

	req := &service.SubmitReviewRequest{
		SceneKey:      "chat.order",
		SceneID:       sceneID,
		UserID:        mock_user_mapping.UIDMapping[msg.From.UserName],
		ReviewContent: string(content),
		SubmitTime:    time.Now(),
	}
	_, err := customer_review_service.GetReviewService().SubmitReview(req)
	if err != nil {
		return err
	}

	reply := botapi.NewMessage(msg.From.ID, "Thank you for your review")
	if _, err := bot.Send(reply); err != nil {
		return err
	}

	return nil
}
