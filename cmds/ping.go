package cmds

import "github.com/bwmarrin/discordgo"

func CommandPing(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" Pong!")
}
