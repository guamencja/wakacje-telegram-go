// self-explainatory
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/guamencja/gownobot/telegram"
	"github.com/joho/godotenv"
)

type config struct {
	Token    string
	Cooldown int
	ChatId   string
}

func getConfig() config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	cooldown, _ := strconv.Atoi(os.Getenv("COOLDOWN"))

	return config{
		Token:    os.Getenv("TOKEN"),
		Cooldown: cooldown,
		ChatId:   os.Getenv("CHAT_ID"),
	}
}

type temp struct {
	MessageId string `json:"message_id"`
}

const tempFile = "temp.json"

func getTempFile() (temp, error) {
	var t temp

	// ensure file exists
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		return t, nil // return empty
	}

	data, err := os.ReadFile(tempFile)
	if err != nil {
		return t, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return t, err
	}

	return t, nil
}

func saveTempFile(t temp) error {
	if err := os.MkdirAll(filepath.Dir(tempFile), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(tempFile, data, 0644)
}

func findNearestHoliday(holidays []Holiday) Holiday {
	var nearest Holiday
	var nearestDiff time.Duration = 1<<63 - 1 // max duration

	now := time.Now()
	for _, h := range holidays {
		eventTime := h.GetCountdownDate()
		diff := eventTime.Sub(now)

		if diff > 0 && diff < nearestDiff {
			nearest = h
			nearestDiff = diff
		}
	}
	return nearest
}

func main() {
	config := getConfig()

	bot := telegram.New(config.Token)

	user, err := bot.GetMe()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Logged in as @%s (%d)", user.Username, user.Id)

	// CHECK FOR MESSAGE ID

	// load temp file
	t, err := getTempFile()
	if err != nil {
		log.Fatalf("failed to read temp file: %v", err)
	}

	if t.MessageId == "" {
		id, err := bot.SendMessageText(config.ChatId, "loading...")
		if err != nil {
			log.Println(err)
		}

		if err := bot.PinMessage(config.ChatId, id); err != nil {
			log.Println(err)
		}

		t.MessageId = id
		if err := saveTempFile(t); err != nil {
			log.Println("failed to save temp file:", err)
		}
	}

	var holidays []Holiday
	loadHolidays(&holidays)

	// COUNTDOWN
	/*summer := getSummer()*/
	for range time.Tick(time.Second * time.Duration(config.Cooldown)) {
		nearest := findNearestHoliday(holidays)
		d := GetTimeRemaining( /*summer*/ nearest)

		/*if d.Total <= 0 { // bot odlicza na minusie, zresetuj wartości eventu (prawdopodobnie błąd związany z zmianą roku)
			summer = getSummer()
		}

		str := "wakacji! ☀️🍹"
		if summer.IsItAlready() {
			str = "jesieni! 🌆"
		}*/

		str := nearest.GetName()
		/*if nearest.IsItAlready() {
			str = fmt.Sprintf("końca %s", str)
		}*/

		text := fmt.Sprintf("%d dni, %d godz, %d min i %d sek do %s", d.Days, d.Hours, d.Minutes, d.Seconds, str)

		if err := bot.EditMessageText(config.ChatId, t.MessageId, text); err != nil {
			log.Println(err)
		}
	}
}
