package cli

import (
    "fmt"
    "os"
    "ATAD_PFCLIManager/internal/core/transaction"

    "github.com/olekukonko/tablewriter"
    "github.com/spf13/cobra"
)

func addListCommand(txService *transaction.Service) {
    var listCmd = &cobra.Command{
        Use:   "list",
        Short: "Afișează ultimele tranzacții",
        Run: func(cmd *cobra.Command, args []string) {
            
            transactions, err := txService.List()
            if err != nil {
                fmt.Printf("Eroare la listarea tranzacțiilor: %v\n", err)
                return
            }

            if len(transactions) == 0 {
                fmt.Println("Nicio tranzacție găsită. Adaugă una cu 'atad-pfcli add'.")
                return
            }

            table := tablewriter.NewWriter(os.Stdout)
            table.SetHeader([]string{"ID", "Data", "Descriere", "Suma", "Categorie"})
            table.SetBorder(false)
            table.SetRowLine(true)

            for _, tx := range transactions {
                dataStr := tx.Date.Format("2006-01-02")
                amountStr := fmt.Sprintf("%.2f RON", tx.Amount)
                
                if tx.Amount < 0 {
                    amountStr = "\033[31m" + amountStr + "\033[0m"
                } else {
                    amountStr = "\033[32m" + "+" + amountStr + "\033[0m"
                }

                row := []string{
                    fmt.Sprint(tx.ID),
                    dataStr,
                    tx.Description,
                    amountStr,
                    tx.Category,
                }
                table.Append(row)
            }
            table.Render()
        },
    }

    rootCmd.AddCommand(listCmd)
}