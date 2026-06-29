package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "candie",
	Short: "Turn your favorite films into living wallpapers",
	Long: `Candie is a Linux-first CLI that extracts beautiful frames
from locally stored videos and plays them back as rotating desktop wallpapers.`,
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
