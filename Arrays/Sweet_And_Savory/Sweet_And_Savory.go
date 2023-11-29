/*


You're hosting an event at a food festival and want to showcase the best possible pairing of two dishes from the festival that
complement each other's flavor profile.
Each dish has a flavor profile represented by an integer. A negative integer means a dish is sweet, while a positive integer
means a dish is savory. The absolute value of that integer represents the intensity of that flavor. For example, a flavor profile
of -3 is slightly sweet, one of -10 is extremely sweet, one of 2 is mildly savory, and one of 8 is significantly savory.
You're given an array of these dishes and a target combined flavor profile. Write a function that returns the best possible
pairing of two dishes (the pairing with a total flavor profile that's closest to the target one). Note that this pairing must include
one sweet and one savory dish. You're also concerned about the dish being too savory, so your pairing should never be more
savory than the target flavor profile.

All dishes will have a positive or negative flavor profile; there are no dishes with a 0 value. For simplicity, you can assume that
there will be at most one best solution. If there isn't a valid solution, your function should return [[0, 0]. The returned array
should be sorted, meaning the sweet dish should always come first.

Sample Input

dishes = [-3, -5, 1, 7]
target = 8

Sample Output
[-3, 7] // The dishes could be in reverse order



This function works by first sorting the dishes in ascending order. This is necessary to ensure that the function can efficiently find the best dish pair. The function then initializes the best dish pair and the minimum absolute difference between the target flavor profile and the flavor profile of a dish pair.

The function then iterates over the dishes. For each sweet dish, the function iterates over the dishes again, starting from the index of the sweet dish. The function calculates the total flavor profile of the dish pair and the absolute difference between the target flavor profile and the flavor profile of the dish pair. If the absolute difference is smaller than the minimum absolute difference, the function updates the best dish pair and the minimum absolute difference.

Once the function has finished iterating over the dishes, it returns the best dish pair.



*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func findBestDishPair(dishes []int, target int) []int {
	// Sort the dishes in ascending order.
	sort.Ints(dishes)

	// Initialize the best dish pair.
	bestDishPair := []int{0, 0}

	// Initialize the minimum absolute difference between the target flavor profile and the flavor profile of a dish pair.
	minAbsoluteDifference := math.Inf(1)

	// Iterate over the dishes.
	for i, sweetDish := range dishes {
		// If the sweet dish is more savory than the target flavor profile, break from the loop.
		if sweetDish > target {
			break
		}

		// Iterate over the dishes again, starting from the index of the sweet dish.
		for j := i + 1; j < len(dishes); j++ {
			// Calculate the total flavor profile of the dish pair.
			totalFlavorProfile := sweetDish + dishes[j]

			// Calculate the absolute difference between the target flavor profile and the flavor profile of the dish pair.
			absoluteDifference := int(math.Abs(float64(totalFlavorProfile - target)))

			// If the absolute difference is smaller than the minimum absolute difference, update the best dish pair and the minimum absolute difference.
			if absoluteDifference < minAbsoluteDifference {
				bestDishPair[0] = sweetDish
				bestDishPair[1] = dishes[j]
				minAbsoluteDifference = absoluteDifference
			}
		}
	}

	// Return the best dish pair.
	return bestDishPair
}

func main() {
	dishes := []int{-3, -5, 1, 7}
	target := 8

	bestDishPair := findBestDishPair(dishes, target)

	fmt.Println(bestDishPair)
}
