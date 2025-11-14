package cli

import (
    "fmt"
    "ATAD_PFCLIManager/internal/core/transaction"
    "time"

    "github.com/spf13/cobra"
)

func addAddCommand(txService *transaction.Service) {
    var (
        amount      float64
        description string
        category    string
    )

    var addCmd = &cobra.Command{
        Use:   "add",
        Short: "Adaugă o nouă tranzacție (venit sau cheltuială)",
        Example: `atad-pfcli add --amount -50.5 --desc "Cafea" --category "Mâncare"
atad-pfcli add --amount 2500 --desc "Salariu" --category "Venit"`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Se adaugă tranzacția...")

            err := txService.Add(time.Now(), description, amount, category)
            if err != nil {
                fmt.Printf("Eroare la adăugarea tranzacției: %v\n", err)
                return
            }

            fmt.Println("Tranzacție adăugată cu succes!")
        },
    }

    addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Suma tranzacției (negativ pt. cheltuieli)")
    addCmd.Flags().StringVarP(&description, "desc", "d", "", "Descrierea tranzacției")
    addCmd.Flags().StringVarP(&category, "category", "c", "General", "Categoria tranzacției")

    addCmd.MarkFlagRequired("amount")
    addCmd.MarkFlagRequired("desc")

    rootCmd.AddCommand(addCmd)
}