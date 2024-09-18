package utils

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

func GoTo(paths [][]string, ants int) map[int][]string {
	// slice of path -> ants
	roads := map[int][]string{}

	// score path map
	scoredPaths := PathToScorePath(&paths)

	for i := 1; i <= ants; i++ {
		index, minPath := GetMinPath(&scoredPaths, &paths)
		scoredPaths[index]++
		roads[i] = minPath[1:]
	}

	return roads
}
