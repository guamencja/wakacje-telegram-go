package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Holiday struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	EndDate string `json:"end_date"`
	Message string `json:"message"`
}

func (h Holiday) GetCountdownDate() time.Time {
	layout := "2/1/2006"
	t, _ := time.Parse(layout, h.Date)

	if h.EndDate != "" {
		now := time.Now()
		if !now.Before(t) {

			t, _ = time.Parse(layout, h.EndDate)
			return t
		}
	}

	return t
}

func (h Holiday) GetName() string {
	return h.Name
}

func (h Holiday) GetMessage() string {
	return h.Message
}

func loadHolidays(h *[]Holiday) {
	data, err := os.ReadFile("holidays.json")
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(data, &h); err != nil {
		log.Println(err)
	}
}
