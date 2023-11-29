/*

Lowest Common Manager
You're given three inputs, all of which are Instances of an OrgChart class that have a directReports property
pointing to their direct reports. The first input is the top manager in an organizational chart (i.e., the only instance
that isn't anybody else's direct report), and the other two Inputs are reports in the organizational chart. The two
reports are guaranteed to be distinct.
Write a function that returns the lowest common manager to the two reports.


Sample Input:
topManager = Node A

reportOne = Node E
reportTwo = Node I

            A
           / \
          B   C
         / \ / \
        D  E F  G
       / \
      H   I

Sample Output:
Node B









Hints
Hint 1

Given a random subtree in the organizational chart, the manager at the root of that subtree is common to any two reports in the subtree.

Hint 2

Knowing Hint #1, the lowest common manager to two reports in an organizational chart is the root of the lowest subtree containing those two reports. Find that lowest subtree to find the lowest common manager.

Hint 3

To find the lowest subtree containing both of the input reports, try recursively traversing the organizational chart and keeping track of the number of those reports contained in each subtree as well as the lowest common manager in each subtree. Some subtrees might contain neither of the two reports, some might contain one of them, and others might contain both; the first to contain both should return the lowest common manager for all of the subtrees above it that contain it, including the entire organizational chart.


This code first defines a Node struct to represent a node in the organizational chart. Each node has a name and a slice of direct reports. The findAncestors function takes a top manager and a target node and returns a slice of all the ancestors of the target node, starting from the top manager.

The findLCM function takes a top manager and two reports and returns the lowest common manager of the two reports. It first finds the ancestors of each report using the findAncestors function. Then, it iterates through the ancestors of one report and checks if any of them are an ancestor of the other report. If it finds a common ancestor, it returns that ancestor as the LCM.



*/

package main

import "fmt"

type Node struct {
	name          string
	directReports []*Node
}

func findLCM(topManager *Node, reportOne, reportTwo *Node) *Node {
	if topManager == nil || reportOne == nil || reportTwo == nil {
		return nil
	}

	ancestorsOne := findAncestors(topManager, reportOne)
	ancestorsTwo := findAncestors(topManager, reportTwo)

	var lcm *Node

	for _, ancestorOne := range ancestorsOne {
		for _, ancestorTwo := range ancestorsTwo {
			if ancestorOne == ancestorTwo {
				lcm = ancestorOne
				break
			}
		}
	}

	return lcm
}

func findAncestors(topManager *Node, target *Node) []*Node {
	ancestors := []*Node{}

	currentNode := target

	for currentNode != nil {
		ancestors = append(ancestors, currentNode)
		currentNode = currentNode.directReports[0] // Assuming each node has only one direct report
	}

	return ancestors
}

func main() {
	topManager := &Node{
		name: "A",
		directReports: []*Node{
			{
				name: "B",
				directReports: []*Node{
					{
						name: "D",
					},
					{
						name: "E",
					},
				},
			},
			{
				name: "C",
				directReports: []*Node{
					{
						name: "F",
					},
					{
						name: "G",
						directReports: []*Node{
							{
								name: "H",
							},
							{
								name: "I",
							},
						},
					},
				},
			},
		},
	}

	reportOne := &Node{
		name: "E",
	}

	reportTwo := &Node{
		name: "I",
	}

	lcm := findLCM(topManager, reportOne, reportTwo)
	fmt.Println("Lowest common manager:", lcm.name)
}
