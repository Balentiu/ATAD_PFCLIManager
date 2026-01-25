package cli

import (
	"ATAD_PFCLIManager/internal/core/transaction"
	"ATAD_PFCLIManager/pkg/parser"
	"fmt"

	"github.com/spf13/cobra"
)

func addImportCommand(txService *transaction.Service) {
	var filePath string

	var importCmd = &cobra.Command{
		Use:   "import",
		Short: "Importă tranzacții dintr-un fișier CSV",
		Run: func(cmd *cobra.Command, args []string) {
			if filePath == "" {
				fmt.Println("Te rugăm să specifici calea către fișier cu --file")
				return
			}

			txs, err := parser.ParseCSV(filePath)
			if err != nil {
				fmt.Printf("Eroare la citirea fișierului: %v\n", err)
				return
			}

			count, err := txService.ImportTransactions(txs)
			if err != nil {
				fmt.Printf("Eroare la import: %v\n", err)
				return
			}

			fmt.Printf("Succes! Am importat %d tranzacții.\n", count)
		},
	}

	importCmd.Flags().StringVarP(&filePath, "file", "f", "", "Calea către fișierul CSV")
	rootCmd.AddCommand(importCmd)
}
