package cli

import (
	"github.com/ashish-barmaiya/candie/internal/app"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract wallpapers from a video",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunExtract()
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
