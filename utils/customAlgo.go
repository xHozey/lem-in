package utils

import (
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
		// reintialze the paths when ants are arrived
		*dup = Scoring(paths)
		RateSort(paths, dup)
		step++
	}

	return road
}

func AntsGoing(ant int, paths *[][]string, fillRoom *map[string][]int, antsArrived *map[int]bool, road *map[int]Road, tunnels *map[string][]int, dup *Scoretype, step int) {
	rooms := map[string]int{}

	path := (*paths)[0]
	pathIndex := 0

	for roomIndex, room := range path {

		// Check if tunnel exist
		if roomIndex+1 < len(path) {
			tunnelKey := room + "-" + path[roomIndex+1]
			isTunnelExist := slices.Contains((*tunnels)[tunnelKey], step+roomIndex)
			if isTunnelExist {
				(*dup)[pathIndex].score++

				RateSort(paths, dup)
				break
			} else {
				(*tunnels)[tunnelKey] = append((*tunnels)[tunnelKey], step+roomIndex)
			}
		}

		// Check if room is fill
		isRoomFill := slices.Contains((*fillRoom)[room], step+roomIndex)
		if isRoomFill {
			break
		}

		// mark room to show that is fill in a specific step
		if roomIndex != len(path)-1 && roomIndex != 0 {
			rooms[room] = step + roomIndex
		}

		// if ant arrived
		// add it to the path length
		// add the road and the fill rooms to the main variables
		if roomIndex == len(path)-1 {
			(*dup)[pathIndex].score++

			RateSort(paths, dup)

			(*road)[ant] = Road{TheRoad: path, Step: step}

			(*antsArrived)[ant] = true

			for r, s := range rooms {
				(*fillRoom)[r] = append((*fillRoom)[r], s)
			}

			return
		}
	}
}
