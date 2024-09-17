package utils

import (
	"fmt"
	"strings"
)

func GetVertex(vertecies []*Vertex, key string) *Vertex {
	for _, vertex := range vertecies {
		if vertex.Key == key {
			return vertex
		}
	}

	return nil
}

func Printer(paths map[int]Road) {
	res := make([][]string, findResLength(paths))
	for ant, path := range paths {
		for i, room := range path.TheRoad[1:] {
			p := fmt.Sprintf("L%d-%s", ant, room)
			if len(res) > i {
				res[i+path.Step-1] = append(res[i+path.Step-1], p)
			}
		}
	}

	for _, p := range res {
		fmt.Println(strings.Join(p, " "))
	}
}

func findResLength(paths map[int]Road) int {
	maxPath := 0
	maxStep := 0

	for _, path := range paths {
		if maxPath < len(path.TheRoad) {
			maxPath = len(path.TheRoad)
		}

		if maxStep < path.Step {
			maxStep = path.Step
		}
	}

	return maxPath + maxStep - 2
}

func Sort(paths *[][]string) {
	for i := 0; i < len((*paths)); i++ {
		for j := 0; j < len((*paths)); j++ {
			if len((*paths)[i]) < len((*paths)[j]) {
				temp := (*paths)[i]
				(*paths)[i] = (*paths)[j]
				(*paths)[j] = temp
			}
		}
	}
}

func StepContains(steps []Step, step Step) bool {
	for _, s := range steps {
		if s.RoomIndex == step.RoomIndex && s.Turn == step.Turn {
			return true
		}
	}

	return false
}
