package telegram

import (
	"link-collector-bot/clients/telegram"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage int
}

func New(client *telegram.Client) {
	
}