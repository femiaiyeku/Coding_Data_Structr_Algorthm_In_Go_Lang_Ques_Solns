/*



You're given an array of integers asteroids, where each Integer represents the size of an asterold. The sign of
the integer represents the direction the asteroid is moving (positive-right, negative=left). All asteroids move at
the same speed, meaning that two asteroids moving in the same direction can never collide.
For example, the integer 4 represents an asterold of size 4 moving to the right. Similarly, -7 represents an
asteroid of size 7 moving to the left.
If two asteroids collide, the smaller asterold (based on absolute value) explodes. If two colliding asterolds are the
same size, they both explode.
Write a function that takes in this array of asterolds and returns an array of Integers representing the asterolds
after all collisions occur.


Sample Input:
asteroids = [-5, 10, -15]

Sample Output:
[-5, 10]

Explanation:
-5 and 10 never collide because they move in the same direction.
-15 collides with 10, and -15 explodes.
10 collides with -5, and -5 explodes.

Sample Input:
asteroids = [-5, -5, -5, -5]

Sample Output:
[]
Explanation:
All asteroids collide and explode.



This code uses a stack to keep track of the asteroids moving from right to left. It iterates through the asteroids from the back, checking if there are any asteroids on the stack moving in the same direction. If there are, it compares their sizes and removes the smaller or both asteroids if they are the same size. Finally, it converts the stack to a slice and returns the remaining asteroids.




*/

package main

import (
	"fmt"
)

func collide(asteroids []int) []int {
	// Initialize a stack to keep track of the asteroids
	stack := []int{}

	for i := len(asteroids) - 1; i >= 0; i-- {
		asteroid := asteroids[i]

		// Check if there are any asteroids on the stack moving in the same direction
		sameDirectionFound := false
		for len(stack) > 0 && asteroids[stack[len(stack)-1]]*asteroid > 0 {
			topAsteroid := asteroids[stack[len(stack)-1]]

			// If the top asteroid is smaller, it explodes
			if abs(topAsteroid) < abs(asteroid) {
				stack = stack[:len(stack)-1]
			} else if abs(topAsteroid) == abs(asteroid) { // If the asteroids are the same size, they both explode
				stack = stack[:len(stack)-1]
				asteroid = 0
			} else { // If the top asteroid is larger, it survives
				sameDirectionFound = true
				break
			}
		}

		// If there are no asteroids on the stack or we found an asteroid moving in the same direction, push the current asteroid onto the stack
		if len(stack) == 0 || !sameDirectionFound {
			stack = append(stack, asteroid)
		}
	}

	// Convert the stack to a slice
	remainingAsteroids := make([]int, len(stack))
	copy(remainingAsteroids, stack)

	return remainingAsteroids
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	asteroids := []int{-5, 10, -15, 10, -5}
	remainingAsteroids := collide(asteroids)
	fmt.Println("Remaining asteroids:", remainingAsteroids)
}
