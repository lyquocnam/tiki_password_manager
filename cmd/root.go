package cmd

import (
	"fmt"
	"github.com/lyquocnam/tiki_password_manager/app"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "tiki [command] [options]",
	Short: fmt.Sprintf(`"default username: %v password: %v"`, app.DemoUser, app.DemoPassword),
	Long:  fmt.Sprintf(`"default username: %v password: %v"`, app.DemoUser, app.DemoPassword),
}

func Execute() {
	rootCmd.AddCommand(loginCommand, createUserCommand, vaLidatePasswordCommand, changePasswordCommand)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
