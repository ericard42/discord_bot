package internal

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messagesManager(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0][0] == '!' {
		messageToRet := commandsManager(args)
		_, err := s.ChannelMessageSend(m.ChannelID, messageToRet)
		if err != nil {
			return
		}
	}

	//for i := 0; i < len(args); i++ {
	//	fmt.Println("'" + args[i] + "'")
	//	if args[i][0] == '!' {
	//		fmt.Println("a Command" + args[i])
	//	}
	//}
	//
	//if m.Content == "ping" {
	//	_, err := s.ChannelMessageSend(m.ChannelID, "pong")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}
	//_, err := s.ChannelMessageSend(m.ChannelID, m.Content)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
