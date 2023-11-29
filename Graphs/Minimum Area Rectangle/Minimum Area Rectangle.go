/*


You're given an array of points plotted on a 2D graph (the xy-plane). Write a function that returns the minimum area of any rectangle that can be formed using any 4 of these points such that the rectangle's sides are parallel to the x and y axes (i.e., only rectangles with horizontal and vertical sides should be considered--no rectangles with diagonal sides). If no rectangle can be formed, your function should return 0.

The input array will contain points represented by arrays of two integers [x, y]. The input array will never contain duplicate points.

Sample Input
points =
[
    [1, 5],
    [5, 1],
    [4, 2],
    [2, 4],
    [2, 2],
    [1, 2],
    [4, 5],
    [2, 5],
    [-1, -2],
]
Sample Output
3
// The rectangle with corners [1, 5], [2, 5], [1, 2], and [2, 2]
// has the minimum area: 3.
Hints
Hint 1

The brute-force approach to this problem is to simply generate all possible combinations of 4 points and to see if they form a rectangle. You can calculate the area of all of these rectangles and then return the minimum area that you find. Is there a better approach than this?

Hint 2

A more optimal approach is to find vertical or horizontal edges that are parallel to the y or x axes, respectively. If you find two parallel edges (two vertical edges, for example) that share a vertical or horizontal coordinate (y values in the case of vertical edges), then those edges form a rectangle.

Hint 3

Another approach is to pick any two points that don't have the same x or y values (i.e., points that could be at opposite ends of a rectangle diagonal) and to see if you can create a rectangle with them and two other points. Given two points where p1 = (x1, y1) and p2 = (x2, y2), if points p3 = (x1, y2) and p4 = (x2, y1) exist, then these 4 points form a rectangle.






*/

package rectangle_mania

// important question

import "fmt"

type Direction int

const (
	None Direction = iota - 1
	Up
	Down
	Left
	Right
)

// RectangleMania
// O(n^2) time | O(n) space - where n is the number of coordinates
func RectangleMania(coords [][]int) int {
	coordsTable := getCoordsTable(coords)
	return getRectangleCount(coords, coordsTable)
}

type CoordsTable map[string]struct{}

func getCoordsTable(coords [][]int) CoordsTable {
	table := CoordsTable{}
	for _, coord := range coords {
		table[coordToString(coord)] = struct{}{}
	}
	return table
}

func getRectangleCount(coords [][]int, coordsTable CoordsTable) int {
	count := 0
	for _, coord1 := range coords {
		for _, coord2 := range coords {
			if !isInUpperRight(coord1, coord2) {
				continue
			}
			upperCoord := []int{coord1[0], coord2[1]}
			rightCoord := []int{coord2[0], coord1[1]}
			_, found1 := coordsTable[coordToString(upperCoord)]
			_, found2 := coordsTable[coordToString(rightCoord)]
			if found1 && found2 {
				count++
			}
		}
	}
	return count
}

func isInUpperRight(coord1, coord2 []int) bool {
	return coord2[0] > coord1[0] && coord2[1] > coord1[1]
}

func coordToString(coord []int) string {
	return fmt.Sprintf("%d-%d", coord[0], coord[1])
}

// RectangleMania2
// O(n^2) time | O(n) space - where n is the number of coordinates
func RectangleMania2(coords [][]int) int {
	coordsTable := getCoordsTable2(coords)
	return getRectangleCount2(coords, coordsTable)
}

type CoordsTable2 struct {
	Xs, Ys map[int][][]int
}

func getCoordsTable2(coords [][]int) CoordsTable2 {
	table := CoordsTable2{
		Xs: map[int][][]int{},
		Ys: map[int][][]int{},
	}
	for _, coord := range coords {
		x, y := coord[0], coord[1]
		table.Xs[x] = append(table.Xs[x], coord)
		table.Ys[y] = append(table.Ys[y], coord)
	}
	return table
}

func getRectangleCount2(coords [][]int, coordsTable CoordsTable2) int {
	count := 0
	for _, coord := range coords {
		lowerLeftY := coord[1]
		count += clockwiseCountRectangles2(coord, coordsTable, Up, lowerLeftY)
	}
	return count
}

func clockwiseCountRectangles2(coord []int, coordsTable CoordsTable2, direction Direction, lowerLeftY int) int {
	if direction == Down {
		relevantCoords := coordsTable.Xs[coord[0]]
		for _, coord2 := range relevantCoords {
			lowerRightY := coord2[1]
			if lowerRightY == lowerLeftY {
				return 1
			}
		}
		return 0
	}

	if direction == Up {
		rectangleCount := 0
		relevantCoords := coordsTable.Xs[coord[0]]
		for _, coord2 := range relevantCoords {
			if coord2[1] > coord[1] {
				rectangleCount += clockwiseCountRectangles2(coord2, coordsTable, Right, lowerLeftY)
			}
		}
		return rectangleCount
	}

	if direction == Right {
		rectangleCount := 0
		relevantCoords := coordsTable.Ys[coord[1]]
		for _, coord2 := range relevantCoords {
			if coord2[0] > coord[0] {
				rectangleCount += clockwiseCountRectangles2(coord2, coordsTable, Down, lowerLeftY)
			}
		}
		return rectangleCount
	}
	return 0
}

