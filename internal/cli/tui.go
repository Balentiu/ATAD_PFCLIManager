package cli

import (
	"ATAD_PFCLIManager/internal/core/transaction"
	"ATAD_PFCLIManager/pkg/tui"
	"log"

	"github.com/spf13/cobra"
)

func addBrowseCommand(txService *transaction.Service) {
	var browseCmd = &cobra.Command{
		Use:   "browse",
		Short: "Deschide interfața interactivă de navigare",
		Run: func(cmd *cobra.Command, args []string) {
			txs, err := txService.List()
			if err != nil {
				log.Fatal(err)
			}
			if err := tui.StartTUI(txs); err != nil {
				log.Fatal(err)
			}
		},
	}
	rootCmd.AddCommand(browseCmd)
}
