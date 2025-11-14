package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `
    CREATE TABLE IF NOT EXISTS transactions (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "date" DATETIME,
        "description" TEXT,
        "amount" REAL,
        "category" TEXT
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatalf("Nu s-a putut crea tabela: %v", err)
    }

    return db
}