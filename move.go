package gochessstructures

type Move struct {
	Id       int    `json:"id"`
	Ply      int    `json:"ply"`
	Cbn      string `json:"cbn"`
	Clock    int    `json:"clock"`
	Time     int    `json:"time"`
	Revision int    `json:"revision"`
}
