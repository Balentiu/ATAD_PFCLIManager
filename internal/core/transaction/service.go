package transaction

import (
    "fmt"
    "time"
)

type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) Add(date time.Time, desc string, amount float64, category string) error {
    if desc == "" {
        return fmt.Errorf("descrierea nu poate fi goală")
    }
    if amount == 0 {
        return fmt.Errorf("suma nu poate fi zero")
    }
    if category == "" {
        category = "General"
    }

    tx := Transaction{
        Date:        date,
        Description: desc,
        Amount:      amount,
        Category:    category,
    }

    return s.repo.AddTransaction(tx)
}

func (s *Service) List() ([]Transaction, error) {
    return s.repo.GetTransactions()
}