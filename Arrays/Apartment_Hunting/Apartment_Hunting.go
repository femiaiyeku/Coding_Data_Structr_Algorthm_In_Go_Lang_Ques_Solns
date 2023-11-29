/*

You're looking to move into a new apartment on specific street,
and you're given a list of contiguous blocks on that street where each block contains an apartment that you could move into.
You also have a list of requirements: a list of buildings that are important to you.
For instance, you might value having a school and a gym near your apartment.
The list of blocks that you have contains information at every block about all of the buildings that are present and absent at the block in question. For instance,
for every block, you might know whether a school, a pool, an office, and a gym are present.
In order to optimize your life, you want to pick an apartment block such that you minimize the farthest distance you'd have to walk from your apartment to reach any of your required buildings.
Write a function that takes in a list of contiguous blocks on a specific street and a list of your required buildings and
that returns the location (the index) of the block that's most optimal for you.
If there are multiple most optimal blocks, your function can return the index of any one of them.

Sample Input
blocks = [
    {
        "gym": False,
        "school": True,
        "store": False,
        },
    {
        "gym": True,
        "school": False,
        "store": False,
        },
    {
        "gym": True,
        "school": True,
        "store": False,
        },
    {
        "gym": False,
        "school": True,
        "store": False,
        },
    {
        "gym": False,
        "school": True,
        "store": True,
        },
    ]
reqs = ["gym", "school", "store"]

Sample Output
3 // at index 3, the farthest you'd have to walk to reach a gym, a school, or a store is 1 block; at any other index, you'd have to walk farther


Hints
Hint 1

For every block, you want to go through every requirement, and for every requirement, you want to find the closest other block with that requirement (or rather, the smallest distance to another block with that requirement). Once you've done that for every requirement and for every block, you want to pick, for every block, the distance of the farthest requirement. You can do this with three nested "for" loops.

Hint 2

Is there a way to optimize on the solution mentioned in Hint #1 (that uses three nested "for" loops) by precomputing the smallest distances of every requirement from every block?

Hint 3

For every requirement, you should be able to precompute its smallest distances from every block by doing two simple passes though the array of blocks: one pass from left to right and one pass from right to left. Once you have these precomputed values, you can iterate through all of the blocks and pick the biggest of all the precomputed distances at that block.



To use the function, simply pass in the list of contiguous blocks and the list of required buildings as arguments. The function will return the index of the optimal block.

For example, using the sample input provided in the question, we would get the following output:

optimalBlockIndex := findOptimalApartmentBlock(blocks, reqs)
fmt.Println(optimalBlockIndex) // 3
This indicates that the optimal block is at index 3.


*/

package main

func findOptimalApartmentBlock(blocks []Block, reqs []string) int {

	// Create a map to store the farthest distance to each required building from each block.
	farthestDistances := make(map[int]int)

	// Iterate over each block.
	for i := 0; i < len(blocks); i++ {

		// Find the farthest distance to each required building from the current block.
		farthestDistance := 0
		for _, req := range reqs {
			if !blocks[i][req] {
				farthestDistance = Max(farthestDistance, findFarthestDistance(blocks, i, req))
			}
		}

		// Store the farthest distance to each required building from the current block in the map.
		farthestDistances[i] = farthestDistance
	}

	// Find the block with the smallest farthest distance to all required buildings.
	optimalBlockIndex := 0
	for i, farthestDistance := range farthestDistances {
		if farthestDistance < farthestDistances[optimalBlockIndex] {
			optimalBlockIndex = i
		}
	}

	// Return the index of the optimal block.
	return optimalBlockIndex
}

// Find the farthest distance to a given required building from a given block.
func findFarthestDistance(blocks []Block, blockIndex int, req string) int {

	// Initialize the farthest distance.
	farthestDistance := 0

	// Iterate over all subsequent blocks.
	for i := blockIndex + 1; i < len(blocks); i++ {
		if blocks[i][req] {
			// Found the required building.
			break
		}
		farthestDistance++
	}

	// Return the farthest distance.
	return farthestDistance
}

// Returns the maximum of two integers.

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Block map[string]bool
