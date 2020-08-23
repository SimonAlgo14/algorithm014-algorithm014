package main

import "fmt"

type Node  map[string] bool

type Graph struct {
	nodes map[string]Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]Node),
	}
}

func (g *Graph) AddNode(name string) {
	if !g.ContainsNode(name) {
		g.nodes[name] = make(Node)
	}
}

func (g *Graph) AddEdge(from string, to string) error {
	f, ok := g.nodes[from]
	if !ok {
		return fmt.Errorf("node %q not found", from)
	}
	_, ok = g.nodes[to]
	if !ok {
		return fmt.Errorf("node %q not found", to)
	}

	f.addEdge(to)
	return nil
}

func (g *Graph) ContainsNode(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (n Node) addEdge(name string) {
	n[name] = true
}

func (n Node) edges() []string {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}