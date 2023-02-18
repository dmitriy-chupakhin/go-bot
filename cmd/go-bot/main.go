package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Command() == "help" {
				helpCommand(update.Message, bot)
				continue
			}

			defaultMessage(update.Message, bot)
		}
	}
}

func defaultMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)
}

func helpCommand(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You need help? Hmm..")

	bot.Send(msg)
}
