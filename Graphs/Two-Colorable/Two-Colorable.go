/*


You're given a list of edges representing a connected, unweighted, undirected graph with at least one node. Write a
function that returns a boolean representing whether the given graph is two-colorable.
A graph is two-colorable (also called bipartite) if all of the nodes can be assigned one of two colors such that no nodes
of the same color are connected by an edge.
The given list is what's called an adjacency list, and it represents a graph. The number of vertices in the graph is equal to
the length of edges, where each index i in edges contains vertex's siblings, in no particular order. Each
individual edge is represented by a positive integer that denotes an index in the list that this vertex is connected to.
Note that this graph is undirected, meaning that if a vertex appears in the edge list of another vertex, then the inverse
will also be true.
Also note that this graph may contain self-loops. A self-loop is an edge that has the same destination and origin; in
other words, it's an edge that connects a vertex to itself. Any self-loop should make a graph not 2-colorable.

Sample Input:

edges = [
    [1, 3],
    [0, 2],
    [1, 3],
    [0, 2]
]

Sample Output:

true

// The following graph is two-colorable:
// 0: R
// 1: B
// 2: R
// 3: B
// Every vertex can have a color that is different than the color of its immediate neighbors.







*/

package main

import (
	"fmt"
)

func isBipartite(edges [][]int) bool {
	// Create a color map to store the color of each node
	colorMap := make(map[int]int, len(edges))

	// Iterate through the nodes and assign colors
	for node := 0; node < len(edges); node++ {
		if colorMap[node] == 0 {
			if !dfs(edges, node, 1, colorMap) {
				return false // Graph is not bipartite
			}
		}
	}

	return true // Graph is bipartite
}

func dfs(edges [][]int, node int, currentColor int, colorMap map[int]int) bool {
	// Assign the current color to the current node
	colorMap[node] = currentColor

	// Check each neighbor of the current node
	for _, neighbor := range edges[node] {
		if neighbor == node { // Self-loop detected
			return false
		}

		if colorMap[neighbor] == 0 { // Unvisited neighbor
			if !dfs(edges, neighbor, -currentColor, colorMap) {
				return false
			}
		} else if colorMap[neighbor] != -currentColor { // Visited neighbor with different color
			return false
		}
	}

	return true
}

func main() {
	edges := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}

	isTwoColorable := isBipartite(edges)
	fmt.Println(isTwoColorable)
}
