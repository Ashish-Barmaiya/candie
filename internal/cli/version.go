package cli

import (
	"github.com/ashish-barmaiya/candie/internal/app"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Candie version",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
