/*


You're given an array of non-negative integers where each non-zero integer represents the height of a pillar of width 1 . 
Imagine water being poured over all of the pillars; write a function that returns the surface area of the water trapped between the pillars viewed from the front. Note that spilled water should be ignored. 
You can refer to the first three minutes of this question's video explanation for a visual example. 


Sample Input

heights = [0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3]

Sample Output

48



Hints
Hint 1

In order to calculate the amount of water above a single point in the input array, you must know the height of the tallest pillar to its left and the height of the tallest pillar to its right.

Hint 2

If a point can hold water above it, then the smallest of the two heights mentioned in Hint #1 minus the height at that respective point should lead you to the amount of water above it.

Hint 3

Try building an array of the left and right max heights for each point in the input array. You should be able to build this array and to compute the final amount of water above each point in just two loops over the input array.





*/



package main

import (
    "fmt"
)

func calculateWaterTrapped(heights []int) int {
    // Initialize two pointers, left and right, to the beginning and end of the array, respectively
    left := 0
    right := len(heights) - 1

    // Initialize the maximum left and right heights
    maxLeftHeight := heights[left]
    maxRightHeight := heights[right]

    // Initialize the total trapped water
    totalTrappedWater := 0

    // Continue iterating until the pointers meet
    while left < right {
        // Check if the current left height is smaller than the maximum left height
        if heights[left] < maxLeftHeight {
            // Calculate the water trapped at the current left position
            currentWaterTrapped := (maxLeftHeight - heights[left]) * (right - left)

            // Update the total trapped water
            totalTrappedWater += currentWaterTrapped

            // Update the current left position
            left++
        } else {
            // Update the maximum left height
            maxLeftHeight = heights[left]
        }

        // Check if the current right height is smaller than the maximum right height
        if heights[right] < maxRightHeight {
            // Calculate the water trapped at the current right position
            currentWaterTrapped := (maxRightHeight - heights[right]) * (right - left)

            // Update the total trapped water
            totalTrappedWater += currentWaterTrapped

            // Update the current right position
            right--
        } else {
            // Update the maximum right height
            maxRightHeight = heights[right]
        }
    }

    // Return the total trapped water
    return totalTrappedWater
}

func main() {
    // Set the input heights.
    heights := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}

    // Calculate the surface area of the water trapped between the pillars.
    surfaceArea := calculateWaterTrapped(heights)

    // Print the surface area.
    fmt.Println(surfaceArea)
}
