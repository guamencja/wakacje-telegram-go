// self-explainatory
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/guamencja/gownobot/telegram"
	"github.com/joho/godotenv"
)

type config struct {
	Token     string
	Cooldown  int
	ChatId    string
	MessageId string
}

func getConfig() config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	cooldown, _ := strconv.Atoi(os.Getenv("COOLDOWN"))

	return config{
		Token:     os.Getenv("TOKEN"),
		Cooldown:  cooldown,
		ChatId:    os.Getenv("CHAT_ID"),
		MessageId: os.Getenv("MESSAGE_ID"),
	}
}

func main() {
	config := getConfig()

	bot := telegram.New(config.Token)

	user, err := bot.GetMe()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Logged in as @%s (%d)", user.Username, user.Id)

	summer := getSummer()
	for range time.Tick(time.Second * time.Duration(config.Cooldown)) {
		d := getTimeRemaining(summer)

		if d.Total <= 0 { // bot odlicza na minusie, zresetuj wartości eventu (prawdopodobnie błąd związany z zmianą roku)
			summer = getSummer()
		}

		str := "wakacji! ☀️🍹"
		if summer.IsItAlready() {
			str = "jesieni! 🌆"
		}

		text := fmt.Sprintf("%d dni, %d godz, %d min i %d sek do %s", d.Days, d.Hours, d.Minutes, d.Seconds, str)
		if err := bot.EditMessageText(config.ChatId, config.MessageId, text); err != nil {
			log.Println(err)
		}
	}
}
