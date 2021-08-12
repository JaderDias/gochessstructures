package gochessstructures

type Player struct {
	Name               string  `json:"name"`
	Title              string  `json:"title"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	Elo                int     `json:"elo"`
	AvatarUrl          string  `json:"avatarUrl"`
	AvatarUrlLarge     string  `json:"avatarUrlLarge"`
	MasterPlayerUrl    string  `json:"masterPlayerUrl"`
	Country            Country `json:"country"`
	FideId             int     `json:"fideId"`
	Id                 int     `json:"id"`
	ShortTeam          string  `json:"shortTeam"`
	Team               string  `json:"team"`
	FideStandardRating int     `json:"fideStandardRating"`
	FideRapidRating    int     `json:"fideRapidRating"`
	FideBlitzRating    int     `json:"fideBlitzRating"`
	Username           string  `json:"username"`
}
