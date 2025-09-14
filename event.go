package main

import "time"

type Event interface {
	GetCountdownDate() time.Time
	GetName() string
}
