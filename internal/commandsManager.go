package internal

import "discord-bot/internal/commands"

func commandsManager(args []string) (messageToSend string) {
	mapCommandes := map[string]func([]string) string{
		"!help":    commands.CommandHelp,
		"!profile": commands.CommandProfile,
	}
	val, ok := mapCommandes[args[0]]
	if ok {
		return val(args)
	} else {
		return "Not a command"
	}
}
