/*


Write a function that takes in an array of unique integers and returns its powerset.
The powerset P (X) of a set X is the set of all subsets of X . For example, the powerset of [ 1, 2] is
[[], [1], [2], [1,2]]
Note that the sets in the powerset do not need to be in any particular order.

Sample Input: [1, 2, 3]



Sample Output: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]





Hints
Hint 1
Try thinking about the base cases. What is the powerset of the empty set? What is the powerset of sets of length 1?
Hint 2
If you were to take the input set X and add an element to it, how would the resulting powerset change?
Hint 3
Can you solve this problem recursively? Can you solve it iteratively? What are the advantages and disadvantages of using either approach?





This code first defines a function powerset that takes an array of integers and returns its powerset as a slice of slices of integers. The function first initializes an empty slice of slices of integers to store the powerset. Then, it iterates over all possible subsets of the input array, using bitmasks to efficiently represent the subsets. For each subset, it creates a slice of integers to represent the subset and adds it to the powerset. Finally, it returns the powerset.

The main function calls the powerset function with the given input array and prints the result.


*/

package main

import (
	"fmt"
)

func powerset(nums []int) [][]int {
	powerset := [][]int{}

	for i := 0; i < 1<<len(nums); i++ {
		subset := []int{}

		for j := 0; j < len(nums); j++ {
			if (i>>j)&1 == 1 {
				subset = append(subset, nums[j])
			}
		}

		powerset = append(powerset, subset)
	}

	return powerset
}

func main() {
	nums := []int{1, 2, 3}

	powerset := powerset(nums)
	fmt.Println("Powerset of", nums, ":", powerset)
}
