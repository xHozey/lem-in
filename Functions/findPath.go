package ants

import "sort"

// Function store depth first search needed arguments
func (g *Graph) PathFinder(start, end string) [][]string {
	visited := make(map[string]bool)
	var currentPath []string
	var possiblePath [][]string
	g.Dfs(start, end, visited, currentPath, &possiblePath)
	sort.SliceStable(possiblePath, func(i, j int) bool {
		return len(possiblePath[i]) < len(possiblePath[j])
	})
	return possiblePath
}

// function travl in our graph with dfs method then backtrack to get path from start to end
func (g *Graph) Dfs(start, end string, visited map[string]bool, currentPath []string, possiblePath *[][]string) {
	currentPath = append(currentPath, start)
	if start == end {
		currentPathCopy := make([]string, len(currentPath))
		copy(currentPathCopy, currentPath)
		*possiblePath = append(*possiblePath, currentPathCopy)
	}

	visited[start] = true

	v := g.getVertix(start)
	for _, neigbors := range v.Adjacent {
		if !visited[neigbors.Key] {
			g.Dfs(neigbors.Key, end, visited, currentPath, possiblePath)
		}
	}

	currentPath = currentPath[:len(currentPath)-1]
	visited[start] = false
}
