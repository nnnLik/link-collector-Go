package main

import (
	"flag"
	"log"
)

func main() {
	t, err = mustToken()
}

func mustToken() string {
	token := flag.String("bot-token", "", "token for accessing bot")

	flag.Parse()

	if *token  == "" {
		log.Fatal("token is not specified")
	}
}