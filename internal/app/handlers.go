package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r *Router) HandleError(chatId int64, err error) {
	fmt.Println(err)
	r.api.Send(tgbotapi.NewMessage(chatId, "Error while performing operation"))
}

func (r *Router) Help(chatId int64) {
	msg := `
/help - show command list
/add - add transaction. Format {+/-}1000{USD} {Message} where optional parameters in {}
/balance - show balance in selected currency
/list - show a list of recent transactions
    `
	reply := tgbotapi.NewMessage(chatId, msg)
	r.api.Send(reply)
}

func (r *Router) Add(msg *tgbotapi.Message) error {
	amount, err := strconv.ParseFloat(msg.CommandArguments(), 64)
	if err != nil {
		return err
	}
	transaction, err := r.service.FinanceService.AddTransaction(amount)
	if err != nil {
		return err
	}
	r.api.Send(tgbotapi.NewMessage(msg.Chat.ID, strconv.Itoa(int(transaction.Amount))))
	return nil
}

func (r *Router) List(msg *tgbotapi.Message) {
	list := r.service.FinanceService.GetTransactions()
	result := ""
	for _, item := range list {
		result = result + "\n" + item
	}
	r.api.Send(tgbotapi.NewMessage(msg.Chat.ID, result))
}
