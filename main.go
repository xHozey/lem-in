package main

import (
	"fmt"
	"os"

	"lem/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || !utils.CheckFileName(args[0]) {
		fmt.Fprintln(os.Stderr, "Usage: go run . filename.txt")
		return
	}

	fileData, err := utils.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please Write a valid file")
		return
	}

	parsedData, err := utils.ParseData(fileData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please Write a valid file")
		return
	}

	graph := utils.Graph{}

	for _, room := range parsedData.Rooms {
		graph.AddVertex(room)
	}

	for _, links := range parsedData.Links {
		graph.AddEdge(links[0], links[1])
	}

	paths := [][]string{}
	startVertex := utils.GetVertex(graph.Vertecies, parsedData.Start)
	targetVertex := utils.GetVertex(graph.Vertecies, parsedData.End)

	graph.CDFS(startVertex, targetVertex, []string{}, &paths)

	dup := utils.Duplicated(&paths)
	utils.RateSort(&paths, &dup)

	fmt.Println("Paths")
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println("##################################")

	fmt.Println("Dup")
	for i, path := range dup {
		if i == 6 {
			break
		}
		fmt.Println(path)
	}
	fmt.Println("##################################")

	res := utils.GoTo(&paths, &dup, parsedData.Ants)

	strRes := utils.Printer(res)
	fmt.Print(strRes)

	fmt.Println("##############################")
	rres := utils.SepRoads(&paths)
	fmt.Println(rres)
}
