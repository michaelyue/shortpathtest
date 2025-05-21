package main

import (
	"fmt"
)

// save BFS element data
type Position struct {
	x int
	y int
}

// get shortest path and fill in the grid
// for example
//	0 -1 -1 -1
// -1 -1 -1 -1
// -1 -1  0 -1
// -1 -1 -1 -1
// here 0 stands for uber position
// -1 means not filled with shortest path yet

// Time Complexity is O(n) , n equals gird size row*col
// Space Complexity is O(n), should save same grid and return
func shortestPath(grid [][]int) [][]int {
	// in case empty, just return empty
	if len(grid) == 0 || len(grid[0]) == 0 {
		return [][]int{}
	}

	// define uber move directions
	directions := []Position{
		{0, 1},  // move up
		{0, -1}, // move down
		{1, 0},  // move right
		{-1, 0}, // move left
	}

	rows := len(grid)
	cols := len(grid[0])

	// initialize the result grid
	// and get uber's position as well
	uberQueue := make([]Position, 0)
	result := make([][]int, rows)
	for i, v := range grid {
		result[i] = make([]int, cols)
		for j, col := range v {
			result[i][j] = 0
			if col == 0 {
				uberQueue = append(uberQueue, Position{i, j})
			}
		}
	}

	// BFS(breadth first search) algo to fill in all the shortest path for nearest Uber
	for len(uberQueue) > 0 {
		currUber := uberQueue[0]
		uberQueue = uberQueue[1:]

		for _, direction := range directions {
			x := currUber.x + direction.x
			y := currUber.y + direction.y
			if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == -1 {
				result[x][y] = result[currUber.x][currUber.y] + 1
				grid[x][y] = 3 // set as filled status
				uberQueue = append(uberQueue, Position{x, y})
			}
		}
	}
	return result
}

func main() {
	fmt.Println("hello world")

	// a few test cases
	fmt.Println(shortestPath([][]int{
		{0, -1},
		{-1, -1},
	}))

	fmt.Println(shortestPath([][]int{
		{0},
	}))
	fmt.Println(shortestPath([][]int{
		{-1},
	}))
	fmt.Println(shortestPath([][]int{
		{-1, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, -1, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, -1, -1, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, -1, -1, -1},
		{0, -1, -1, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, -1, -1, -1},
		{0, -1, 0, -1},
	}))
	fmt.Println(shortestPath([][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}))
	fmt.Println(shortestPath([][]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}))

}
