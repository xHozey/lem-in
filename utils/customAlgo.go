package utils

import (
	"fmt"
	"slices"
)

type Road struct {
	TheRoad []string
	Step    int
}

func (g *Graph) CDFS(source *Vertex, target *Vertex, path []string, paths *[][]string) {
	source.Visted = true
	path = append(path, source.Key)

	if source.Key == target.Key {
		temp := make([]string, len(path))
		copy(temp, path)
		*paths = append(*paths, temp)
	} else {
		for _, v := range source.Adjacments {
			if !v.Visted {
				g.CDFS(v, target, path, paths)
			}
		}
	}

	source.Visted = false
}

func GoTo(paths [][]string, ants int) map[int]Road {
	road := map[int]Road{}
	antsCount := ants
	antsArrived := map[int]bool{}
	fillRoom := map[string][]int{}

	step := 1
	for antsCount > 0 {
		fmt.Println(antsCount)
		for i := 1; i <= ants; i++ {
			if antsArrived[i] {
				antsCount--
				continue
			}

			AntsGoing(i, &paths, &fillRoom, &antsArrived, &road, step)
		}
		step++
	}

	return road
}

func AntsGoing(ant int, paths *[][]string, fillRoom *map[string][]int, antsArrived *map[int]bool, road *map[int]Road, step int) {
	rooms := map[string]int{}
	for _, path := range *paths {
		/* (*fillPath)[pIndex] = true */

		// Clean up the road

		roadOfAnt := []string{}

		for roomIndex, room := range path {

			isRoomFill := slices.Contains((*fillRoom)[room], step)

			if isRoomFill {
				break
			}
			if roomIndex != len(path)-1 && roomIndex != 0 {
				rooms[room] = step
			}

			roadOfAnt = append(roadOfAnt, room)

			if roomIndex == len(path)-1 {

				(*road)[ant] = Road{TheRoad: roadOfAnt, Step: step}

				(*antsArrived)[ant] = true

				for r, s := range rooms {
					(*fillRoom)[r] = append((*fillRoom)[r], s)
				}

				return
			}
		}
	}
}
