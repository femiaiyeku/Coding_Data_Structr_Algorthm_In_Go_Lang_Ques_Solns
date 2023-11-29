/*

You are given an array of points plotted on a 2D graph(the xy-plane).
Write a function that returns the minimum area of any rectangle that can be formed using any 4 of these points such that the rectangle's sides are parallel to the x and y axes.
If no rectangle can be formed, your function should return 0.

The input array will contain points represented by arrays of two integers [x, y].The input array will never contain duplicate points.

Sample Input:
points = [[1, 5], [-5, -1], [-1, -5], [5, 1]]

Sample Output:
12

Explanation:
The rectangle with corners [-5, -1], [-1, -5], [1, 5], and [5, 1],
whose sides are parallel to the x and y axes, can be formed from these points,
 so your function should return 12.


 Hint 1

The brute-force approach to this problem is to simply generate all possible combinations of 4 points and to see if they form a rectangle. You can calculate the area of all of these rectangles and then return the minimum area that you find. Is there a better approach than this?

Hint 2

A more optimal approach is to find vertical or horizontal edges that are parallel to the y or x axes, respectively. If you find two parallel edges (two vertical edges, for example) that share a vertical or horizontal coordinate (y values in the case of vertical edges), then those edges form a rectangle.

Hint 3

Another approach is to pick any two points that don't have the same x or y values (i.e., points that could be at opposite ends of a rectangle diagonal) and to see if you can create a rectangle with them and two other points. Given two points where p1 = (x1, y1) and p2 = (x2, y2), if points p3 = (x1, y2) and p4 = (x2, y1) exist, then these 4 points form a rectangle.


*/

func findMinAreaRectangle(points [][]int) int {
	// Initialize the minimum area.
	minArea := 0

	// Iterate over all pairs of points.
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			// Calculate the width and height of the rectangle formed by the two points.
			width := abs(points[i][0] - points[j][0])
			height := abs(points[i][1] - points[j][1])

			// Iterate over the remaining points to find two other points that can form a rectangle with the current two points.
			for k := j + 1; k < len(points); k++ {
				for l := k + 1; l < len(points); l++ {
					// Check if the four points can form a rectangle.
					if isRectangle(points[i], points[j], points[k], points[l]) {
						// Calculate the area of the rectangle.
						area := width * height

						// If the area is less than the current minimum area, update the minimum area.
						if area < minArea || minArea == 0 {
							minArea = area
						}
					}
				}
			}
		}
	}

	return minArea
}

// Returns true if the four points can form a rectangle.
func isRectangle(point1, point2, point3, point4 []int) bool {
	// Check if the opposite sides are equal in length.
	if abs(point1[0]-point2[0]) != abs(point3[0]-point4[0]) ||
		abs(point1[1]-point2[1]) != abs(point3[1]-point4[1]) {
		return false
	}

	// Check if the diagonals are equal in length.
	if (point1[0]-point3[0])*(point1[0]-point3[0])+(point1[1]-point3[1])*(point1[1]-point3[1]) !=
		(point2[0]-point4[0])*(point2[0]-point4[0])+(point2[1]-point4[1])*(point2[1]-point4[1]) {
		return false
	}

	return true
}

// Returns the absolute value of a number.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
  