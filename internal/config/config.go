package config

import "os"

type Config struct {
	BotToken  string
	Port      string
	PublicURL string
}

func Load() *Config {
	return &Config{
		BotToken:  os.Getenv("BOT_TOKEN"),
		Port:      os.Getenv("PORT"),
		PublicURL: os.Getenv("PUBLIC_URL"),
	}
}
