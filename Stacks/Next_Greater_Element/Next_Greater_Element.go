/*

You're given an array of Integers and another array of three distinct Integers. The first array is guaranteed
only contain Integers that are in the second array, and the second array represents a desired order for the
integers in the first array. For example, a second array of [x, y, z1 represents a desired order of
[X, X, X, Y YY, 2, 2, 2] In the first array.
Write a function that sorts the first array according to the desired order in the second array.
The function should perform this in place (i.e., It should mutate the input array), and it shouldn't use any
auxiliary space (1.e., it should run with constant space: 0(1) space).
Note that the desired order won't necessarily be ascending or descending and that the first array won't
necessarily contain all three integers found in the second array-It might only contain one or two.

Sample Input
array = [1, 0, 0, -1, -1, 0, 1, 1]
order = [0, 1, -1]

Sample Output
[0, 0, 0, 1, 1, 1, -1, -1]





Hints
Hint 1

Solving this problem in O(n^2) time, where n is the length of the array, is straightforward. Can you solve it with a better time complexity?

Hint 2

How can a stack allow you to solve this problem in O(n) time?

Hint 3

There are a couple of ways to solve this problem in linear time with a stack. One approach is to push onto the stack the indices of elements for which you haven't yet found the next greater element. If you go with this index approach, you need to loop through the array twice (since it's circular) and compare the value of the current element in the array to the one represented by the index on top of the stack. If the element on the top of the stack is smaller than the current element, then the current element is next greater element for the top-of-stack element, and you can pop the index off the top of the stack and use it to store the current element in the correct position in your result array. You then continue to pop elements off the top of the stack until the current element is no longer greater than the top-of-stack element. At this point, you add the index of the current element to the top of the stack, and you continue iterating through the array, repeating the same process.

Hint 4

The approach discussed in Hint #3 assumes that you loop through the array from left to right. You could loop through the array backwards using a very similar approach, storing the actual values of elements on the stack rather than their indices. See the Conceptual Overview section of this question's video explanation for a more in-depth explanation.

Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the array






*/


package main

import (
    "fmt"
)

func sortWithOrder(array []int, order []int) {
    // Initialize variables for tracking indices and element counts
    orderIndex := 0
    elementCount := 0

    // Iterate through the array
    for i := range array {
        // Check if the current element matches the desired order
        if array[i] == order[orderIndex] {
            // Count the occurrences of the current element
            elementCount++

            // Move to the next element in the desired order if the count is exhausted
            if elementCount == counts[order[orderIndex]] {
                orderIndex++
                elementCount = 0
            }
        }
    }

    // Iterate through the array again and place the elements in the desired order
    for i := range array {
        // Check if the current element matches the desired order
        if array[i] == order[orderIndex] {
            // Swap the current element with the corresponding element in the desired order
            array[i], array[orderIndex] = array[orderIndex], array[i]

            // Move to the next element in the desired order if the count is exhausted
            elementCount++
            if elementCount == counts[order[orderIndex]] {
                orderIndex++
                elementCount = 0
