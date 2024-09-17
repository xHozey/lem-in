package main

import (
	"fmt"

	ant "ants/Functions"
)

func main() {
	graphList := &ant.Graph{}
	data := ant.ReadFile()
	roomsAndLinks, start, end := ant.HandleData(data)
	linksFrom, linksTo, rooms := ant.GetRoomsAndLinks(roomsAndLinks)
	for i := 0; i < len(rooms); i++ {
		graphList.AddVertix(rooms[i])
	}
	for i := 0; i < len(linksFrom); i++ {
		graphList.AddIndirectedEdge(linksFrom[i], linksTo[i])
	}
	graphList.Dfs(start, end)
	ant.Travel()
	fmt.Println(ant.PossiblePaths)
}
