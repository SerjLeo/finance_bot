package app

import (
	"fmt"
	"github.com/SerjLeo/finance_bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log"
)

type Router struct {
	api     *tgbotapi.BotAPI
	service *service.Service
}

func NewRouter(service *service.Service, token string) (*Router, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Router{
		bot, service,
	}, nil
}

func (r *Router) Listen() error {
	r.api.Debug = true

	log.Printf("Authorized on account %s", r.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := r.api.GetUpdatesChan(u)

	for update := range updates {
		r.HandleUpdate(update)
	}

	return nil
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {

	switch callback.Message.Command() {
	case "education":
		break
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callback.Message.Command())
	}
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		r.HandleError(msg.Chat.ID, errors.New("message is not a command"))
		return
	}
	fmt.Println(msg.Command())
	switch msg.Command() {
	case "help":
		r.Help(msg.Chat.ID)
	case "list":
		r.List(msg)
	case "add":
		err := r.Add(msg)
		if err != nil {
			r.HandleError(msg.Chat.ID, err)
		}
	case "test":
		r.api.Send(tgbotapi.NewMessage(msg.Chat.ID, "Test message!"))
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", msg.Command())
	}
}
