package aoc

import (
	"fmt"
	"time"

	"github.com/pepperonirollz/aoc-2025/utils"
)

func Day4Part1(path string) int64 {
	data, _ := utils.Parse(path)
	sum := int64(0)
	for i, line := range data {
		for j := 0; j < len(line); j++ {
			current := line[j]
			neighborCount := 0
			neighbors := getNeighbors(data, i, j)
			if current == '@' {
				for _, neighbor := range neighbors {
					if neighbor == '@' {
						neighborCount++
					}
				}
				if neighborCount < 4 {
					sum++
				}
			}
		}
	}

	return sum
}

func Day4Part2(path string) int64 {
	data, _ := utils.Parse(path)
	grid := make([][]byte, len(data))
	for i, line := range data {
		grid[i] = []byte(line)
	}

	numRemoved := 0
	iteration := 0

	fmt.Print("\033[2J\033[H") // Clear screen and move cursor to top
	fmt.Println("Initial grid:")
	printGrid(grid)
	time.Sleep(100 * time.Millisecond)

	for {
		toBeRemoved := make([][2]int, 0)

		for i := range grid {
			for j := 0; j < len(grid[i]); j++ {
				current := grid[i][j]
				neighborCount := 0
				neighbors := getNeighborsGrid(grid, i, j)
				if current == '@' {
					for _, neighbor := range neighbors {
						if neighbor == '@' {
							neighborCount++
						}
					}
					if neighborCount < 4 {
						toBeRemoved = append(toBeRemoved, [2]int{i, j})
					}
				}
			}
		}

		if len(toBeRemoved) == 0 {
			break
		}

		for _, coord := range toBeRemoved {
			row := coord[0]
			col := coord[1]
			grid[row][col] = '.'
			numRemoved++
		}

		iteration++
		fmt.Print("\033[2J\033[H") // Clear screen and move cursor to top
		fmt.Printf("Iteration %d (removed %d cells this round):\n\n", iteration, len(toBeRemoved))
		printGrid(grid)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\n\nFinal - Total removed: %d\n", numRemoved)
	return int64(numRemoved)
}

func getNeighbors(data []string, i, j int) []byte {
	var neighbors []byte
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := i + dir[0]
		newCol := j + dir[1]

		if newRow >= 0 && newRow < len(data) && newCol >= 0 && newCol < len(data[newRow]) {
			neighbors = append(neighbors, data[newRow][newCol])
		}
	}

	return neighbors
}

func getNeighborsGrid(grid [][]byte, i, j int) []byte {
	var neighbors []byte
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := i + dir[0]
		newCol := j + dir[1]

		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[newRow]) {
			neighbors = append(neighbors, grid[newRow][newCol])
		}
	}

	return neighbors
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
