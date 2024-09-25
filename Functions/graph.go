package ants

// Add nodes to our struct
func (g *Graph) AddVertix(name string) {
	if !g.ContainsVertix(name) {
		g.Vertices = append(g.Vertices, &Vertix{Key: name})
	}
}

// Add an indirected edge between from and to
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

// Get vertix from provided key
func (g *Graph) getVertix(key string) *Vertix {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

// Check vertix adjacent list if it exist
func (g *Graph) containsAdjacent(key string, s *Vertix) bool {
	for _, v := range s.Adjacent {
		if v.Key == key {
			return true
		}
	}
	return false
}

// Check if the vertix exists
func (g *Graph) ContainsVertix(key string) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}
