package utils

import (
	"fmt"
	"strings"
)

type Scoretype []struct {
	path  int
	score int
}

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

func Sort(paths *[][]string, scoring *Scoretype) {
	for i := 0; i < len(*scoring); i++ {
		for j := 0; j < len(*scoring); j++ {
			if (*scoring)[i].score <= (*scoring)[j].score {
				temp := (*scoring)[i]
				(*scoring)[i] = (*scoring)[j]
				(*scoring)[j] = temp
			}
		}
	}

	res := make([][]string, len(*paths))
	for i, s := range *scoring {
		res[i] = (*paths)[s.path]
		(*scoring)[i].path = i
	}

	*paths = res
}

func Duplicated(paths *[][]string) Scoretype {
	res := make(Scoretype, len(*paths))
	for i, path := range *paths {
		res[i].score -= len(path)
		res[i].path = i
		for _, room := range path[1 : len(path)-1] {
			res[i].score += RoomsCounter[room]
		}
	}

	return res
}

func GetRoomCount(room string, paths *[][]string) int {
	res := 0
	for _, path := range *paths {
		for _, r := range path {
			if room == r {
				res++
			}
		}
	}

	return res
}
