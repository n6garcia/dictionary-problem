package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		fmt.Println("vertex already exists!")
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func contains(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from string, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("cannot add edge")
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k string) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) deleteVertex(k string) {
	delIndex := g.getVertexIndex(k)

	if delIndex == -1 {
		fmt.Println("vertex doesn't exist!")
	} else {
		// delete vertex
		g.vertices = remove(g.vertices, delIndex)

		// delete edges
		for _, v := range g.vertices {
			for i, u := range v.adjacent {
				if u.key == k {
					v.adjacent = remove(v.adjacent, i)
				}
			}
		}
	}
}

func (g *Graph) getVertexIndex(k string) int {
	for i, v := range g.vertices {
		if v.key == k {
			return i
		}
	}
	return -1
}

func remove(s []*Vertex, i int) []*Vertex {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex: %s", v.key)
		fmt.Printf(" outEdges: ")
		for _, v := range v.adjacent {
			fmt.Printf(" %s ", v.key)
		}
	}
}

func main() {
	fmt.Println("starting!")

	tGraph := &Graph{}

	for i := 0; i < 5; i++ {
		tGraph.AddVertex(strconv.Itoa(i))
	}

	tGraph.AddEdge(strconv.Itoa(1), strconv.Itoa(3))
	tGraph.AddEdge(strconv.Itoa(1), strconv.Itoa(2))
	tGraph.AddEdge(strconv.Itoa(6), strconv.Itoa(2))
	tGraph.AddEdge(strconv.Itoa(4), strconv.Itoa(3))

	tGraph.deleteVertex(strconv.Itoa(0))
	tGraph.deleteVertex(strconv.Itoa(6))
	tGraph.deleteVertex(strconv.Itoa(2))

	tGraph.Print()
}
