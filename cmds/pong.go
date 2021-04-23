package cmds

import "github.com/bwmarrin/discordgo"

func CommandPong(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" Ping!")
}
