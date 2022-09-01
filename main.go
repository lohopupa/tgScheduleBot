package main

import (
	"fmt"
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func main() {

	tgtoken := "5163964075:AAFyUJXUy8PS-lRO0em4ZejUZHAgfNcF2Vw"
	bot, err := tgapi.NewBotAPI(tgtoken)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	bot.Debug = true
	uc := tgapi.NewUpdate(0)
	uc.Timeout = 30

	for update := range bot.GetUpdatesChan(uc) {
		if update.Message == nil {
			continue
		}

		text := update.Message.Text
		if !strings.HasPrefix(text, "/") {
			//TODO: send default message
			log.Println("Receive not a command")
			continue
		}

		cmd, args := getCommand(text)
		switch cmd {
		case "/selectGroup":
			groupPrefix := ""
			if len(args) > 0 {
				groupPrefix = args[0]
			}
			fmt.Println(cmd, groupPrefix)
		case "/hello", "/start":
			msg := tgapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, %s", update.Message.Chat.UserName))
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func getCommand(text string) (string, []string) {
	t := strings.Split(text, " ")
	return t[0], t[1:]
}
