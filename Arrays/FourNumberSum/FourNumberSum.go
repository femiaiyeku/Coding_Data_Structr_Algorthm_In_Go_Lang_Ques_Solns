/*

Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum. The function should find all quadruplets in the array that sum up to the target sum and return a two-dimensional array of all these quadruplets in no
particular order. If no four numbers sum up to the target sum, the function should return an empty array.
Sample Input
array = [7, 6, 4, -1, 1, 2]"

targetSum = 16
Sample Output
((7, 6, 4, -1), (7, 6, 1, 2))


This code first checks if the input array is empty or has less than four elements. If so, it returns an empty array.

Then, it iterates over the array and selects four distinct elements at a time. It calculates the sum of the four selected elements and compares it to the target sum. If the sum is equal to the target sum, it adds the four elements to the result array.

Finally, it returns the result array.


*/

package main

import "fmt"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	result := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if i == j || i == k || j == k {
					continue
				}

				sum := nums[i] + nums[j] + nums[k]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[i+j+k-sum]})
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{7, 6, 4, -1, 1, 2}
	target := 16

	quadruplets := fourSum(nums, target)

	fmt.Println(quadruplets)
}
