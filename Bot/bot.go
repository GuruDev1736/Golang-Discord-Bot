package bot

import (
	"fmt"

	config "github.com/GuruDev1736/Golang_Discord_Bot/Config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = user.ID

	goBot.AddHandler(MessageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
	}

}
