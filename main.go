package main

import (
	"log"
	"tiki/app"
	"tiki/cmd"
)

func main() {
	err := app.LoadDataFromFile()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
