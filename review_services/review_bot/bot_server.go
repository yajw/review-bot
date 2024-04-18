package review_bot

import (
	"sync"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/libs/logger"
	"github.com/yajw/review-bot/review_services/review_bot/command_handlers"
)

const (
	// todo: store token in some secure place
	token = "7058887310:AAHRZyG_Fv5HVxz33f9gyI-0LBjc9esVTuA"
)

var (
	once sync.Once
)

func Start() {
	once.Do(func() {
		start()
	})
}

func start() {
	bot, err := botapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	// remove in production env
	bot.Debug = true

	updateConfig := botapi.NewUpdate(0)
	updateConfig.Timeout = 3

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		msg := update.Message
		if msg == nil {
			continue
		}

		if msg.IsCommand() {
			if err := command_handlers.Handle(bot, msg.Text, msg.Chat.ID, msg.Chat.UserName, msg.MessageID); err != nil {
				logger.Error("error: %+v", err)
			}
		}
	}
}
