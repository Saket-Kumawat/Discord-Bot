package bot

import (
	"fmt"

	"github.com/Saket-Kumawat/Discord-Bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	var err error
	goBot, err = discordgo.New("Bot " + config.Token) // one space after bot

	if err != nil {
		fmt.Println("1", err.Error())

		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("2", err.Error())

	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println("3", err.Error())
		return
	}
	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")

	}

	if m.Content == "hi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi how about you!")

	}

}
