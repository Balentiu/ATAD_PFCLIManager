package budget

import (
	"ATAD_PFCLIManager/internal/core/transaction"
)

type Service struct {
	repo   Repository
	txRepo transaction.Repository
}

func NewService(r Repository, tr transaction.Repository) *Service {
	return &Service{repo: r, txRepo: tr}
}

func (s *Service) SetBudget(category string, amount float64) error {
	return s.repo.UpsertBudget(Budget{Category: category, Amount: amount})
}

func (s *Service) GetBudgetReport() ([]BudgetReport, error) {
	budgets, _ := s.repo.GetAllBudgets()
	txs, _ := s.txRepo.GetTransactions()

	spentMap := make(map[string]float64)
	for _, tx := range txs {
		if tx.Amount < 0 {
			spentMap[tx.Category] += (tx.Amount * -1)
		}
	}

	var reports []BudgetReport
	for _, b := range budgets {
		status := "OK"
		if spentMap[b.Category] > b.Amount {
			status = "!! DEPĂȘIT !!"
		}
		reports = append(reports, BudgetReport{
			Category: b.Category,
			Budget:   b.Amount,
			Spent:    spentMap[b.Category],
			Status:   status,
		})
	}
	return reports, nil
}
