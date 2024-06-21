// everything countdown related
package main

import "time"

type timeRemaining struct {
	Total   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func getTimeRemaining(e Event) timeRemaining {
	now := time.Now()
	event := e.getCountdownDate()

	difference := event.Sub(now)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return timeRemaining{
		Total:   total,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}
