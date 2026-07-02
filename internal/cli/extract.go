package cli

import (
	"github.com/ashish-barmaiya/candie/internal/app"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract <movie>",
	Short: "Extract wallpapers from a video",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.RunExtract(args[0])
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
