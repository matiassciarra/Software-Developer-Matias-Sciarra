package structs

import (
	"time"
)

type DbItem struct {
	ID         uint64
	Action     string
	Brokerage  string
	Company    string
	RatingFrom string
	RatingTo   string
	TargetFrom float32
	TargetTo   float32
	Ticker     string
	Time       time.Time
}
