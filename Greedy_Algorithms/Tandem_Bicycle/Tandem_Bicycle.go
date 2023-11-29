/*



A tandem bicycle 1s a bicycle that's operated by two people: person A and person B. Both people pedal the bicycle,
 but the person that pedals faster dictates the speed of the bicycle. So if person A pedals at a speed of 5 , and person B pedals at a speed of 4 ,
   the tandem bicycle moves at a speed of 5 (i.e., tandemSpeed = max (speedA, speedB) ).
You're given two lists of positive integers: one that contains the speeds of riders wearing red shirts and one that contains the speeds of riders wearing blue shirts.
 Each rider is represented by a single positive integer, which is the speed that they pedal a tandem bicycle at.
 Both lists have the same length, meaning that there are as many red-shirt riders as there are blue-shirt riders. Your goal is to pair every rider wearing a red shirt
 with a rider wearing a blue shirt to operate a tandem bicycle.
Write a function that returns the maximum possible total speed or the minimum possible total speed of all of the tandem bicycles being ridden based on an input parameter,
 fastest . If fastest = true , your function should return the maximum possible total speed; otherwise it should return the minimum total speed.
"Total speed" 1s defined as the sum of the speeds of all the tandem bicycles being ridden. For example, if there are 4 riders (2 red­shirt riders and 2 blue-shirt riders)
 who have speeds of 1, 3, 4, 5 , and if they're paired on tandem bicycles as follows:


[1, 4], [5, 3] , then the total speed of these tandem bicycles is 4 + 5 = 9.

Sample Input
redShirtSpeeds = [5, 5, 3, 9, 2]

blueShirtSpeeds = [3, 6, 7, 2, 1]

fastest = true

Sample Output
32

// 1f fastest = true , we should pair the riders wearing red shirts with the riders wearing blue shirts in the following way:
// [1, 9], [2, 7], [3, 6], [5, 3], [5, 2]
// The rider wearing red shirt 1 rides the tandem bicycle with the rider wearing blue shirt 9.
// The rider wearing red shirt 2 rides the tandem bicycle with the rider wearing blue shirt 7.
// The rider wearing red shirt 3 rides the tandem bicycle with the rider wearing blue shirt 6.
// The rider wearing red shirt 5 rides the tandem bicycle with the rider wearing blue shirt 3.
// The rider wearing red shirt 5 rides the tandem bicycle with the rider wearing blue shirt 2.
// Note that there are other ways to pair the riders wearing red shirts with the riders wearing blue shirts that also yield the maximum possible total speed.








*/

package main

import (
	"fmt"
	"sort"
)

func maximumTotalSpeed(redShirtSpeeds, blueShirtSpeeds []int, fastest bool) int {
	// Sort both arrays in either ascending or descending order based on the fastest parameter
	if fastest {
		sort.Slice(redShirtSpeeds, func(i, j int) bool {
			return redShirtSpeeds[i] > redShirtSpeeds[j]
		})
		sort.Slice(blueShirtSpeeds, func(i, j int) bool {
			return blueShirtSpeeds[i] > blueShirtSpeeds[j]
		})
	} else {
		sort.Ints(redShirtSpeeds)
		sort.Ints(blueShirtSpeeds)
	}

	// Calculate the total speed
	totalSpeed := 0
	for i := 0; i < len(redShirtSpeeds); i++ {
		if fastest {
			totalSpeed += redShirtSpeeds[i]
		} else {
			totalSpeed += blueShirtSpeeds[i]
		}
	}

	return totalSpeed
}

func main() {
	redShirtSpeeds := []int{5, 5, 3, 9, 2}
	blueShirtSpeeds := []int{3, 6, 7, 2, 1}
	fastest := true

	totalSpeed := maximumTotalSpeed(redShirtSpeeds, blueShirtSpeeds, fastest)
	fmt.Println(totalSpeed)
}
