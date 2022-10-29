package service

type Dependencies struct {
}

type Transaction struct {
	Id       int64   `json:"id,omitempty"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency,omitempty"`
	Type     bool    `json:"type"`
}

type IFinanceService interface {
	AddTransaction(amount float64) (Transaction, error)
	DeleteTransaction(id int64) error
}

type Service struct {
	FinanceService IFinanceService
}
