package main

import (
	"fmt"
	"os"

	"lem/utils"
)

func main() {
	// Check args
	args := os.Args[1:]
	if len(args) != 1 || !utils.CheckFileName(args[0]) {
		fmt.Fprintln(os.Stderr, "Usage: go run . filename.txt")
		return
	}

	// Read File And parse data
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

	// Intialize Graph
	graph := utils.Graph{}

	for _, room := range parsedData.Rooms {
		graph.AddVertex(room)
	}

	for _, links := range parsedData.Links {
		graph.AddEdge(links[0], links[1])
	}

	// Get All possible paths
	paths := [][]string{}
	startVertex := utils.GetVertex(graph.Vertecies, parsedData.Start)
	targetVertex := utils.GetVertex(graph.Vertecies, parsedData.End)

	graph.CDFS(startVertex, targetVertex, []string{}, &paths)
	utils.Sort(&paths)

	// Sepearate Paths in unique combinations
	disjointPaths := utils.SepRoads(&paths)

	// Get All Possible Results and print the most optimal
	results := []string{}

	for _, combPaths := range disjointPaths {
		dup := utils.Duplicated(&combPaths)
		utils.RateSort(&combPaths, &dup)

		res := utils.GoTo(&combPaths, &dup, parsedData.Ants)

		strRes := utils.Printer(res)
		results = append(results, strRes)
	}

	fmt.Print(utils.GetMinPath(results))
}
