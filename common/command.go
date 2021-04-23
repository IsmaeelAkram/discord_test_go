package common

import "github.com/bwmarrin/discordgo"

type Command struct {
	Aliases []string
	Handler func(session *discordgo.Session, message *discordgo.MessageCreate)
}
