package gochessstructures

type Game struct {
	Black       Player `json:"black"`
	Board       int    `json:"board"`
	CurrentFEN  string `json:"currentFEN"`
	EventSlug   string
	HalfMoves   int
	Id          int    `json:"id"`
	Moves       []Move `json:"moves"`
	ProcessStep int
	Result      string `json:"result"`
	RoundId     int    `json:"roundId"`
	RoundSlug   string
	Slug        string `json:"slug"`
	Table       int    `json:"table"`
	TimeControl string `json:"timeControl"`
	URL			string
	White       Player `json:"white"`
}
