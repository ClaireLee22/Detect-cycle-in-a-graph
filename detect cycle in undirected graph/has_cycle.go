package main

import "fmt"

func hasCycle(numNodes int, edges [][]int) bool {
	graph := buildGraph(numNodes, edges)
	fmt.Printf("adjacency list: %+v\n\n", graph)

	visited := map[int]bool{}

	for node := 0; node < numNodes; node += 1 {
		fmt.Println("------------------------------------------")
		fmt.Println("dfs node ", node)
		if detectCycle(graph, node, visited, -1) {
			return false
		}
	}

	return true
}

func detectCycle(
	graph map[int][]int,
	node int,
	visited map[int]bool,
	parent int) bool {
	if _, found := visited[node]; found {
		fmt.Printf("node %d has been visited\n\n", node)
		return true
	}

	visited[node] = true
	for _, descendant := range graph[node] {
		fmt.Printf("current node: %d\n", node)
		fmt.Printf("visited: %+v\n", visited)
		fmt.Printf("parent: %d\n", parent)
		if descendant != parent {
			fmt.Printf("visited descendant: node %d\n\n", descendant)
		} else {
			fmt.Printf("not visited descendant, node %d. It is parent.\n\n", descendant)
		}
		if descendant != parent && detectCycle(graph, descendant, visited, node) {
			return true
		}
	}
	return false
}

func buildGraph(numNodes int, prerequisites [][]int) map[int][]int {
	graph := map[int][]int{}
	for i := 0; i < numNodes; i++ {
		graph[i] = []int{}
	}

	for _, prerequisite := range prerequisites {
		a, b := prerequisite[0], prerequisite[1]
		graph[b] = append(graph[b], a)
		graph[a] = append(graph[a], b)
	}

	return graph
}

func main() {
	numCourses := 6
	prerequisites := [][]int{
		{1, 0},
		{3, 0},
		{5, 0},
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 4},
		{3, 5},
	}
	fmt.Println("has cycle?", hasCycle(numCourses, prerequisites))
}

/*
output:
adjacency list: map[0:[1 3 5] 1:[0 2 3] 2:[1] 3:[0 1 4 5] 4:[3 5] 5:[0 4 3]]

------------------------------------------
dfs node  0
current node: 0
visited: map[0:true]
parent: -1
visited descendant: node 1

current node: 1
visited: map[0:true 1:true]
parent: 0
not visited descendant, node 0. It is parent.

current node: 1
visited: map[0:true 1:true]
parent: 0
visited descendant: node 2

current node: 2
visited: map[0:true 1:true 2:true]
parent: 1
not visited descendant, node 1. It is parent.

current node: 1
visited: map[0:true 1:true 2:true]
parent: 0
visited descendant: node 3

current node: 3
visited: map[0:true 1:true 2:true 3:true]
parent: 1
visited descendant: node 0

node 0 has been visited

has cycle? false
*/
