package utils

import "fmt"

var RoomsCounter = map[string]int{}

type Graph struct {
	Vertecies []*Vertex
}

type Vertex struct {
	Key        string
	Adjacments []*Vertex
	Visted     bool
}

func (g *Graph) AddVertex(key string) {
	if GetVertex(g.Vertecies, key) == nil {
		g.Vertecies = append(g.Vertecies, &Vertex{Key: key})
	} else {
		fmt.Printf("Vertex %s already exist!!\n", key)
	}
}

func (g *Graph) AddEdge(from string, to string) {
	fromVertex := GetVertex(g.Vertecies, from)
	toVertex := GetVertex(g.Vertecies, to)

	if fromVertex == nil {
		fmt.Printf("Vertex %s doesn't exist!!\n", from)
		return
	}

	if toVertex == nil {
		fmt.Printf("Vertex %s doesn't exist!!\n", to)
		return
	}

	if GetVertex(fromVertex.Adjacments, to) != nil || GetVertex(toVertex.Adjacments, from) != nil {
		fmt.Printf("Edge %s <--> %s already exist!!", from, to)
		return
	}

	RoomsCounter[from]++
	RoomsCounter[to]++
	fromVertex.Adjacments = append(fromVertex.Adjacments, toVertex)
	toVertex.Adjacments = append(toVertex.Adjacments, fromVertex)
}
