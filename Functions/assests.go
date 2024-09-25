package ants

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func GiveScore(paths *[][]string) []int {
	score := make([]int, len(*paths))
	for i, path := range *paths {
		score[i] = len(path)
	}
	return score
}

func SortByScore(paths *[][]string, scores []int) {
	sort.Slice(*paths, func(i, j int) bool {
		return scores[i] < scores[j]
	})
}

func GoTo(paths *[][]string, dup *[]int, ants int) map[int]Road {
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

		*dup = GiveScore(paths)
		SortByScore(paths, *dup)
		step++
	}
	return road
}

func AntsGoing(ant int, paths *[][]string, fillRoom *map[string][]int, antsArrived *map[int]bool, road *map[int]Road, tunnels *map[string][]int, dup *[]int, step int) {
	rooms := map[string]int{}

	path := (*paths)[0]
	pathIndex := 0

	for roomIndex, room := range path {

		// Check if tunnel exist
		if roomIndex+1 < len(path) {
			tunnelKey := room + "-" + path[roomIndex+1]
			isTunnelExist := slices.Contains((*tunnels)[tunnelKey], step+roomIndex)
			if isTunnelExist {
				(*dup)[pathIndex]++
				SortByScore(paths, *dup)
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
			(*dup)[pathIndex]++
			SortByScore(paths, *dup)

			(*road)[ant] = Road{Path: path, Step: step}

			(*antsArrived)[ant] = true

			for r, s := range rooms {
				(*fillRoom)[r] = append((*fillRoom)[r], s)
			}

			return
		}
	}
}

func Printer(paths map[int]Road) string {
	res := map[int][]string{}
	str := ""
	for ant, path := range paths {
		for i, room := range path.Path[1:] {
			p := fmt.Sprintf("L%d-%s", ant, room)
			res[i+path.Step-1] = append(res[i+path.Step-1], p)
		}
	}

	for i := 0; i < len(res); i++ {
		str += strings.Join(res[i], " ") + "\n"
	}

	return str
}

func GetMinPath(paths []string) string {
	min := paths[0]
	mins := []string{}
	minLen := len(strings.Split(paths[0], "\n"))

	for _, path := range paths {
		if minLen >= len(strings.Split(path, "\n")) {
			minLen = len(strings.Split(path, "\n"))
			min = path
			mins = append(mins, path)
		}
	}

	minWords := len(strings.Split(mins[0], " "))
	for _, path := range mins {
		if minWords > len(strings.Split(path, " ")) {
			minWords = len(strings.Split(path, " "))
			min = path
		}
	}

	return min
}
