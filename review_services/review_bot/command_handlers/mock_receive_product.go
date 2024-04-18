package command_handlers

import (
	"fmt"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/review_services/customer_review_service"
)

func MockReceiveProduct(bot *botapi.BotAPI, command string, chatID int64, userName string, msgID int) error {
	tpl, err := customer_review_service.GetReviewService().GetReviewTemplate("chat.order")
	if err != nil {
		return err
	}

	text := fmt.Sprintf(tpl.Template, userName)
	msg := botapi.NewMessage(chatID, text)
	// msg.ReplyToMessageID = msgID

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}
