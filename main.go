package main

import (
	ant "ants/Functions"
)

var (
	GraphList       = &ant.Graph{}
	VerticesLength  = len(GraphList.Vertices)
	VisitedVertices = make(map[string]bool)
)

func main() {
	data := ant.ReadFile()
	roomsAndLinks := ant.HandleData(data)
	linksFrom, linksTo, rooms := ant.GetRoomsAndLinks(roomsAndLinks)
	for i := 0; i < len(rooms); i++ {
		GraphList.AddVertix(rooms[i])
	}
	for i := 0; i < len(linksFrom); i++ {
		GraphList.AddIndirectedEdge(linksFrom[i], linksTo[i])
	}
	GraphList.PrintGraph()
	search("1")
}

func search(at string) string {
	if VisitedVertices[at] {
		return search(at)
	}
	VisitedVertices[at] = true

	return ""
}
