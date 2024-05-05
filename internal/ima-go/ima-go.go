package imaGo

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
)

const (
	Charset string = "@#S%?*+;:,."
	PngExt  string = "png"
	PngMime string = "image/png"
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
	default:
		return "", errors.New("file extension not supported. Use .png")
	}
}

func parseFile(path string) (File, error) {
	var spPath = strings.Split(path, "/")

	if len(spPath) <= 1 {
		return File{}, errors.New("invalid path")
	}

	var filename = spPath[len(spPath)-1]

	if filename == "" {
		return File{}, errors.New("invalid path")
	}

	var spFilename = strings.Split(filename, ".")

	if len(spFilename) <= 1 {
		return File{}, errors.New("invalid file name")
	}

	var extension = spFilename[len(spFilename)-1]

	if extension == "" {
		return File{}, errors.New("invalid extension")
	}

	var mime, errM = matchFileType(extension)

	if errM != nil {
		return File{}, errM
	}

	file, err := os.Open(path)

	if err != nil {
		return File{}, errors.New("failed to open file")
	}

	return File{name: filename, extension: extension, mime: mime, content: file}, nil
}

func getRelativeRgbBrightness(color color.Color) int {
	var r, g, b, a = color.RGBA()

	var r8 = uint8(r >> 8)
	var g8 = uint8(g >> 8)
	var b8 = uint8(b >> 8)
	var a8 = uint8(a >> 8)

	var averageBrightness = (uint32(r8) + uint32(g8) + uint32(b8) + uint32(a8)) / 4

	return int((averageBrightness * 100) / 0xFF)
}

func getCharByBrightness(brightness int) rune {
	var charIndex = (brightness * (len(Charset) - 1)) / 100

	return rune(Charset[charIndex])
}

func getPixelCharByCoords(x int, y int, image image.Image) string {
	var colorValues = image.At(x, y)
	var brightness = getRelativeRgbBrightness(colorValues)

	return string(getCharByBrightness(brightness))
}

func Run(path string) error {
	if path == "" {
		return errors.New("path has to be defined")
	}

	var file, errF = parseFile(path)

	if errF != nil {
		return errF
	}

	var imageData, err = png.Decode(file.content)

	if err != nil {
		return err
	}

	var scaleFactor = 10.0

	var xMax = imageData.Bounds().Max.X
	var yMax = imageData.Bounds().Max.Y

	for y := 0; y < yMax; y += 2 {
		for x := 0; x < int(float64(xMax)*scaleFactor); x += 10 {
			var originalX = int(float64(x) / scaleFactor)
			var char = getPixelCharByCoords(originalX, y, imageData)

			fmt.Print(char)
		}

		fmt.Println()
	}

	return nil
}
