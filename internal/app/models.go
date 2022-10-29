package app

import (
	"github.com/SerjLeo/finance_bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type dbConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
}

type botConfig struct {
	Token string `yaml:"token"`
}

type Config struct {
	Database dbConfig  `yaml:"database"`
	Bot      botConfig `yaml:"bot"`
}

type Router struct {
	api     *tgbotapi.BotAPI
	service *service.Service
}
