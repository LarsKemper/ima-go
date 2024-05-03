package app

import (
	"errors"
	"fmt"
	"log"
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

func Run(args []string) {
	if len(args) <= 0 {
		log.Fatal("path has to be defined")
	}

	var path = args[0]

	if path == "" {
		return
	}

	var file, err = parseFile(path)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
