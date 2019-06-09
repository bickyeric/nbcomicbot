package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/bickyeric/arumba/telegram/message"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type MessageHandler map[string]message.Handler

// NewMessageHandler ...
func NewMessageHandler(app arumba.Arumba, bot arumba.Bot, kendang connection.IKendang) MessageHandler {
	telegraphCreator := telegraph.NewCreatePage()
	readerService := comic.NewRead(app, kendang, telegraphCreator)
	handlers := map[string]message.Handler{}

	handlers[message.StartCommand] = message.NewStart(bot, readerService)

	handlers[message.HelpCommand] = message.NewHelp(bot)

	handlers[message.ReadCommand] = message.NewRead(bot, readerService)

	handlers[message.FeedbackCommand] = message.NewFeedback(bot)

	handlers[message.FollowCommand] = message.NewFollow(bot)

	handlers[message.GenericCommand] = message.NewGeneric(
		bot, bot,
		comic.Search{
			Repo: app.ComicRepo,
		},
	)

	return handlers
}

func (handler MessageHandler) Handle(m *tgbotapi.Message) {
	log.WithFields(
		log.Fields{
			"text":    m.Text,
			"chat_id": m.Chat.ID,
		},
	).Info("Handling message")

	h, ok := handler[m.Command()]
	if ok {
		h.Handle(m)
	} else {
		handler[message.GenericCommand].Handle(m)
	}
}
