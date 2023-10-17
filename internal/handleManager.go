package internal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messagesManager(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")
	fmt.Println(args)

	for i := 0; i < len(args); i++ {
		fmt.Println("'" + args[i] + "'")
	}

	if m.Content[0] != '!' {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	_, err := s.ChannelMessageSend(m.ChannelID, m.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
