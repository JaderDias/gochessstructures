package gochessstructures

type Event struct {
	EndAt       string `json:"endAt"`
	Games       []Game `json:"games"`
	Id          int    `json:"id"`
	ImageUrl    string `json:"imageUrl"`
	Name        string `json:"name"`
	ProcessStep int
	RoundCount  int     `json:"roundCount"`
	Rounds      []Round `json:"rounds"`
	Slug        string  `json:"slug"`
	StartAt     string  `json:"startAt"`
}
