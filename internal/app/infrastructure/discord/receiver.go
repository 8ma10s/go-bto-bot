package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/ssym0614/go-bto-bot/internal/app/domain"
	"log"
	"strings"
)

type Receiver struct {
	token          string
	commandPrefix  string
	commandChannel chan domain.MessageInteractor
}

func (r *Receiver) Start(ctx context.Context) error {

	defer close(r.commandChannel)

	dg, err := discordgo.New("Bot " + r.token)
	if err != nil {
		return err
	}

	dg.AddHandler(r.handleMessage)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	if err := dg.Open(); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		// Cleanly close down the Discord session.
		log.Println("Shutting down Discord...")
		if err := dg.Close(); err != nil {
			log.Println("Failed to gracefully shut down Discord Bot")
		}
		log.Println("Graceful shutdown complete!")
		return ctx.Err()
	}
}

func (r *Receiver) MessageReceptionChannel() chan domain.MessageInteractor {
	return r.commandChannel
}

func NewReceiver(commandPrefix string, token string) *Receiver {
	return &Receiver{commandPrefix: commandPrefix, token: token, commandChannel: make(chan domain.MessageInteractor)}
}

func (r *Receiver) handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		return
	}

	if c == nil {
		return
	}

	if c.Name != "debug" {
		return
	}

	log.Printf("%s: %s@%s\n", m.Author.Username, m.Content, c.Name)

	if r.isCommand(m.Content) {
		r.commandChannel <- &messageInteractor{s: s, m: m}
	}
}

func (r *Receiver) isCommand(message string) bool {
	return strings.HasPrefix(message, r.commandPrefix)
}
