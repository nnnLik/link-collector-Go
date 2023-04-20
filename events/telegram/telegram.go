package telegram

import (
	"link-collector-bot/clients/telegram"
	"link-collector-bot/events"
	"link-collector-bot/lib/e"
	"link-collector-bot/storage"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

type Meta struct {
	ChatID int
	Username string
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
	tg: client,
	storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)

	if err != nil {
		return nil, e.Wrap("cannot get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}
	
	result := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		result = append(result, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return result, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", e.ErrUnknowEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)

	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	result, ok := event.Meta.(Meta)

	if !ok {
		return Meta{}, e.Wrap("can't get meta", e.ErrUnknowMetaType)
	}
	
	return result, nil
}

func event(u telegram.Update) events.Event {
	uTyp := fetchType(u)
	uTxt := fetchText(u)

	result := events.Event{
		Type: uTyp,
		Text: uTxt,
	}
	
	if uTyp == events.Message {
		result.Meta = Meta{
			ChatID: u.Message.Chat.ID,
			Username: u.Message.From.Username,
		}
	}

	return result
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}

	return u.Message.Text
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	}

	return events.Message
}