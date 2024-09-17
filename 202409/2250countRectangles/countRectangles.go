package main

import (
	"sort"
)

func main() {

}

// 使用二分
func countRectangles2(rectangles [][]int, points [][]int) []int {
	// Initialize the result array to store counts
	result := make([]int, len(points))

	// Map to store y-coordinates and their corresponding points with x-coordinates and indices
	yToPointsMap := make(map[int][]pair)
	for i, point := range points {
		yToPointsMap[point[1]] = append(yToPointsMap[point[1]], pair{point[0], i})
	}

	// Process each y-coordinate
	for height, pointsList := range yToPointsMap {
		lengthsList := getLengths(height, rectangles)
		for _, point := range pointsList {
			// Binary search to find the number of rectangles that cover the point
			idx := sort.Search(len(lengthsList), func(i int) bool { return lengthsList[i] >= point.X })
			result[point.Index] = len(lengthsList) - idx
		}
	}

	return result
}

// 不使用二分
func countRectangles(rectangles [][]int, points [][]int) []int {
	// Initialize the result array to store counts
	result := make([]int, len(points))

	// Map to store y-coordinates and their corresponding points with x-coordinates and indices
	yToPointsMap := make(map[int][]pair)
	for i, point := range points {
		yToPointsMap[point[1]] = append(yToPointsMap[point[1]], pair{point[0], i})
	}

	// Process each unique y-coordinate
	for height, pointsList := range yToPointsMap {
		// Get all rectangle lengths that can cover points at this height
		lengthsList := getLengths(height, rectangles)
		for _, point := range pointsList {
			// Count the number of rectangles whose length is >= point.x using linear search
			count := 0
			for _, length := range lengthsList {
				if length >= point.X {
					count++
				}
			}
			// Store the count of rectangles covering the point
			result[point.Index] = count
		}
	}
	return result
}

type pair struct {
	X     int
	Index int
}

func getLengths(minHeight int, rectangles [][]int) []int {
	var lengths []int
	for _, rect := range rectangles {
		if rect[1] >= minHeight {
			lengths = append(lengths, rect[0])
		}
	}
	sort.Ints(lengths)
	return lengths
}
