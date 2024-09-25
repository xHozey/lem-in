package main

import (
	"fmt"
	"strings"

	ant "ants/Functions"
)

func main() {
	graphList := &ant.Graph{}
	ants, start, end, fileData := ant.GetData(graphList)
	path := graphList.PathFinder(start, end)
	var results []string
	for _, combo := range path {
		score := ant.GiveScore(&combo)
		ant.SortByScore(&combo, score)
		res := ant.GoTo(&combo, &score, ants)
		strRes := ant.Printer(res)
		results = append(results, strRes)
	}

	if len(results) == 0 {
		fmt.Println("Error: No Reslult possible")
		return
	}
	fmt.Println(fileData + "\n")
	fmt.Println(strings.TrimSpace(ant.GetMinPath(results)))
}
