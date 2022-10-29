package service

import "github.com/SerjLeo/finance_bot/internal/repository"

type Dependencies struct {
	Repo *repository.Repo
}

type IFinanceService interface {
	AddTransaction(amount float64) (*repository.Transaction, error)
	DeleteTransaction(id int64) error
	GetTransactions() []string
}

type Service struct {
	FinanceService IFinanceService
}
