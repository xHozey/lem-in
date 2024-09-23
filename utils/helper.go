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

func Printer(paths map[int]Road) string {
	res := map[int][]string{}
	str := ""
	for ant, path := range paths {
		for i, room := range path.TheRoad[1:] {
			p := fmt.Sprintf("L%d-%s", ant, room)
			res[i+path.Step-1] = append(res[i+path.Step-1], p)
		}
	}

	for i := 0; i < len(res); i++ {
		str += strings.Join(res[i], " ") + "\n"
	}

	return str
}

func RateSort(paths *[][]string, scoring *Scoretype) {
	for i := 0; i < len(*scoring); i++ {
		for j := 0; j < len(*scoring); j++ {
			if (*scoring)[i].score < (*scoring)[j].score {
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

func Sort(paths *[][]string) {
	for i := 0; i < len(*paths); i++ {
		for j := 0; j < len(*paths); j++ {
			if len((*paths)[i]) < len((*paths)[j]) {
				temp := (*paths)[i]
				(*paths)[i] = (*paths)[j]
				(*paths)[j] = temp
			}
		}
	}
}

func Duplicated(paths *[][]string) Scoretype {
	res := make(Scoretype, len(*paths))
	for i, path := range *paths {
		res[i].path = i
		res[i].score += len(path)
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

func CheckIfExist(roads [][]string, path []string) bool {
	direct := len(path) == 2
	for _, road := range roads {
		if len(road) == 0 || len(path) == 0 {
			return false
		}

		if direct && len(road) == 2 {
			return true
		}

		if DeepEqual(road[1:len(road)-1], path[1:len(path)-1]) {
			return true
		}
	}

	return false
}

func SepRoads(paths *[][]string) map[int][][]string {
	res := map[int][][]string{}
	index := 0

	// Intialize
	for _, path := range *paths {
		passed := false
		if len(res) == 0 {
			res[index] = append(res[index], path)
		} else {
			for i, road := range res {
				if !CheckIfExist(road, path) {
					passed = true
					res[i] = append(res[i], path)
				}
			}
			if !passed {
				index++
				res[index] = append(res[index], path)
			}
		}
	}

	// Fill Again
	for key, val := range res {
		for _, path := range *paths {
			if !CheckIfExist(val, path) {
				res[key] = append(res[key], path)
			}
		}
	}

	return res
}

func DeepEqual(path1 []string, path2 []string) bool {
	for _, room := range path1 {
		for _, room2 := range path2 {
			if room2 == room {
				return true
			}
		}
	}

	return false
}

func GetMinPath(paths []string) string {
	min := paths[0]
	minLen := len(strings.Split(paths[0], "\n"))

	for _, path := range paths {
		if minLen > len(strings.Split(path, "\n")) {
			minLen = len(strings.Split(path, "\n"))
			min = path
		}
	}

	return min
}
