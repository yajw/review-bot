package main

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "7058887310:AAHRZyG_Fv5HVxz33f9gyI-0LBjc9esVTuA"

func main() {
	bot, err := botapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := botapi.NewUpdate(0)
	updateConfig.Timeout = 3

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := botapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
