package main

import (
	"log"

	"ATAD_PFCLIManager/internal/cli"
	"ATAD_PFCLIManager/internal/core/budget"
	"ATAD_PFCLIManager/internal/core/transaction"
	"ATAD_PFCLIManager/internal/database"
)

func main() {

	db, err := database.InitDB("./finance.db")
	if err != nil {
		log.Fatalf("Eroare fatală la inițializarea bazei de date: %v", err)
	}
	defer db.Close()

	txRepo := transaction.NewSQLiteRepository(db)
	txService := transaction.NewService(txRepo)
	bRepo := budget.NewSQLiteRepository(db)
	bService := budget.NewService(bRepo, txRepo)

	appServices := cli.AppServices{
		TxService:     txService,
		BudgetService: bService,
	}
	cli.SetupCommands(appServices)

	cli.Execute()
}
