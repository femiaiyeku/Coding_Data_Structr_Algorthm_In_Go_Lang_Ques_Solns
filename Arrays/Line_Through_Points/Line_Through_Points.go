/*

You·re given an array of points plotted on a 2D graph (the xy-plane).
Write a function that returns the maximum number of points that a single line (or potentially multiple lines) on the graph passes through.
The ,nput array will contain points represented by an array of two integers [x, y] .
 The input array will never contain duplicate points and will always contain at least one point.

Sample Input
points = [
[-1, 0],
[0, 0],
[1, 0],
[2, 0],
[3, 0],
[3, 1],
]

Sample Output
4 // A line passes through points [-1, 0], [0, 0], [1, 0], [2, 0]


Hints

Hint 1

The brute-force approach to solve this problem is to consider every single pair of points and to form a line using them. Then, for each line, you determine how many points lie on that line by using the equation of the line you formed and checking if each point's coordinates solve the equation. This solution runs in O(n^3) time; can you come up with a better approach?

Hint 2

What does it mean if two lines have the same slope and share a point?

Hint 3

If two lines have the same slope and share a point, they're the same line. Try using a hash table to store the slopes of lines that pass through certain points. How does this help you write an algorithm that runs in O(n^2) time?

Hint 4

Loop through every single pair of points, picking a p2 for every p1 in order to form a line. For every pair (p1, p2), store the slope of the formed line in a hash table, and map it to the number of points on that line. If you ever find two identical slopes for lines that both use the same point p1, you can consider these lines to be one and the same, meaning that points p1, p2a, and p2b are all on the same line; in those cases, update the number of points on the slope (the line) in the hash table accordingly. You'll need to reset the hash table at each change of p1. See the Conceptual Overview section of this question's video explanation for a more in-depth explanation.



This function works by first sorting the points by x-coordinate. This is necessary to ensure that the function can efficiently find the minimum area rectangle. The function then creates a map to store the y-coordinates of the points at each x-coordinate. This map is used to quickly find the y-coordinates of the points that can be used to form rectangles with the current x-coordinate.

The function then iterates over the x-coordinates of the points. For each x-coordinate, the function sorts the y-coordinates at that x-coordinate. The function then iterates over the y-coordinates at the current x-coordinate and calculates the area of a rectangle with the current x-coordinate, the current y-coordinate, the next y-coordinate, and the next x-coordinate. The function updates the minimum area if necessary.

If no rectangle can be formed, the function returns 0. Otherwise, the function returns the minimum area.

*/

package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
}

type Line struct {
	slope      float64
	yIntercept float64
}

func findMaxPointsOnLine(points []Point) int {
	// Sort the points by x-coordinate.
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	// Create a map to store the lines that pass through each point.
	lineMap := make(map[Point][]Line)

	// Iterate over the points and calculate the slope and y-intercept of the line that passes through each pair of points.
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			slope := float64(points[j].y-points[i].y) / float64(points[j].x-points[i].x)
			yIntercept := float64(points[i].y) - slope*float64(points[i].x)

			line := Line{slope, yIntercept}

			// Add the line to the maps for both points.
			lineMap[points[i]] = append(lineMap[points[i]], line)
			lineMap[points[j]] = append(lineMap[points[j]], line)
		}
	}

	// Initialize the maximum number of points on a line.
	maxPoints := 0

	// Iterate over the points and count the number of lines that pass through each point.
	for point, lines := range lineMap {
		numLines := 0
		for _, line := range lines {
			numLines++
		}

		// Update the maximum number of points on a line if necessary.
		if numLines > maxPoints {
			maxPoints = numLines
		}
	}

	return maxPoints
}

func main() {
	points := []Point{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{3, 1},
	}

	maxPoints := findMaxPointsOnLine(points)

	fmt.Println(maxPoints)
}
