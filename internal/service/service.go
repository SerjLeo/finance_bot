package service

func InitService(deps *Dependencies) *Service {
	return &Service{
		FinanceService: NewFinanceService(deps.Repo),
	}
}
