/*


Write a function that takes in a "special" array and returns its product sum.
A "special" array is a non-empty array that contains either integers or other "special" arrays. The product sum of a
"special" array is the sum of its elements, where "special" arrays inside it are summed themselves and then
multiplied by their level of depth.
The depth of a "special" array is how far nested it is. For instance, the depth of [[] is 1; the depth of the inner
array in [[]] is 2; the depth of the innermost array in [[[[J]] is 3.
Therefore, the product sum of [x, y] is x + y; the product sum of [x, [y, z]] is x + 2 (y + z);
the product sum of [x, [y, [z]]] is x + 2 * (y + 3z)

Sample Input:
array = [5, 2, [7, -1], 3, [6, [-13, 8], 4]]

Sample Output:
12
// calculated as: 5 + 2 + 2 * (7 - 1) + 3 + 2 * (6 + 3 * (-13 + 8) + 4)






This code first defines a type SpecialArray to represent a "special" array. Then, it defines a function productSum that takes a "special" array and a multiplier (representing the depth of the array) and returns the product sum of the array. The function iterates over the elements of the array and checks if each element is another "special" array or an integer. If it is another "special" array, it recursively calls itself to calculate the product sum of the nested array and multiplies the result by the multiplier. If it is an integer, it simply adds the integer to the sum. Finally, it returns the sum.

The main function creates a "special" array with the given input data and calls the productSum function to calculate its product sum. It then prints the result.


*/

package main

import (
	"fmt"
)

type SpecialArray []interface{}

func productSum(array SpecialArray, multiplier int) int {
	sum := 0

	for _, element := range array {
		if cast, ok := element.(SpecialArray); ok {
			sum += productSum(cast, multiplier+1)
		} else if cast, ok := element.(int); ok {
			sum += multiplier * cast
		}
	}

	return sum
}

func main() {
	array := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}}
	productSumResult := productSum(array, 1)
	fmt.Println("Product sum of", array, ":", productSumResult)
}
