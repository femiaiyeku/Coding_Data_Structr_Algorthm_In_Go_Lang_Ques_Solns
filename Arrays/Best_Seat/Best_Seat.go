/*
You walk into a theatre you're about to see a show in.
The usher within the theatre walks you to your row and mentions you're allowed to sit anywhere within the given row.
Naturally you'd like to sit in the seat that gives you the most space. You also would prefer this space to be evenly distributed on either side of you {e.g. if there are three empty seats in a row, you would prefer to sit in the middle of those three seats).
Given the theatre row represented as an integer array, return the seat index of where you should sit.
Ones represent occupied seats and zeroes represent empty seats.
You may assume that someone is always sitting in the first and last seat of the row.
Whenever there are two equally good seats, you should sit in the seat with the lower index.
If there is no seat to sit in, return -1. The given array will always have a length of at least one and contain only ones and zeroes.


Sample Input
seats = [1, 0, 0, 0, 1, 0, 1]

Sample Output
4



To use the function, simply pass in the theatre row as an argument. The function will return the seat index of where you should sit.

For example, using the sample input provided in the question, we would get the following output:

Go
seats := []int{1, 0, 0, 0, 1, 0, 1}

bestSeatIndex := findBestSeat(seats)

fmt.Println(bestSeatIndex) // 4

*/

package main

func findBestSeat(seats []int) int {

	// Initialize the best seat index.
	bestSeatIndex := -1

	// Iterate over the theatre row.
	for i := 1; i < len(seats)-1; i++ {

		// Check if the current seat is empty.
		if seats[i] == 0 {

			// Calculate the number of empty seats on either side of the current seat.
			leftEmptySeats := 0
			rightEmptySeats := 0
			for j := i - 1; j >= 0; j-- {
				if seats[j] == 0 {
					leftEmptySeats++
				} else {
					break
				}
			}
			for j := i + 1; j < len(seats); j++ {
				if seats[j] == 0 {
					rightEmptySeats++
				} else {
					break
				}
			}

			// If the current seat has the most space on either side, update the best seat index.
			if (leftEmptySeats >= bestSeatIndex || rightEmptySeats >= bestSeatIndex) && (leftEmptySeats == rightEmptySeats) {
				bestSeatIndex = i
			}
		}
	}

	// Return the best seat index.
	return bestSeatIndex
}
