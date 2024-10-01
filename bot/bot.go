package bot

import (
	"log"
	"math/rand"
	"time"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"my_telegram_game/game"

)

var secretNumber int

func StartBot() error {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Авторизовано як %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	rand.Seed(time.Now().Unix())
  secretNumber = rand.Intn(10) + 1

	for update := range updates {
		if update.Message != nil {
			msg := game.HandleGame(update.Message.Text, secretNumber, update.Message.Chat.ID)
			bot.Send(msg)
		}
	}

	return nil
}