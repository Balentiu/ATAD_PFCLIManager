package cli

import (
	"ATAD_PFCLIManager/internal/core/budget"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func addBudgetCommand(budgetService *budget.Service) {
	budgetCmd := &cobra.Command{Use: "budget", Short: "Management bugete"}

	setCmd := &cobra.Command{
		Use:  "set [categorie] [suma]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			amount, _ := strconv.ParseFloat(args[1], 64)
			budgetService.SetBudget(args[0], amount)
			fmt.Println("✅ Buget salvat!")
		},
	}

	checkCmd := &cobra.Command{
		Use: "check",
		Run: func(cmd *cobra.Command, args []string) {
			reports, _ := budgetService.GetBudgetReport()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Categorie", "Buget", "Cheltuit", "Status"})
			for _, r := range reports {
				table.Append([]string{r.Category, fmt.Sprintf("%.2f", r.Budget), fmt.Sprintf("%.2f", r.Spent), r.Status})
			}
			table.Render()
		},
	}

	budgetCmd.AddCommand(setCmd, checkCmd)
	rootCmd.AddCommand(budgetCmd)
}
