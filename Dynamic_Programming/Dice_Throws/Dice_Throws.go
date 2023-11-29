/*

You're given a set of numDice dice, each with numSides sides, and a target integer, which represents a target
sum to obtain when rolling all of the dice and summing their values. Write a function that returns the total number of
dice-roll permutations that sum up to exactly that target value.
All three input values will always be positive integers. Each of the dice has an equal probability of landing on any
number from 1 to numSides. Identical total dice rolls obtained from different individual dice rolls (for example,
[2, 3] vs. [3, 2] ) count as different dice-roll permutations. If there's no possible dice-roll combination that sums
up to the target given the input dice, your function should return 0.


Sample Input


numDice = 2
numSides = 6
target = 7

Sample Output

6 // The following dice-roll permutations sum up to the target:





*/

// rollDice is a function that calculates the total number of dice-roll permutations that sum up to a target value.
// It takes three input parameters: numDice (the number of dice), numSides (the number of sides on each die), and target (the target sum).
// The function returns an integer representing the total number of valid dice-roll permutations.
// If there is no possible dice-roll combination that sums up to the target, the function returns 0.
// The function uses memoization to optimize the calculation by storing the number of valid dice rolls for each possible sum.
// It recursively calls the rollDiceHelper function to calculate the number of valid dice rolls for the target sum.
// The rollDiceHelper function is a helper function that performs the actual calculation.
// It checks if the number of valid dice rolls for the current sum has already been calculated and stored in the memoization table.
// If not, it rolls each die and recursively calculates the number of valid dice rolls for the remaining target sum.
// The count of valid dice rolls for the current sum is stored in the memoization table for future use.
// The rollDice function initializes the memoization table and calls the rollDiceHelper function to perform the calculation.
// Finally, it prints the total number of valid dice-roll permutations.

package main

import "fmt"

func rollDice(numDice, numSides, target int) int {
	// If the target is less than or equal to 0, there are no valid dice rolls.
	if target <= 0 {
		return 0
	}

	// If there is only one die, the only valid roll is the target value.
	if numDice == 1 {
		if target <= numSides {
			return 1
		} else {
			return 0
		}
	}

	// Create a memoization table to store the number of valid dice rolls for each possible sum.
	memo := make(map[int]int)

	// Recursively calculate the number of valid dice rolls for the target sum.
	return rollDiceHelper(memo, numDice, numSides, target)
}

func rollDiceHelper(memo map[int]int, numDice, numSides, target int) int {
	// If the target sum is 0 and there are no more dice to roll, we have found a valid dice roll.
	if target == 0 && numDice == 0 {
		return 1
	}

	// If there are no more dice to roll and the target sum is not 0, there are no valid dice rolls.
	if numDice == 0 {
		return 0
	}

	// Check if the number of valid dice rolls for the current sum has already been calculated.
	if count, ok := memo[target]; ok {
		return count
	}

	// Initialize the count of valid dice rolls for the current sum.
	count := 0

	// Roll each die and recursively calculate the number of valid dice rolls for the remaining target sum.
	for i := 1; i <= numSides; i++ {
		count += rollDiceHelper(memo, numDice-1, numSides, target-i)
	}

	// Store the number of valid dice rolls for the current sum in the memoization table.
	memo[target] = count

	// Return the number of valid dice rolls for the current sum.
	return count
}

func main() {
	// Set the input values.
	numDice := 2
	numSides := 6
	target := 7

	// Calculate the total number of valid dice-roll permutations that sum up to the target value.
	count := rollDice(numDice, numSides, target)

	// Print the total number of valid dice-roll permutations.
	fmt.Println(count)
}
