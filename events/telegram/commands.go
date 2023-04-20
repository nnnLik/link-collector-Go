package telegram

import (
	"log"
	"strings"
)

const (
	RndCmd = "/rnd"
	HelpCmd = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:
	}
}

func isAddCmd(text string) bool {
	return isURL(text)
}
// TODO
func isURL(text string) bool {
	return true
}