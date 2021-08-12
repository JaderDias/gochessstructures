package gochessstructures

type Event struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`
	RoundCount  int    `json:"roundCount"`
	ProcessStep int
}
