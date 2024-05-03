package app

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ExtPng  string = "png"
	ExtJpg  string = "jpg"
	ExtJpeg string = "jpeg"
	MimePng string = "image/png"
	MimeJpg string = "image/jpeg"
)

type File struct {
	name      string
	extension string
	mime      string
	content   []byte
}

func matchMime(extension string) (string, error) {
	switch extension {
	case ExtPng:
		return MimePng, nil
	case ExtJpg:
		return MimeJpg, nil
	case ExtJpeg:
		return MimeJpg, nil
	default:
		return "", errors.New("file extension not supported")
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

	var mime, errM = matchMime(extension)

	if errM != nil {
		return File{}, errM
	}

	var content, errC = os.ReadFile(path)

	if errC != nil {
		return File{}, errors.New("failed to read file content")
	}

	return File{filename, extension, mime, content}, nil
}

func Run(args []string) error {
	if len(args) <= 0 {
		return errors.New("not arguments given")
	}

	var path = args[0]

	if path == "" {
		return errors.New("path has to be defined")
	}

	var file, err = parseFile(path)

	if err != nil {
		return err
	}

	fmt.Println(file)

	return nil
}
