package structs

import (
	"time"
)

type DbItem struct {
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
