package service

type FinanceService struct {
}

func (s *FinanceService) AddTransaction(amount float64) (Transaction, error) {
	return Transaction{
		Amount:   amount,
		Type:     true,
		Currency: "USD",
		Id:       0,
	}, nil
}

func (s *FinanceService) DeleteTransaction(id int64) error {
	return nil
}

func NewFinanceService() *FinanceService {
	return &FinanceService{}
}
