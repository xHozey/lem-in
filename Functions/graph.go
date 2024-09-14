package ants

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
	Key      int
	Adjacent []*Vertix
}

func (g *Graph) AddVertix(name int) {
	if !g.Contains(name) {
		g.Vertices = append(g.Vertices, &Vertix{Key: name})
	}
}

func (g *Graph) AddIndirectedEdge(from, to int) {
	fromVertix := g.getVertix(from)
	toVertix := g.getVertix(to)
	if fromVertix != nil && toVertix != nil {
		if !g.Contains(to) {
			fromVertix.Adjacent = append(fromVertix.Adjacent, toVertix)
		}
		if !g.Contains(from) {
			toVertix.Adjacent = append(toVertix.Adjacent, fromVertix)
		}
	}
}

func (g *Graph) getVertix(key int) *Vertix {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Graph) Contains(key int) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}
