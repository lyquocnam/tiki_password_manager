package main

import (
	"github.com/lyquocnam/tiki_password_manager/app"
	"github.com/lyquocnam/tiki_password_manager/cmd"
	"log"
)

func main() {
	err := app.LoadDataFromFile()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
