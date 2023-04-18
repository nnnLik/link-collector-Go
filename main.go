package main

import (
	"flag"
	"link-collector-bot/clients/telegram"
	"log"
)

const(
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())
} 

func mustToken() string {
	token := flag.String("bot-token", "", "token for accessing bot")

	flag.Parse()

	if *token  == "" {
		log.Fatal("token is not specified")
	}
}