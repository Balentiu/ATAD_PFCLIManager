package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, fmt.Errorf("nu s-a putut deschide baza de date: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("nu s-a putut conecta la baza de date: %w", err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS transactions (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "date" DATETIME NOT NULL,
        "description" TEXT NOT NULL,
        "amount" REAL NOT NULL,
        "category" TEXT NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Printf("Atenție: Nu s-a putut crea tabela: %v", err)
		return nil, fmt.Errorf("nu s-a putut crea tabela: %w", err)
	}

	createBudgetTableSQL := `
	CREATE TABLE IF NOT EXISTS budgets (
    "category" TEXT NOT NULL PRIMARY KEY,
    "amount" REAL NOT NULL
	);`

	_, err = db.Exec(createBudgetTableSQL)
	if err != nil {
		return nil, fmt.Errorf("nu s-a putut crea tabela de bugete: %w", err)
	}

	return db, nil
}
