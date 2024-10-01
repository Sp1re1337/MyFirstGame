package game

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func HandleGame(input string, secretNumber int, chatID int64) tgbotapi.MessageConfig {
	guess, err := strconv.Atoi(input)
	if err != nil {
		return tgbotapi.NewMessage(chatID, "Будь ласка, введіть число від 1 до 10.")
	}

	if guess == secretNumber {
		return tgbotapi.NewMessage(chatID, fmt.Sprintf("Вітаємо! Ви вгадали число %d!", secretNumber))
	} else {
		return tgbotapi.NewMessage(chatID, "Невірно. Спробуйте ще раз!")
	}
}