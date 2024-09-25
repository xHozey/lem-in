package ants

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
	Key      string
	Adjacent []*Vertix
}

type Road struct {
	Path []string
	Step int
}

type Step struct {
	Turn      int
	RoomIndex int
}
