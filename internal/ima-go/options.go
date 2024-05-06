package imago

import (
	"errors"
	"github.com/LarsKemper/ima-go/internal/helper"
	"github.com/spf13/cobra"
)

type Options struct {
	Path      string
	Color     bool
	Precision int
	Charset   string
	Invert    bool
	XScale    float64
	YScale    float64
}

func (options Options) Validate() error {
	if options.Path == "" {
		return errors.New("path is required")
	}

	if options.Precision < 1 {
		return errors.New("precision must be at least 1")
	}

	if options.Charset == "" {
		return errors.New("charset is required")
	}

	if options.XScale <= 0 {
		return errors.New("xScale must be greater than 0")
	}

	if options.YScale <= 0 {
		return errors.New("yScale must be greater than 0")
	}

	return nil
}

func GetOptions(cmd *cobra.Command) Options {
	path, err := cmd.Flags().GetString("path")
	helper.HandleError(err)

	color, err := cmd.Flags().GetBool("color")
	helper.HandleError(err)

	precision, err := cmd.Flags().GetInt("precision")
	helper.HandleError(err)

	charset, err := cmd.Flags().GetString("charset")
	helper.HandleError(err)

	invert, err := cmd.Flags().GetBool("invert")
	helper.HandleError(err)

	xScale, err := cmd.Flags().GetFloat64("xScale")
	helper.HandleError(err)

	yScale, err := cmd.Flags().GetFloat64("yScale")
	helper.HandleError(err)

	return Options{
		Path:      path,
		Color:     color,
		Precision: precision,
		Charset:   charset,
		Invert:    invert,
		XScale:    xScale,
		YScale:    yScale,
	}
}
