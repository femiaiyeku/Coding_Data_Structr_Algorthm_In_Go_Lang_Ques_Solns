/*


You're given a Node class that has a name and an array of optional children nodes. When put together, nodes
form an acyclic tree-like structure.
Implement the breadthFirstSearch method on the Node class, which takes in an empty array, traverses the tree
using the Breadth-first Search approach (specifically navigating the tree from left to right), stores all of the nodes' names
in the input array, and returns it.
If you're unfamiliar with Breadth-first Search, we recommend watching the Conceptual Overview section of this
question's video explanation before starting to code.

Sample Input:

graph = A
         / | \
        B C D
      / \ / \
        E F G H
        / \ \
        I J K

Sample Output:

["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"]








*/

package main

import "fmt"

type Node struct {
	name     string
	children []Node
}

func (node *Node) breadthFirstSearch(visited []string) []string {
	queue := []Node{node}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if !contains(visited, currentNode.name) {
			visited = append(visited, currentNode.name)

			for _, child := range currentNode.children {
				queue = append(queue, child)
			}
		}
	}

	return visited
}

func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}

	return false
}

func main() {
	root := &Node{name: "A"}
	root.children = []Node{
		{name: "B", children: []Node{{name: "E"}, {name: "F"}}},
		{name: "C", children: []Node{{name: "G"}, {name: "H"}}},
		{name: "D", children: []Node{{name: "I"}, {name: "J"}, {name: "K"}}},
	}

	visited := []string{}
	result := root.breadthFirstSearch(visited)
	fmt.Println(result)
}
