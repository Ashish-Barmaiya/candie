package cli

import (
	"github.com/ashish-barmaiya/candie/internal/app"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play extracted wallpapers",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunPlay()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
