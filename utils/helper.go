package utils

func GetVertex(vertecies []*Vertex, key string) *Vertex {
	for _, vertex := range vertecies {
		if vertex.Key == key {
			return vertex
		}
	}

	return nil
}

func Printer(ant int, paths map[int][]string) {
	roads := []string{}

	for key, path := range paths {
		line := 0
		for _, room := range path {
			if len(roads) > line {
				// add room to the index
			} else {
				// append
			}
			line++
		}
	}
}

func PathToScorePath(paths *[][]string) map[int]int {
	res := map[int]int{}
	for i, path := range *paths {
		res[i] = len(path)
	}

	return res
}

func GetMinPath(pathsScore *map[int]int, paths *[][]string) (int, []string) {
	min := 0
	minVal := (*pathsScore)[0]

	for key, path := range *pathsScore {
		if minVal >= path {
			minVal = path
			min = key
		}
	}

	return min, (*paths)[min]
}
