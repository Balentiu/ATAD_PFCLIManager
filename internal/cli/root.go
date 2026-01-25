package cli

import (
	"ATAD_PFCLIManager/internal/core/budget"
	"ATAD_PFCLIManager/internal/core/report"
	"ATAD_PFCLIManager/internal/core/transaction"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type AppServices struct {
	TxService     *transaction.Service
	BudgetService *budget.Service
	ReportService *report.Service
}

var rootCmd = &cobra.Command{
	Use:   "atad-pfcli",
	Short: "Manager de Finanțe Personale (ATAD_PFCLIManager)",
	Long:  `Un instrument CLI pentru a urmări veniturile, cheltuielile și bugetele.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetupCommands(services AppServices) {
	addAddCommand(services.TxService)
	addListCommand(services.TxService)
	addImportCommand(services.TxService)
	addBudgetCommand(services.BudgetService)
	addReportCommand(services.ReportService)
}
