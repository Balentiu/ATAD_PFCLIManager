package report

import (
	"ATAD_PFCLIManager/internal/core/transaction"
	"fmt"
	"strings"
)

type Service struct {
	txRepo transaction.Repository
}

func NewService(tr transaction.Repository) *Service {
	return &Service{txRepo: tr}
}

func (s *Service) GetCategoryBreakdown() (map[string]float64, error) {
	txs, err := s.txRepo.GetTransactions()
	if err != nil {
		return nil, err
	}

	totalSpent := 0.0
	breakdown := make(map[string]float64)

	for _, tx := range txs {
		if tx.Amount < 0 {
			amount := tx.Amount * -1
			breakdown[tx.Category] += amount
			totalSpent += amount
		}
	}

	return breakdown, nil
}

func (s *Service) GenerateBarChart(label string, value float64, max float64) string {
	const barWidth = 20
	percentage := (value / max) * 100
	fillWidth := int((value / max) * barWidth)
	if fillWidth > barWidth {
		fillWidth = barWidth
	}

	bar := strings.Repeat("█", fillWidth) + strings.Repeat("░", barWidth-fillWidth)
	return fmt.Sprintf("%-12s | %s | %.2f RON (%.1f%%)", label, bar, value, percentage)
}
