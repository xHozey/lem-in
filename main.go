package main

import (
	"fmt"

	ant "ants/Functions"
)

func main() {
	graphList := &ant.Graph{}
	ants, start, end := ant.GetData(graphList)
	fmt.Println(ants)
	graphList.PrintGraph()
	path := graphList.PathFinder(start, end)
	fmt.Println(path)
}
