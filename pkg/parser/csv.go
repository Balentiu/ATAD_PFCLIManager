package parser

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"ATAD_PFCLIManager/internal/core/transaction"
)

func ParseCSV(filePath string) ([]transaction.Transaction, error)
{
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []transaction.Transaction
	for i, record := range records {
		if i == 0 { continue }
		date, _ := time.Parse("2006-01-02", record[0])
		amount, _ := strconv.ParseFloat(record[2], 64)
		transactions = append(transactions, transaction.Transaction{
			Date:        date,
			Description: record[1],
			Amount:      amount,
			Category:    "Neclasificat",
		})
	}
	return transactions, nil
}