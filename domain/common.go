package domain

type Vertices struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Label struct {
	Name       string  `json:"name"`
	Confidence float32 `json:"confidence"`
}
