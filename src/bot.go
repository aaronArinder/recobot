package recobot

import (
	"fmt"

	"github.com/aaronArinder/recobot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}

	if message.Content == "ping" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "pong")
	}
}
