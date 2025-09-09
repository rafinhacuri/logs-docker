package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

func initDiscord() error {
	tokenDiscord := os.Getenv("DISCORD_TOKEN")

	var err error
	discordSession, err = discordgo.New("Bot " + tokenDiscord)
	if err != nil {
		return err
	}

	if err := discordSession.Open(); err != nil {
		return err
	}

	return nil
}

func sendDiscordMessage(sistema, mensagem, tipo string) error {
	channelID := os.Getenv("DISCORD_CHANNEL")

	var color int
	switch tipo {
	case "error":
		color = 0xF04747
	default:
		color = 0x7289DA
	}

	embed := &discordgo.MessageEmbed{
		Title:       sistema,
		Description: mensagem,
		Color:       color,
	}

	_, err := discordSession.ChannelMessageSendEmbed(channelID, embed)
	return err
}
