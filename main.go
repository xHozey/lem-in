package main

import (
	ant "ants/Functions"
)

func main() {
	graph := &ant.Graph{}
	data := ant.ReadFile()
	roomsAndLinks := ant.HandleData(data)
	linksFrom, linksTo, rooms := ant.GetRoomsAndLinks(roomsAndLinks)
	for i := 0; i < len(rooms); i++ {
		graph.AddVertix(rooms[i])
	}
	for i := 0; i < len(linksFrom); i++ {
		graph.AddIndirectedEdge(linksFrom[i], linksTo[i])
	}
	graph.PrintGraph()
}
