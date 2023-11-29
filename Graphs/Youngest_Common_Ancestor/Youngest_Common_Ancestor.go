/*


You're given three inputs, all of which are instances of an Ancestral Tree class that have an ancestor property pointing to their youngest ancestor.
The first input is the top ancestor in an ancestral tree (i.e., the only instance that has no ancestor--its ancestor property points to None I null ),
and the other two inputs are descendants in the ancestral tree.
Write a function that returns the y)ungest common ancestor to the two descendants.
Note that a descendant is considered its own ancestor. So in the simple ancestral tree below,
the youngest common ancestor to nodes A and B is node A.


// The youngest common ancestor to nodes A and B is node A.
// The youngest common ancestor to nodes A and C is node A.
// The youngest common ancestor to nodes A and I is node H.
// The youngest common ancestor to nodes E and I is node H.
// The youngest common ancestor to nodes H and G is node H.
// The youngest common ancestor to nodes D and G is node A.
// This tree has five levels of depth.
// The youngest common ancestor to nodes A and E is node A.

Sample Input:

// The nodes are from the ancestral tree below.
topAncestor = node A
descendantOne = node E
descendantTwo = node I

             A
          /   \
         B     C
      /   \  /
     D     E F
    / \
      G   H I

Sample Output:

node B




*/

package main

import "fmt"

type Ancestor struct {
	ancestor *Ancestor
}

func findYoungestCommonAncestor(topAncestor, descendantOne, descendantTwo *Ancestor) *Ancestor {
	// Initialize a set to store the ancestors of descendantOne
	descendantOneAncestors := make(map[*Ancestor]bool)
	currentAncestor := descendantOne

	// Collect the ancestors of descendantOne
	for currentAncestor != nil {
		descendantOneAncestors[currentAncestor] = true
		currentAncestor = currentAncestor.ancestor
	}

	// Find the first ancestor of descendantTwo that is also an ancestor of descendantOne
	currentAncestor = descendantTwo

	// Check each ancestor of descendantTwo
	for currentAncestor != nil {
		if descendantOneAncestors[currentAncestor] {
			return currentAncestor // Youngest common ancestor found
		}
		currentAncestor = currentAncestor.ancestor
	}

	// If no common ancestor is found, return the top ancestor
	return topAncestor
}

func main() {
	// Create the ancestral tree
	topAncestor := &Ancestor{}
	descendantOne := &Ancestor{ancestor: &Ancestor{ancestor: &Ancestor{ancestor: topAncestor}}}
	descendantTwo := &Ancestor{ancestor: &Ancestor{ancestor: &Ancestor{ancestor: &Ancestor{ancestor: topAncestor}}}}

	// Find the youngest common ancestor
	youngestCommonAncestor := findYoungestCommonAncestor(topAncestor, descendantOne, descendantTwo)

	// Print the youngest common ancestor
	fmt.Println(youngestCommonAncestor)
}
