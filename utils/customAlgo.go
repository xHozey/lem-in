package utils

import (
	"fmt"
	"slices"
)

type Road struct {
	TheRoad []string
	Step    int
}

type Step struct {
	Turn      int
	RoomIndex int
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

func GoTo(paths *[][]string, dup *Scoretype, ants int) map[int]Road {
	road := map[int]Road{}
	antsCount := ants
	antsArrived := map[int]bool{}
	fillRoom := map[string][]int{}
	tunnels := map[string][]int{}

	step := 1
	for antsCount > 0 {
		for i := 1; i <= ants; i++ {
			if antsArrived[i] {
				antsCount = ants - len(antsArrived)
				continue
			}
			AntsGoing(i, paths, &fillRoom, &antsArrived, &road, &tunnels, dup, step)
		}
		*dup = Duplicated(paths)
		Sort(paths, dup)
		step++
	}

	return road
}

func AntsGoing(ant int, paths *[][]string, fillRoom *map[string][]int, antsArrived *map[int]bool, road *map[int]Road, tunnels *map[string][]int, dup *Scoretype, step int) {
	rooms := map[string]int{}

	path := (*paths)[0]
	pathIndex := 0
	/* (*fillPath)[pIndex] = true */

	// Clean up the road

	roadOfAnt := []string{}

	for roomIndex, room := range path {

		isRoomFill := slices.Contains((*fillRoom)[room], step+roomIndex)

		if roomIndex+1 < len(path) {
			tunnelKey := room + "-" + path[roomIndex+1]
			isTunnelExist := slices.Contains((*tunnels)[tunnelKey], step+roomIndex)
			if isTunnelExist {
				break
			} else {
				(*tunnels)[tunnelKey] = append((*tunnels)[tunnelKey], step+roomIndex)
			}
		}

		if isRoomFill {
			break
		}
		if roomIndex != len(path)-1 && roomIndex != 0 {
			rooms[room] = step + roomIndex
		}

		roadOfAnt = append(roadOfAnt, room)

		if roomIndex == len(path)-1 {
			(*dup)[pathIndex].score++

			Sort(paths, dup)

			fmt.Println("Dup")
			for _, path := range *dup {
				fmt.Println(path)
			}
			fmt.Println("##################################")

			(*road)[ant] = Road{TheRoad: roadOfAnt, Step: step}

			(*antsArrived)[ant] = true

			for r, s := range rooms {
				(*fillRoom)[r] = append((*fillRoom)[r], s)
			}

			return
		}
	}
}
