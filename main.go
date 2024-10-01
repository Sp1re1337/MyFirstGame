package main

import (
	"log"
	"my_telegram_game/bot"
)

func main() {
	err := bot.StartBot()
	if err != nil {
		log.Panic(err)
	}
}