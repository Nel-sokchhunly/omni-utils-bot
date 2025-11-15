package main

import (
	"log"
	"utils-bot/internal/bot"
	"utils-bot/internal/config"
)

func main() {
	cfg := config.Load()
	if err := bot.RunWebhook(cfg); err != nil {
		log.Fatalf("Failed to run bot: %v", err)
	}
}
