package command_handlers

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler func(bot *botapi.BotAPI, command string, chatID int64, userName string, msgID int) error

var handlers = map[string]handler{
	"mock_receive_product": MockReceiveProduct,
	"submit_review":        SubmitReview,
}

func Handle(bot *botapi.BotAPI, command string, chatID int64, userName string, msgID int) error {
	if handler, ok := handlers[command]; ok {
		return handler(bot, command, chatID, userName, msgID)
	}

	return nil

}
