package model

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PathRequest struct {
	Start Coordinate `json:"start"`
	End   Coordinate `json:"end"`
}

type Path struct {
	Path []Coordinate `json:"path"`
}

type Robot struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

type Space struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}
