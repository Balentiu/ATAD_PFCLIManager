package cli

import (
	"ATAD_PFCLIManager/internal/core/report"
	"fmt"

	"github.com/spf13/cobra"
)

func addReportCommand(reportService *report.Service) {
	var reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Generează rapoarte vizuale ale cheltuielilor",
		Run: func(cmd *cobra.Command, args []string) {
			breakdown, err := reportService.GetCategoryBreakdown()
			if err != nil {
				fmt.Printf("Eroare la generarea raportului: %v\n", err)
				return
			}

			if len(breakdown) == 0 {
				fmt.Println("Nu există date pentru a genera un raport.")
				return
			}

			fmt.Println("\n--- DEFALCARE CHELTUIELI PE CATEGORII ---")

			maxVal := 0.0
			for _, val := range breakdown {
				if val > maxVal {
					maxVal = val
				}
			}

			for cat, val := range breakdown {
				fmt.Println(reportService.GenerateBarChart(cat, val, maxVal))
			}
			fmt.Println()
		},
	}

	rootCmd.AddCommand(reportCmd)
}
