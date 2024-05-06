package imago

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/aybabtme/rgbterm"
)

const (
	PngExt   string = "png"
	PngMime  string = "image/png"
	JpgExt   string = "jpg"
	JpegExt  string = "jpeg"
	JpegMime string = "image/jpeg"
)

type File struct {
	name      string
	extension string
	mime      string
	content   io.Reader
}

func matchFileType(extension string) (string, error) {
	switch extension {
	case PngExt:
		return PngMime, nil
	case JpgExt:
		return JpegMime, nil
	case JpegExt:
		return JpegMime, nil
	default:
		return "", errors.New("file extension not supported. Use .png, .jpg or .jpeg")
	}
}

func decodeByType(file File) (image.Image, error) {
	switch file.mime {
	case PngMime:
		return png.Decode(file.content)
	case JpegMime:
		return jpeg.Decode(file.content)
	default:
		return nil, errors.New("failed to decode file type. Use .png or .jpg")
	}
}

func parseFile(path string) (File, error) {
	spPath := strings.Split(path, "/")

	if len(spPath) <= 1 {
		return File{}, errors.New("invalid path")
	}

	filename := spPath[len(spPath)-1]

	if filename == "" {
		return File{}, errors.New("invalid path")
	}

	spFilename := strings.Split(filename, ".")

	if len(spFilename) <= 1 {
		return File{}, errors.New("invalid file name")
	}

	extension := spFilename[len(spFilename)-1]

	if extension == "" {
		return File{}, errors.New("invalid extension")
	}

	mime, errM := matchFileType(extension)

	if errM != nil {
		return File{}, errM
	}

	file, err := os.Open(path)
	if err != nil {
		return File{}, errors.New("failed to open file")
	}

	return File{name: filename, extension: extension, mime: mime, content: file}, nil
}

func getRelativeRgbBrightness(colorValues color.Color, options Options) int {
	r, g, b, a := colorValues.RGBA()

	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	a8 := uint8(a >> 8)

	averageBrightness := (uint32(r8) + uint32(g8) + uint32(b8) + uint32(a8)) / 4
	relativeBrightness := int((averageBrightness * 100) / 0xFF)

	if options.Invert {
		return 100 - relativeBrightness
	}

	return relativeBrightness
}

func getCharByBrightness(brightness int, options Options) rune {
	charIndex := (brightness * (len(options.Charset) - 1)) / 100

	return rune(options.Charset[charIndex])
}

func getPixelCharByCoords(x, y int, imageData image.Image, options Options) string {
	colorValues := imageData.At(x, y)
	brightness := getRelativeRgbBrightness(colorValues, options)

	char := string(getCharByBrightness(brightness, options))

	if options.Color {
		r, g, b, _ := colorValues.RGBA()
		char = rgbterm.FgString(char, uint8(r), uint8(g), uint8(b))
	}

	return char
}

func Run(options Options) error {
	if options.Path == "" {
		return errors.New("path has to be defined")
	}

	file, errF := parseFile(options.Path)

	if errF != nil {
		return errF
	}

	imageData, err := decodeByType(file)
	if err != nil {
		return err
	}

	xMax := imageData.Bounds().Max.X
	yMax := imageData.Bounds().Max.Y

	for y := 0; y < int(float64(yMax)*options.YScale); y += options.Precision {
		for x := 0; x < int(float64(xMax)*options.XScale); x += options.Precision {
			originalX := int(float64(x) / options.XScale)
			originalY := int(float64(y) / options.YScale)

			char := getPixelCharByCoords(originalX, originalY, imageData, options)

			fmt.Print(char)
		}

		fmt.Println()
	}

	return nil
}
