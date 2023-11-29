/*


Given an array of buildings and a direction that all of the buildings face, return an array of the indices of the buildings that can see the sunset.
A building can see the sunset if it's strictly taller than all of the buildings that come after it in the direction that it faces.
The input array named buildings contains positive, non-zero integers representing the heights of the buildings. A building at index i thus has a height denoted by buildings [i] . All of the buildings face the same direction, and this direction is either east or west, denoted by the input string named di rec ti on , which will always be equal to either "EAST" or "WEST" . In relation to the input array, you can interpret these directions as right for east and left for west.
Important note: the indices in the ouput array should be sorted in ascending order.

Sample Input
buildings =[3, 5, 4, 4, 3, 1, 3, 2)
direction = "EAST"

Sample Output

[1, 3, 6, 7]


This code iterates through the buildings from the back, checking if there are any taller buildings to the right (EAST) for EAST direction and to the left (WEST) for WEST direction. If there are no taller buildings, it means the building has a sunset view, and its index is added to the sunsetViewIndices array. Finally, the indices are sorted in ascending order and returned.




*/

package main

import (
	"fmt"
)

func buildingsWithSunsetView(buildings []int, direction string) []int {
	sunsetViewIndices := []int{}

	// Iterate through the buildings from the back
	for i := len(buildings) - 1; i >= 0; i-- {
		buildingHeight := buildings[i]

		// Check if there are any taller buildings to the right (EAST)
		if direction == "EAST" {
			tallerBuildingFound := false
			for j := i + 1; j < len(buildings); j++ {
				if buildings[j] > buildingHeight {
					tallerBuildingFound = true
					break
				}
			}

			if !tallerBuildingFound {
				sunsetViewIndices = append(sunsetViewIndices, i)
			}
		} else if direction == "WEST" {
			// Check if there are any taller buildings to the left (WEST)
			tallerBuildingFound := false
			for j := i - 1; j >= 0; j-- {
				if buildings[j] > buildingHeight {
					tallerBuildingFound = true
					break
				}
			}

			if !tallerBuildingFound {
				sunsetViewIndices = append(sunsetViewIndices, i)
			}
		} else {
			fmt.Println("Invalid direction:", direction)
			return []int{}
		}
	}

	// Sort the indices in ascending order
	sortedIndices := make([]int, len(sunsetViewIndices))
	copy(sortedIndices, sunsetViewIndices)
	for i := range sortedIndices {
		for j := i; j >= 0; j-- {
			if sortedIndices[j] > sortedIndices[i] {
				temp := sortedIndices[j]
				sortedIndices[j] = sortedIndices[i]
				sortedIndices[i] = temp
			}
		}
	}

	return sortedIndices
}

func main() {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"

	sunsetViewIndices := buildingsWithSunsetView(buildings, direction)
	fmt.Println("Buildings with sunset view:", sunsetViewIndices)
}
