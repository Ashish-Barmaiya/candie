package cli

import (
	"github.com/ashish-barmaiya/candie/internal/app"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Candie installation and environment",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunDoctor()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
