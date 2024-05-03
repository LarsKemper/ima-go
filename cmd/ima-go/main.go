package main

import (
	"github.com/LarsKemper/ima-go/internal/app"
	"os"
)

func main() {
	app.Run(os.Args[1:])
}
