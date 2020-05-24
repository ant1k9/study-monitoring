package main

import (
	"log"

	"github.com/ant1k9/study-monitoring/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
