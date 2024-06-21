// implementation of Event
package main

import (
	"time"
)

type Summer struct {
	start time.Time
	end   time.Time
}

func getSummerStart(year int) time.Time {
	tz, _ := time.LoadLocation("Europe/Warsaw")
	start := time.Date(year, time.June, 1, 0, 0, 0, 0, tz)
	offset := 27 - 1 - int(start.Weekday())
	return start.AddDate(0, 0, offset)
}

func getSummerEnd(year int) time.Time {
	tz, _ := time.LoadLocation("Europe/Warsaw")
	end := time.Date(year, time.September, 1, 0, 0, 0, 0, tz)
	if end.Weekday() != time.Monday {
		offset := (7 + int(time.Monday) - int(end.Weekday())) % 7
		return end.AddDate(0, 0, offset)
	}
	return end
}

func getSummer() *Summer {
	now := time.Now()
	year := now.Year()

	start := getSummerStart(year)
	end := getSummerEnd(year)

	// jeśli wakacje się skończyły, aby uniknąć liczenia na minusie, odliczaj do następnego roku
	if now.After(end) {
		start = getSummerStart(now.Year() + 1)
	}

	return &Summer{
		start: start,
		end:   end,
	}
}

func (s Summer) IsItAlready() bool {
	now := time.Now()
	return !now.Before(s.start)
}

func (s Summer) getCountdownDate() time.Time {
	if s.IsItAlready() {
		return s.end
	}

	return s.start
}
