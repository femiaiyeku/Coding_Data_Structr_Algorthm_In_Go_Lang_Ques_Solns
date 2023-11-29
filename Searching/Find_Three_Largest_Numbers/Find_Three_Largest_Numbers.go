/*


Write a function that takes in an array of at least three integers and,
without sorting the input array, returns a sorted array of the three largest integers 1n the input array.
The function should return duplicate integers 1f necessary; for example, 1t should return [10, 18, 12] for an input array of

[10, 5, 9, 10,12]

Sample Input:
array = [141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7]

Sample Output:
[18, 141, 541]


Hints
Hint 1
Can you keep track of the three largest numbers in an array as you traverse the input array?
Hint 2
Following the suggestion in Hint #1, try traversing the input array and updating the three largest numbers if necessary by shifting them accordingly.









This code first defines a function findThreeLargest that takes an array of integers as input and returns a slice of the three largest integers in the array. The function first checks if the input array has at least three elements. If it does not, it returns an empty slice.

Otherwise, it initializes three variables to store the indices of the three largest integers in the array. It then iterates over the remaining elements of the array and checks if each element is larger than the current largest, second-largest, or third-largest integer. If it is, it updates the corresponding variable accordingly.

Finally, it returns a slice containing the three largest integers in the array.

The main function creates an array of integers and calls the findThreeLargest function to find the three largest integers in the array. It then prints the result.










*/

package main

import (
	"fmt"
)

func findThreeLargest(A []int) []int {
	if len(A) < 3 {
		return nil
	}

	largest, secondLargest, thirdLargest := A[0], A[1], A[2]

	for i := 3; i < len(A); i++ {
		if A[i] > largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = A[i]
		} else if A[i] > secondLargest {
			thirdLargest = secondLargest
			secondLargest = A[i]
		} else if A[i] > thirdLargest {
			thirdLargest = A[i]
		}
	}

	return []int{largest, secondLargest, thirdLargest}
}

func main() {
	array := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	threeLargest := findThreeLargest(array)
	fmt.Println("The three largest elements are:", threeLargest)
}
