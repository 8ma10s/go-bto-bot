package discord

import (
	"github.com/bwmarrin/discordgo"
)

type messageInteractor struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func (m *messageInteractor) Message() string {
	return m.m.Content[1:]
}

func (m *messageInteractor) Send(message string) error {
	if _, err := m.s.ChannelMessageSend(m.m.ChannelID, message); err != nil {
		return err
	}

	return nil
}
