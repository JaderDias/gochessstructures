package gochessstructures

type Game struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	RoundId     int    `json:"roundId"`
	Result      string `json:"result"`
	Table       int    `json:"table"`
	Board       int    `json:"board"`
	CurrentFEN  string `json:"currentFEN"`
	White       Player `json:"white"`
	Black       Player `json:"black"`
	TimeControl string `json:"timeControl"`
	ProcessStep int
	RoundSlug   string
	EventSlug   string
	Moves       []Move `json:"moves"`
}
