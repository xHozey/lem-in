package ants

import (
	"fmt"
	"sort"
)

// Function store depth first search needed arguments
func (g *Graph) PathFinder(start, end string) map[int][][]string {
	visited := make(map[string]bool)
	var currentPath []string
	var possiblePath [][]string
	g.Dfs(start, end, visited, currentPath, &possiblePath)
	sort.SliceStable(possiblePath, func(i, j int) bool {
		return len(possiblePath[i]) < len(possiblePath[j])
	})
	valid := allValidPaths(possiblePath)
	return valid
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

func allValidPaths(paths [][]string) map[int][][]string {
	combinations := make(map[int][][]string)
	i := 0

	for _, p := range paths {
		visited := false

		if len(combinations[i]) == 0 {
			combinations[i] = append(combinations[i], p)
			continue
		}
		for _, v := range combinations[i] {
			if !disjoint(p, v) {
				visited = true
				break
			}
		}
		if !visited {
			combinations[i] = append(combinations[i], p)
		} else {
			i++
			combinations[i] = append(combinations[i], p)
		}
	}
	for i, v := range combinations {
		alreadyAdded := make(map[string]bool)
		for _, p1 := range v {
			alreadyAdded[getKey(p1)] = true
		}

		for _, p2 := range paths {
			if alreadyAdded[getKey(p2)] {
				continue
			}
			addable := true
			for _, p1 := range combinations[i] {
				if !disjoint(p1, p2) {
					addable = false
					break
				}
			}
			if addable {
				combinations[i] = append(combinations[i], p2)
				alreadyAdded[getKey(p2)] = true
			}
		}
	}

	return combinations
}

func getKey(path []string) string {
	return fmt.Sprintf("%v", path)
}

func disjoint(path1 []string, path2 []string) bool {
	visited := make(map[string]bool)
	for _, p := range path1[1 : len(path1)-1] {
		visited[p] = true
	}
	for _, p := range path2[1 : len(path2)-1] {
		if visited[p] {
			return false
		}
	}
	return true
}
