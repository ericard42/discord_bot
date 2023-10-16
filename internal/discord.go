package internal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var discordSession *discordgo.Session = nil

func StartDiscordBot() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	discordSession = sess
	if err != nil {
		log.Fatal(err)
	}
	discordSession.AddHandler(messagesManager)
	discordSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentMessageContent
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func StopDiscordBot() {
	fmt.Println("Closing the bot")
	discordSession.Close()
}
