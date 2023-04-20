package main

import (
	"flag"
	tgClient "link-collector-bot/clients/telegram"
	event_consumer "link-collector-bot/consumer/event-consumer"
	"link-collector-bot/events/telegram"
	"link-collector-bot/storage/files"
	"log"
)

const(
	tgBotHost = "api.telegram.org"
	storagePath = "storage"
	batchSize = 100
)

func main() {
	eventProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("server started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("server is crashed", err)
	}
}

func mustToken() string {
	token := flag.String("bot-token", "", "token for accessing bot")

	flag.Parse()

	if *token  == "" {
		log.Fatal("token is not specified")
	}

	return *token
}