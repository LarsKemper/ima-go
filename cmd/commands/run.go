package commands

import (
	"github.com/LarsKemper/ima-go/internal/helper"
	imago "github.com/LarsKemper/ima-go/internal/ima-go"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run ima-go",
	Long:  `Run ima-go to generate ASCII art from an image file. Example: ima-go run --path path/to/file.png`,
	Run: func(cmd *cobra.Command, args []string) {
		options := imago.GetOptions(cmd)

		err := options.Validate()
		helper.HandleError(err)

		helper.HandleError(imago.Run(options))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// required flags
	runCmd.PersistentFlags().StringP("path", "p", "", "Path to the file to convert to ASCII art")
	helper.HandleError(runCmd.MarkPersistentFlagRequired("path"))

	// optional flags
	runCmd.PersistentFlags().BoolP("color", "c", false, "Use color in the output")
	runCmd.PersistentFlags().IntP("precision", "r", 10, "Precision of the ASCII art")
	runCmd.PersistentFlags().StringP("charset", "s", "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ", "Characters to use for the ASCII art, from light to dark")
	runCmd.PersistentFlags().BoolP("invert", "i", false, "Invert the grayscale")
	runCmd.PersistentFlags().Float64P("xScale", "x", 12.0, "X scale the output")
	runCmd.PersistentFlags().Float64P("yScale", "y", 5.0, "Y scale the output")
}
