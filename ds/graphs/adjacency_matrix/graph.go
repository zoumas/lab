package graph

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph implements a simple undirected graph.
// A simple graph has no loops or parallel edges.
type Graph struct {
	adjacencyMatrix [][]int
}

func NewGraph(vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		matrix[i] = make([]int, vertices)
	}

	return &Graph{adjacencyMatrix: matrix}
}

func NewGraphFromFile(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line, err := scanner.Text(), scanner.Err()
	if err != nil {
		return nil, err
	}

	vertices, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}
	g := NewGraph(vertices)

	for scanner.Scan() {
		line, err := scanner.Text(), scanner.Err()
		if err != nil {
			log.Printf("Scanner error: %q", err)
			return nil, err
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, errors.New("Bad file format: Need two vertices to form an edge")
		}

		v, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		u, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		g.AddEdge(v, u)
	}

	return g, nil
}

func (g Graph) outOfBounds(v int) bool {
	return v < 0 || v >= len(g.adjacencyMatrix)
}

func (g *Graph) AddEdge(v, u int) {
	if g.outOfBounds(v) || g.outOfBounds(u) {
		return
	}

	g.adjacencyMatrix[v][u] = 1
	g.adjacencyMatrix[u][v] = 1
}

func (g Graph) String() string {
	var b strings.Builder
	vertices := len(g.adjacencyMatrix)
	for i := 0; i < vertices-1; i++ {
		b.WriteString(fmt.Sprintln(g.adjacencyMatrix[i]))
	}
	b.WriteString(fmt.Sprint(g.adjacencyMatrix[vertices-1]))
	return b.String()
}
