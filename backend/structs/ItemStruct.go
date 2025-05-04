package structs

type Item struct {
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	Company    string `json:"company"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Ticker     string `json:"ticker"`
	Time       string `json:"time"`
}