// RectangleMania1
// O(n^2) time | O(n^2) space - where n is the number of coordinates
func RectangleMania1(coords [][]int) int {
	coordsTable := getCoordsTable1(coords)
	return getRectangleCount1(coords, coordsTable)
}

type CoordsTable1 map[string]map[Direction][][]int

func getCoordsTable1(coords [][]int) CoordsTable1 {
	table := CoordsTable1{}
	for _, coord1 := range coords {
		directions := map[Direction][][]int{
			Up:    {},
			Right: {},
			Down:  {},
			Left:  {},
		}
		for _, coord2 := range coords {
			coord2Direction := getCoordDirection1(coord1, coord2)
			if coord2Direction != None {
				directions[coord2Direction] = append(directions[coord2Direction], coord2)
			}
		}
		table[coordToString(coord1)] = directions
	}
	return table
}

func getCoordDirection1(coord1, coord2 []int) Direction {
	if coord2[1] == coord1[1] {
		if coord2[0] > coord1[0] {
			return Right
		} else if coord2[0] < coord1[0] {
			return Left
		}
	} else if coord2[0] == coord1[0] {
		if coord2[1] > coord1[1] {
			return Up
		} else if coord2[1] < coord1[1] {
			return Down
		}
	}
	return None
}

func getRectangleCount1(coords [][]int, coordsTable CoordsTable1) int {
	count := 0
	for _, coord := range coords {
		count += clockwiseCountRectangles1(coord, coordsTable, Up, coord)
	}
	return count
}

func clockwiseCountRectangles1(coord []int, coordsTable CoordsTable1, direction Direction, origin []int) int {
	if direction == Left {
		for _, element := range coordsTable[coordToString(coord)][Left] {
			if element[0] == origin[0] && element[1] == origin[1] {
				return 1
			}
		}
		return 0
	}
	rectangleCount := 0
	nextDirection := direction.NextClockwise()
	for _, nextCoord := range coordsTable[coordToString(coord)][direction] {
		rectangleCount += clockwiseCountRectangles1(nextCoord, coordsTable, nextDirection, origin)
	}
	return rectangleCount
}

func (d Direction) NextClockwise() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	return None
}

// {0, 0}, {0, 1},
// {1, 1}, {1, 0},
// {2, 1}, {2, 0},
// {3, 1}, {3, 0}

//0 [0, 1]
//1 [0, 1]
//2 [0, 1]
//3 [0, 1]

// 0 [0, 1, 2, 3]
// 1 [0, 1, 2, 3]

// 0,2 1,2
// 0,1 1,1 2,1 3,1
// 0,0 1,0 2,0 3,0

// my solution
type yCoordsSet map[string]struct{}

func rectangleMania(coords [][]int) int {
	coordsSet := createCoordsSet(coords)
	rectangles := make(map[string]yCoordsSet)

	for currentIdx := range coords {
		p2x, p2y := coords[currentIdx][0], coords[currentIdx][1]
		for previousIdx := 0; previousIdx < currentIdx; previousIdx++ {
			p1x, p1y := coords[previousIdx][0], coords[previousIdx][1]
			coordinatesShareValue := p1x == p2x || p1y == p2y

			if coordinatesShareValue {
				continue
			}

			coordinate1OnOppositeDiagonalExists := coordsSet[convertCoordinateToString(p1x, p2y)]
			coordinate2OnOppositeDiagonalExists := coordsSet[convertCoordinateToString(p2x, p1y)]
			oppositeDiagonalExists := coordinate1OnOppositeDiagonalExists && coordinate2OnOppositeDiagonalExists

			if oppositeDiagonalExists {
				//fmt.Println(p1x, p1y, p2x, p2y, p1x, p2y, p2x, p1y)

				//fmt.Sprintf("%d:%d", x, y)
				//p1x p2y   p2x p2y
				//p1x p1y   p2x p1y
				y1, y2 := p1y, p2y
				if y1 > y2 {
					y1, y2 = y2, y1
				}

				x1, x2 := p1x, p2x
				if x1 > x2 {
					x1, x2 = x2, x1
				}

				if _, ok := rectangles[convertCoordinateToString(x1, x2)]; !ok {
					yCoords := make(yCoordsSet)
					yCoords[convertCoordinateToString(y1, y2)] = struct{}{}
					rectangles[convertCoordinateToString(x1, x2)] = yCoords
				} else {
					rectangles[convertCoordinateToString(x1, x2)][convertCoordinateToString(y1, y2)] = struct{}{}
				}
			}
		}
	}

	var rectangleCount int
	for _, yCoords := range rectangles {
		rectangleCount += len(yCoords)
	}

	return rectangleCount
}

func createCoordsSet(coords [][]int) map[string]bool {
	coordsSet := make(map[string]bool)
	for _, coordinate := range coords {
		x, y := coordinate[0], coordinate[1]
		pointString := convertCoordinateToString(x, y)
		coordsSet[pointString] = true
	}
	return coordsSet
}

func convertCoordinateToString(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
