package ants

var (
	VisitedVertix = make(map[string]bool)
	currentPath   []string
	ShortestPath  []string
	PossiblePaths [][]string
)

func (g *Graph) Dfs(start, end string) {
	currentPath = append(currentPath, start)
	if start == end {
		ShortestPath = make([]string, len(currentPath))
		copy(ShortestPath, currentPath)
		PossiblePaths = append(PossiblePaths, ShortestPath)

	}
	VisitedVertix[start] = true

	v := g.getVertix(start)
	for _, neigbors := range v.Adjacent {
		if !VisitedVertix[neigbors.Key] {
			g.Dfs(neigbors.Key, end)
		}
	}

	currentPath = currentPath[:len(currentPath)-1]
	VisitedVertix[start] = false
}
