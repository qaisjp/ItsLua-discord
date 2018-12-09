package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const message = `It's Lua, not LUA. https://www.lua.org/about.html
` + "```" + `
"Lua" (pronounced LOO-ah) means "Moon" in Portuguese. As such, it is neither an acronym nor an abbreviation, but a noun. More specifically, "Lua" is a name, the name of the Earth's moon and the name of the language. Like most names, it should be written in lower case with an initial capital, that is, "Lua". Please do not write it as "LUA", which is both ugly and confusing, because then it becomes an acronym with different meanings for different people. So, please, write "Lua" right!
` + "```"

func main() {
	tok, err := ioutil.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + strings.TrimSpace(string(tok)))
	if err != nil {
		panic(err)
	}

	if err := discord.Open(); err != nil {
		panic(err)
	}

	discord.AddHandler(onMessage)

	// Create a new signal receiver
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Started...")

	// Watch for a signal
	<-sc
}

func onMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot {
		return
	}
	if !strings.Contains(e.Message.Content, "LUA") {
		return
	}

	fmt.Println("SENDING")
	_, err := s.ChannelMessageSend(e.ChannelID, message)
	fmt.Println(err)
}
