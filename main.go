package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "NTk2MDgxNzk5Njc2MDM1MDcy.XR0cUA.isqhYSqu0s2I7dNQ6VlhZc2w4nY")
	if err != nil {
		return
	}

	discord.AddHandler(messageCreate)
	err = discord.Open()

	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(s.UserChannelCreate("297046031017836544"))
	if m.Content == "ping" {
		s.ChannelMessageSend("596094778681065493", "pong")
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
