package review_bot

import (
	"sync"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/libs/logger"
	"github.com/yajw/review-bot/review_services/review_bot/callback_handlers"
	"github.com/yajw/review-bot/review_services/review_bot/command_handlers"
)

const (
	// todo: store token in some secure place
	token = "7058887310:AAHRZyG_Fv5HVxz33f9gyI-0LBjc9esVTuA"
)

var (
	once sync.Once

	Bot *botapi.BotAPI
)

func Start() {
	once.Do(func() {
		start()
	})
}

func start() {
	var err error
	Bot, err = botapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	// remove in production env
	Bot.Debug = true

	logger.Info("bot account %s", Bot.Self.UserName)

	updateConfig := botapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		msg := update.Message
		if msg != nil {
			if msg.IsCommand() {
				if err := command_handlers.Handle(Bot, msg); err != nil {
					logger.Error("error: %+v", err)
				}
			} else {
				Bot.Send(botapi.NewMessage(msg.Chat.ID, "ha?"))
			}
		} else if update.CallbackQuery != nil {
			if err := callback_handlers.Handle(Bot, update.CallbackQuery); err != nil {
				logger.Error("error: %+v", err)
				continue
			}

			// let user know callback is handled
			// callback := botapi.NewCallback(update.CallbackQuery.ID, "")
			// if _, err := bot.Request(callback); err != nil {
			// logger.Error("callback request err:%+v", err)
			// continue
			// }
		}
	}
}
