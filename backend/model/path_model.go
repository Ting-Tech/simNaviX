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
