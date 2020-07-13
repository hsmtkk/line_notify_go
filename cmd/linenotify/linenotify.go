package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	notifyImageCommand := &cobra.Command{
		Use:  "notifyimage path",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			fileName := filepath.Base(path)
			reader, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			if err := notifier.NotifyImage(fileName, reader); err != nil {
				log.Fatal(err)
			}
		},
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
	command.AddCommand(notifyImageCommand)
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
