package cli

import (
	"ATAD_PFCLIManager/internal/core/transaction"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func addSearchCommand(txService *transaction.Service) {
	var searchCmd = &cobra.Command{
		Use:   "search [termen]",
		Short: "Caută tranzacții după descriere sau categorie",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query := args[0]
			results, err := txService.Search(query)
			if err != nil {
				fmt.Printf("Eroare la căutare: %v\n", err)
				return
			}

			if len(results) == 0 {
				fmt.Printf("Nu s-a găsit nicio tranzacție pentru: %s\n", query)
				return
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Data", "Descriere", "Suma", "Categorie"})
			for _, tx := range results {
				table.Append([]string{
					fmt.Sprint(tx.ID),
					tx.Date.Format("2006-01-02"),
					tx.Description,
					fmt.Sprintf("%.2f", tx.Amount),
					tx.Category,
				})
			}
			table.Render()
		},
	}

	rootCmd.AddCommand(searchCmd)
}
