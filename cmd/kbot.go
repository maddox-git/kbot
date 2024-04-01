/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

// var (
// 	TeleToken = "7158382843:AAHk8_SwyTWQClav5A8rfaNIShFPeuVBQEg"
// )

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kbot called", appVersion)

		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
			return
		}

		token := os.Getenv("TELE_TOKEN")

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TOKEN env var %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			payload := m.Message().Payload
			log.Print(payload, m.Text())

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello  I'm bot %s", appVersion))
			}

			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
