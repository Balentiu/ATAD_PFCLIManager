package budget

import "database/sql"

type Repository interface {
	UpsertBudget(b Budget) error
	GetAllBudgets() ([]Budget, error)
}

type sqliteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) Repository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) UpsertBudget(b Budget) error {
	query := `INSERT OR REPLACE INTO budgets (category, amount) VALUES (?, ?)`
	_, err := r.db.Exec(query, b.Category, b.Amount)
	return err
}

func (r *sqliteRepository) GetAllBudgets() ([]Budget, error) {
	rows, err := r.db.Query("SELECT category, amount FROM budgets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []Budget
	for rows.Next() {
		var b Budget
		if err := rows.Scan(&b.Category, &b.Amount); err != nil {
			return nil, err
		}
		budgets = append(budgets, b)
	}
	return budgets, nil
}
