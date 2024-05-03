package main

import (
	"github.com/LarsKemper/ima-go/internal/app"
	"log"
	"os"
)

func main() {
	var err = app.Run(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
}
