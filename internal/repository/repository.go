package repository

import (
	"fmt"
)

type Repo struct {
	db []string
}

func (r *Repo) push(s string) {
	r.db = append(r.db, s)
}

func NewRepo() *Repo {
	return &Repo{db: []string{}}
}

func (r *Repo) AddTransaction(t *Transaction) (*Transaction, error) {
	r.push(fmt.Sprintf("%f %s", t.Amount, t.Currency))
	return t, nil
}

func (r *Repo) GetTransactionsList() []string {
	return r.db
}
