/*

This problem deals with measuring cups that are missing important measuring labels. Specifically, a measuring cup only has two measuring lines, a Low {L) line and a High {H) line. This means that these cups can't precisely measure and can only guarantee that the substance poured into them wil be between the Land H line. For example, you might have a measuring cup that has a Low line at 400ml and a high line at 435ml . This means that when you use this measuring cup, you can only be sure that what you're measuring is between 400ml and 435ml .
You're given a list of measuring cups containing their low and high lines as well as one low integer and one high integer representing a range for a target measurement. Write a function that returns a boolean representing whether you can use the cups to accurately measure a volume in the specified [low, high] range {the range is inclusive).
Note that:

• 	Each measuring cup will be represented by a pair of positive integers [L, HJ , where 0 <= L <= H .
• 	You'll always be given at least one measuring cup, and the low and high input parameters will always satisfy the following constraint: 0 <= low <= high .
• 	Once you've measured some liquid, it will immediately be transferred to a larger bowl that will eventually {possibly) contain the target measurement.
• 	You can't pour the contents of one measuring cup into another cup.


Sample Input:
cups = [
[200, 210],
[450, 465],
[800, 850]
]

low = 2100
high = 2300

Sample Output:
true
// We use cup [450, 465] to measure three volumes:
// First measurement: Low = 450, High = 465
// Second measurement: Low = 900, High = 930
// Third measurement: Low = 1350, High = 1395
// This covers the range [2100, 2300] , so we return true .

Explanation
The given solution works fine but I personally found It extremely confusing and hard to navigate, especially when one Is trying to master recursive algorithmic concepts.
In the solution below. I've organised the recursive calls in a way that presents crystal clear base cases and recursive case. In more detail, the cases are:

• 	Base case 1: We get a cache hit and know that the current combination of (low. high) cannot enter the desired range. Return False
• 	Base case 2: We have exceeded the target range. Return False.
• 	Base case 3: We have successfully entered the target range. Return True
• 	Recursive case: Check each available cup and make a recursive call to find out if the new low and new high will be able to enter the target range. If It can, we can bubble up a True return value to the top of the call stack, otherwise we append the failed combInat1on to the set and keep trying.

Another imponant note is that there Is no need to store a cache and boolean values, since as soon as we find a valid range, we stop the algorithm and return True to the top. We will only ever write to the cache 1f the current range failed to enter the target range, therefore we can JUSt use a set of failed combinations.
More comments In the code. Hope this helps'



Hints
Hint 1

Start by considering the last cup that you'll use in your sequence of measurements. If it isn't possible to use any of the cups as the last cup, then you can't measure the desired volume.

Hint 2

If the cup that you're going to use last has a measuring range of [100, 110] and you want to measure in the range of [500, 550], then after you pick this cup as the last cup, you need to measure a range of [400, 440]. Now, you can simply pick the last cup you'll use to measure this new range. If you continue these steps, you'll eventually know if you're able to measure the entire range or not.

Hint 3

Hint #2 should give you an idea of how to solve this problem recursively. Try every cup as the last cup for the starting range, then recursively try to measure the new ranges created after using the selected last cups. If you ever reach a point where one cup can measure the entire range, then you're finished and you can measure the target range. Try to think of a way to optimize this recursive approach, since it might involve a lot of repeated calculations.





This code first defines a function canMeasure that takes a list of measuring cups, a low target volume, and a high target volume and returns a boolean indicating whether the given measuring cups can accurately measure a volume in the specified range. The function first initializes an empty set to store failed combinations. Then, it calls the canMeasureRecursive function to recursively check if the target range can be reached using the given measuring cups.

The canMeasureRecursive function takes the same arguments as the canMeasure function, along with a set of previously failed combinations. The function first checks if the current combination has already been tried and failed. If so, it returns false. Otherwise, it checks if the current range exceeds the target range or falls completely outside the target range. If so, it returns false. Otherwise, if the current range is within the target range, it returns true.

If the current range is not within the target range, the function iterates through the given measuring cups and tries each cup. For each cup, it calculates the new low and high values that would be obtained by pouring the current contents into the cup. Then, it recursively calls itself to check if the new range can be reached using the remaining cups. If the new range can be reached, the function returns true. Otherwise, it adds the failed combination to the set of failed combinations and continues iterating.

Finally, if the function has tried all cups and none of them have led to a successful measurement, it returns false.

The main function calls the canMeasure function with the given input parameters and prints the result.



*/

package main

import (
	"fmt"
	"set"
)

func canMeasure(cups [][]int, low, high int) bool {
	failedCombinations := set.New()

	return canMeasureRecursive(cups, low, high, failedCombinations)
}

func canMeasureRecursive(cups [][]int, currentLow, currentHigh int, failedCombinations *set.Set) bool {
	if failedCombinations.Contains([]int{currentLow, currentHigh}) {
		return false
	}

	if currentLow > high || currentHigh < low {
		return false
	}

	if currentLow <= low && currentHigh >= high {
		return true
	}

	for _, cup := range cups {
		newLow := max(currentLow, cup[0])
		newHigh := min(currentHigh, cup[1])

		if canMeasureRecursive(cups, newLow, newHigh, failedCombinations) {
			return true
		}

		failedCombinations.Add([]int{newLow, newHigh})
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cups := [][]int{{200, 210}, {450, 465}, {800, 850}}
	low := 2100
	high := 2300

	canMeasureResult := canMeasure(cups, low, high)
	fmt.Println(canMeasureResult)
}
