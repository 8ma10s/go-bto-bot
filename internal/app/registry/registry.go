package registry

import (
	"github.com/ssym0614/go-bto-bot/internal/app/config"
	"github.com/ssym0614/go-bto-bot/internal/app/infrastructure/discord"
	"github.com/ssym0614/go-bto-bot/internal/app/service/command"
	"log"
)

type Registry struct {
	config *config.Config
	cw     *command.Worker
}

func (r *Registry) Worker() *command.Worker {
	return r.cw
}

func NewRegistry() *Registry {
	conf, err := config.Load()
	if err != nil {
		log.Panicf("Failed to load config: %s\n", err)
	}

	dr := discord.NewReceiver(conf.CommandPrefix, conf.Discord.Token)

	cw := command.NewWorker(dr)

	return &Registry{cw: cw}
}
