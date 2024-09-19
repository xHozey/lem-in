package ants

import "fmt"

func (g *Graph) ValidPath() {
	var pathCounter []int
	counter := 0
	for i := 0; i < len(PossiblePaths); i++ {
		for j := 0; j < len(PossiblePaths[i]); j++ {
			if PossiblePaths[i][j] == g.getVertix(PossiblePaths[i][j]).Key {
				counter += g.getVertix(PossiblePaths[i][j]).counter
			}
		}
		pathCounter = append(pathCounter, counter)
		counter = 0
	}

	
}
