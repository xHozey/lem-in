package ants

import (
	"sort"
)

func Travel() {
	for i := 0; i < len(PossiblePaths); i++ {
		PossiblePaths[i] = PossiblePaths[i][1:]
		PossiblePaths[i] = PossiblePaths[i][:len(PossiblePaths[i])-1]
	}
	sort.SliceStable(PossiblePaths, func(i, j int) bool {
		return len(PossiblePaths[i]) < len(PossiblePaths[j])
	})
}
