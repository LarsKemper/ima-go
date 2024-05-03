package app

import (
	"errors"
	"fmt"
	"image/png"
	"io"
	"os"
	"strings"
)

const (
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

func readImagePix(file File) error {
	var image, err = png.Decode(file.content)

	if err != nil {
		return err
	}

	fmt.Println(image.At(0, 0))

	return nil
}

func Run(args []string) error {
	if len(args) <= 0 {
		return errors.New("not arguments given")
	}

	var path = args[0]

	if path == "" {
		return errors.New("path has to be defined")
	}

	var file, errF = parseFile(path)

	if errF != nil {
		return errF
	}

	var errR = readImagePix(file)

	if errR != nil {
		return errR
	}

	return nil
}
