package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TG_BOT_API_KEY")
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized as @%s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s",update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
