package internal

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func messagesManager(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			log.Println(err)
			return
		}
	}
	_, err := s.ChannelMessageSend(m.ChannelID, m.Content)
	if err != nil {
		log.Println(err)
		return
	}
}
