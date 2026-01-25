package transaction

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	AddTransaction(tx Transaction) error
	GetTransactions() ([]Transaction, error)
	SearchTransactions(query string) ([]Transaction, error)
}

type sqliteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) Repository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) AddTransaction(tx Transaction) error {
	stmt, err := r.db.Prepare("INSERT INTO transactions(date, description, amount, category) VALUES(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("eroare la pregătirea interogării: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(tx.Date, tx.Description, tx.Amount, tx.Category)
	if err != nil {
		return fmt.Errorf("eroare la executarea interogării: %w", err)
	}
	return nil
}

func (r *sqliteRepository) GetTransactions() ([]Transaction, error) {
	rows, err := r.db.Query("SELECT id, date, description, amount, category FROM transactions ORDER BY date DESC")
	if err != nil {
		return nil, fmt.Errorf("eroare la interogarea tranzacțiilor: %w", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var tx Transaction
		if err := rows.Scan(&tx.ID, &tx.Date, &tx.Description, &tx.Amount, &tx.Category); err != nil {
			return nil, fmt.Errorf("eroare la scanarea rândului: %w", err)
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *sqliteRepository) SearchTransactions(query string) ([]Transaction, error) {
	sqlQuery := `SELECT id, date, description, amount, category FROM transactions 
	             WHERE description LIKE ? OR category LIKE ? 
	             ORDER BY date DESC`

	searchTerm := "%" + query + "%"
	rows, err := r.db.Query(sqlQuery, searchTerm, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var tx Transaction
		if err := rows.Scan(&tx.ID, &tx.Date, &tx.Description, &tx.Amount, &tx.Category); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
