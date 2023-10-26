package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/lab/ds/graphs/directed_graph"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	filename := os.Args[1]

	g, err := graph.NewGraphFromFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(g)
}
