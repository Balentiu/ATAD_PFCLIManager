package transaction

import (
	"fmt"
	"regexp"
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

func (s *Service) AutoCategorize(desc string) string {
	rules := map[string]string{
		"(?i)supermarket|lidl|kaufland|mega": "Mancare",
		"(?i)omv|petrom|mol|benzina":         "Transport",
		"(?i)netflix|hbo|cinema":             "Divertisment",
		"(?i)salariu|bonus":                  "Venituri",
	}

	for pattern, category := range rules {
		matched, _ := regexp.MatchString(pattern, desc)
		if matched {
			return category
		}
	}
	return "Altele"
}

func (s *Service) ImportTransactions(txs []Transaction) (int, error) {
	count := 0
	for _, tx := range txs {
		tx.Category = s.AutoCategorize(tx.Description)

		err := s.repo.AddTransaction(tx)
		if err == nil {
			count++
		}
	}
	return count, nil
}

func (s *Service) Search(query string) ([]Transaction, error) {
	if query == "" {
		return s.List()
	}
	return s.repo.SearchTransactions(query)
}
