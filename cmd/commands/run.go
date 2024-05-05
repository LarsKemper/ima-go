package commands

import (
	"github.com/LarsKemper/ima-go/internal/helper"
	imaGo "github.com/LarsKemper/ima-go/internal/ima-go"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run ima-go",
	Long:  `Run ima-go to generate ASCII art from an image file. Example: ima-go run --path path/to/file.png`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		helper.HandleError(err)

		helper.HandleError(imaGo.Run(path))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringP("path", "p", "", "Path to the file to convert to ASCII art")
	helper.HandleError(runCmd.MarkPersistentFlagRequired("path"))
}
