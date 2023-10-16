package main

import (
	"discord-bot/internal"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	internal.Init()

	internal.StartDiscordBot()
	defer internal.StopDiscordBot()

	fmt.Println("The bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-sc
}
