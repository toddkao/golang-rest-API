package main

import (
	"log"

	"github.com/toddkao/ecomm2/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
