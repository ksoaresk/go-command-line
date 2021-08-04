package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	app := app.Gerar()
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
