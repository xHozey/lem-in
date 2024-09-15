package ants

import "fmt"

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
	Key      string
	Adjacent []*Vertix
}

func (g *Graph) AddVertix(name string) {
	if !g.containsVertix(name) {
		g.Vertices = append(g.Vertices, &Vertix{Key: name})
	}
}

func (g *Graph) AddIndirectedEdge(from, to string) {
	fromVertix := g.getVertix(from)
	toVertix := g.getVertix(to)
	if fromVertix != nil && toVertix != nil {
		if !g.containsAdjacent(to, fromVertix) && !g.containsAdjacent(from, toVertix) {
			fromVertix.Adjacent = append(fromVertix.Adjacent, toVertix)
			toVertix.Adjacent = append(toVertix.Adjacent, fromVertix)

		}
	}
}

func (g *Graph) getVertix(key string) *Vertix {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Graph) containsAdjacent(key string, s *Vertix) bool {
	for _, v := range s.Adjacent {
		if v.Key == key {
			return true
		}
	}
	return false
}

func (g *Graph) containsVertix(key string) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}

func (g *Graph) PrintGraph() {
	for _, val := range g.Vertices {
		fmt.Printf("Room %v :", val.Key)
		for _, v := range val.Adjacent {
			fmt.Printf(" %v ", v.Key)
		}
		fmt.Println("")
	}
}
