package chessgostructures

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ChessPlayer struct {
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

type ChessMove struct {
	Id       int    `json:"id"`
	Ply      int    `json:"ply"`
	Cbn      string `json:"cbn"`
	Clock    int    `json:"clock"`
	Time     int    `json:"time"`
	Revision int    `json:"revision"`
}

type ChessGame struct {
	Id          int         `json:"id"`
	Slug        string      `json:"slug"`
	RoundId     int         `json:"roundId"`
	Result      string      `json:"result"`
	Table       int         `json:"table"`
	Board       int         `json:"board"`
	CurrentFEN  string      `json:"currentFEN"`
	White       ChessPlayer `json:"white"`
	Black       ChessPlayer `json:"black"`
	TimeControl string      `json:"timeControl"`
	ProcessStep int
	RoundSlug   string
	EventSlug   string
	Moves       []ChessMove `json:"moves"`
}

type GameBody struct {
	Game ChessGame `json:"game"`
}

type ChessEvent struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`
	RoundCount  int    `json:"roundCount"`
	ProcessStep int
}

type ChessEvents struct {
	Events []ChessEvent `json:"events"`
}
