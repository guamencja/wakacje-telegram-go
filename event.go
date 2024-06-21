package main

import "time"

type Event interface {
	getCountdownDate() time.Time
}
