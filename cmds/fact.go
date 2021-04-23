package cmds

import (
	"fmt"
	"io"
	"net/http"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

func CommandFact(session *discordgo.Session, message *discordgo.MessageCreate) {
	msg, err := session.ChannelMessageSendEmbed(message.ChannelID, embed.NewGenericEmbed("Loading fact...", ""))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	factApiEndpoint := "https://uselessfacts.jsph.pl/random.txt?language=en"

	resp, err := http.Get(factApiEndpoint)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	factEmbed := embed.NewGenericEmbed("Random Fact", string(body))
	session.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, factEmbed)
}
