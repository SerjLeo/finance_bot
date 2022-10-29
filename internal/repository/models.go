package repository

type Transaction struct {
	Id       int64   `json:"id,omitempty"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency,omitempty"`
	Type     bool    `json:"type"`
}
