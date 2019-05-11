package updater

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/updater/source"
	log "github.com/sirupsen/logrus"
)

// IRunner ...
type IRunner interface {
	Run(source source.ISource)
}

type runner struct {
	bot     arumba.IBot
	kendang connection.IKendang
	saver   episode.UpdateSaver
}

// NewRunner ...
func NewRunner(bot arumba.IBot, kendang connection.IKendang, saver episode.UpdateSaver) IRunner {
	return runner{
		bot:     bot,
		kendang: kendang,
		saver:   saver,
	}
}

// Run ...
func (r runner) Run(source source.ISource) {
	contextLogger := log.WithFields(log.Fields{
		"source": source.Name(),
	})
	contextLogger.Info("Processing updates...")

	updates, err := r.kendang.FetchUpdate("/" + source.Name() + "-update")
	if err != nil {
		r.bot.NotifyError(err)
		contextLogger.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Error fetching updates")
		return
	}

	for _, u := range updates {
		err := r.saver.Perform(u, source.GetID())
		if err != nil {
			if err.Error() == "episode exists" {
				continue
			} else {
				r.bot.NotifyError(err)
				contextLogger.WithFields(log.Fields{
					"update": u,
					"error":  err.Error(),
				}).Warn("Error processing update")
				continue
			}
		}

		r.bot.NotifyNewEpisode(u)
	}
	contextLogger.Info("Updates processed...")
}
