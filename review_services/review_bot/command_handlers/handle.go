package command_handlers

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler func(bot *botapi.BotAPI, msg *botapi.Message) error

var handlers = map[string]handler{
	"mock_receive_product": MockReceiveProduct,
}

func Handle(bot *botapi.BotAPI, msg *botapi.Message) error {
	if handler, ok := handlers[msg.Command()]; ok {
		return handler(bot, msg)
	}

	return nil

}
