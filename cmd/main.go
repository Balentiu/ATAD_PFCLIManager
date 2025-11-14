package main

import (
    "log"

    "ATAD_PFCLIManager/internal/cli"
    "ATAD_PFCLIManager/internal/core/transaction"
    "ATAD_PFCLIManager/internal/database"
    "ATAD_PFCLIManager/pkg/config"
    "ATAD_PFCLIManager/pkg/logger"
)

func main() {
    cfg := config.Load()
    logger := logger.New()
    

    db, err := database.InitDB("./finance.db")
    if err != nil {
        log.Fatalf("Eroare fatală la inițializarea bazei de date: %v", err)
    }
    defer db.Close()

    txRepo := transaction.NewSQLiteRepository(db)
    txService := transaction.NewService(txRepo)
    
    appServices := cli.AppServices{
        TxService: txService,
        BudgetService: budgetService,
    }
    cli.SetupCommands(appServices)

    cli.Execute()
}