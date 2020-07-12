package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/line_notify_go/pkg/linenotify"
	"github.com/spf13/cobra"
)

func main() {
	token, err := loadToken()
	if err != nil {
		log.Fatal(err)
	}
	notifier := linenotify.New(token)
	command := &cobra.Command{
		Use: "linenotify",
	}
	notifyMessageCommand := &cobra.Command{
		Use:  "notifymessage message",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			message := args[0]
			if err := notifier.NotifyMessage(message); err != nil {
				log.Fatal(err)
			}
		},
	}
	statusCommand := &cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {
			if err := notifier.Status(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("valid token")
		},
	}
	command.AddCommand(notifyMessageCommand)
	command.AddCommand(statusCommand)
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

const keyName = "LINE_NOTIFY_TOKEN"

func loadToken() (string, error) {
	token := os.Getenv(keyName)
	if token == "" {
		return "", fmt.Errorf("environment variable %s is not defined", keyName)
	}
	return token, nil
}
