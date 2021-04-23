package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"

	"github.com/IsmaeelAkram/discord_test/cmds"
	"github.com/IsmaeelAkram/discord_test/common"
)

var token, prefix string
var commands = make(map[string]common.Command)

func main() {
	godotenv.Load()
	token = os.Getenv("TOKEN")
	prefix = os.Getenv("PREFIX")
	fmt.Println(token, prefix)

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	registerCommands()

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is ready")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}

func registerCommands() {
	registerCommand("ping", common.Command{
		Aliases: []string{"p"},
		Handler: cmds.CommandPing,
	})
	registerCommand("pong", common.Command{
		Aliases: nil,
		Handler: cmds.CommandPong,
	})
	registerCommand("fact", common.Command{
		Aliases: nil,
		Handler: cmds.CommandFact,
	})
}

func registerCommand(name string, command common.Command) {
	commands[name] = command
	fmt.Printf("Command \"%s\" registered\n", name)
}

func getCommand(cmd string) (common.Command, error) {
	cmdName := strings.Replace(cmd, prefix, "", 1)
	command := commands[cmdName]

	// If struct is zero value
	if command.Handler == nil {
		return command, errors.New("Command " + cmdName + " not found")
	}

	return command, nil
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Printf("%s: %s\n", message.Author.Username, message.Content)
	if !strings.HasPrefix(message.Content, prefix) {
		return
	}
	msg := strings.Split(message.Content, " ")
	command, err := getCommand(msg[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	command.Handler(session, message)
}
