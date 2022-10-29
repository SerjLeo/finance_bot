package service

import "github.com/SerjLeo/finance_bot/internal/repository"

type FinanceService struct {
	db *repository.Repo
}

func (s *FinanceService) AddTransaction(amount float64) (*repository.Transaction, error) {
	return s.db.AddTransaction(&repository.Transaction{
		Amount:   amount,
		Type:     true,
		Currency: "USD",
		Id:       0,
	})
}

func (s *FinanceService) DeleteTransaction(id int64) error {
	return nil
}

func (s *FinanceService) GetTransactions() []string {
	return s.db.GetTransactionsList()
}

func NewFinanceService(db *repository.Repo) *FinanceService {
	return &FinanceService{
		db,
	}
}
