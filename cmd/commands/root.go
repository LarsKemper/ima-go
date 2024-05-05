package commands

import (
	"github.com/LarsKemper/ima-go/internal/helper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ima-go",
	Short: "A simple ascii image generator written in Go",
	Long: `ima-go is a ASCII image generator implemented in Go.
	Create ASCII art effortlessly with this minimalist tool, perfect
	for quick creations and text-based visualizations.`,
}

func Execute() {
	helper.HandleError(rootCmd.Execute())
}
