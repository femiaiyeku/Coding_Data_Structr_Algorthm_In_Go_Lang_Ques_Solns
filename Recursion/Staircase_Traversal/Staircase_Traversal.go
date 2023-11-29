/*



You're given two positive integers representing the height of a staircase and the maximum number of steps that you can advance up the staircase at a time.
Write a function that returns the number of ways in which you can climb the staircase.
For example, if you were given a staircase of height = 3 and maxSteps = 2 you could climb the staircase in 3 ways. You could take 1 step, 1 step,
then 1 step, you could also take 1 step, then 2 steps, and you could take 2 steps, then 1 step.
Note that maxSteps <= height will always be true.

Sample Input:
height = 4
maxSteps = 2

Sample Output:

5

// There are 5 ways to climb the staircase:
// 1. 1 step + 1 step + 1 step + 1 step
// 2. 1 step + 1 step + 2 steps
// 3. 1 step + 2 steps + 1 step
// 4. 2 steps + 1 step + 1 step
// 5. 2 steps + 2 steps




Hints
Hint 1

If you can advance 2 steps at a time, how many ways can you reach a staircase of height 1 and of height 2? Think recursively.

Hint 2

Continuing from Hint #1, if you know the number of ways to climb a staircase of height 1 and of height 2, how many ways are there to climb a staircase of height 3 (assuming the same max steps of 2)?

Hint 3

The number of ways to climb a staircase of height k with a max number of steps s is: numWays[k - 1] + numWays[k - 2] + ... + numWays[k - s]. This is because if you can advance between 1 and s steps, then from each step k - 1, k - 2, ..., k - s, you can directly advance to the top of a staircase of height k. By adding the number of ways to reach all steps that you can directly advance to the top step from, you determine how many ways there are to reach the top step.






This code first defines a function climbStairs that takes the height of the staircase and the maximum number of steps that can be taken at a time as input and returns the number of ways to climb the staircase. The function initializes an array ways to store the number of ways to climb staircases of different heights. It then sets the number of ways to climb a staircase of height 0 to 1 (there is only one way: take no steps).

For each height from 1 to height, the function iterates over all possible step sizes from 1 to maxSteps or height (whichever is smaller). For each step size, it adds the number of ways to climb the staircase of height i-j to the number of ways to climb the staircase of height i. This is because there are ways[i-j] ways to climb the staircase of height i-j and then take a step of size j to reach the staircase of height i.

Finally, the function returns the number of ways to climb the staircase of height height, which is stored in the last element of the ways array.

The main function creates a staircase with the given input data and calls the climbStairs function to count the number of ways to climb it. It then prints the result.


*/

package main

import (
	"fmt"
)

func climbStairs(height int, maxSteps int) int {
	ways := make([]int, height+1)
	ways[0] = 1 // One way to climb a staircase of height 0

	for i := 1; i <= height; i++ {
		for j := 1; j <= maxSteps && j <= i; j++ {
			ways[i] += ways[i-j]
		}
	}

	return ways[height]
}

func main() {
	height := 4
	maxSteps := 2

	ways := climbStairs(height, maxSteps)
	fmt.Println("Number of ways to climb the staircase:", ways)
}
