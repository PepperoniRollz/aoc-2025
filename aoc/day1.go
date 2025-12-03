package aoc

import (
	"strconv"

	"github.com/pepperonirollz/aoc-2025/utils"
)

func Part1(path string) int64 {
	data, _ := utils.Parse(path)
	answer := 0
	loc := 50
	for _, line := range data {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])
		switch direction {
		case 'R':
			for range distance {
				loc = (loc + 1) % 100
			}
			if loc == 0 {
				answer += 1
			}

		case 'L':
			for range distance {
				loc = (loc - 1) % 100
			}
			if loc == 0 {
				answer += 1
			}

		}
	}
	return int64(answer)
}

func Part2(path string) int64 {
	data, _ := utils.Parse(path)
	answer := 0
	loc := 50
	for _, line := range data {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])
		switch direction {
		case 'R':
			for range distance {
				loc = (loc + 1) % 100
				if loc == 0 {
					answer += 1
				}
			}
		case 'L':
			for range distance {
				loc = (loc - 1) % 100
				if loc == 0 {
					answer += 1
				}
			}

		}
	}
	return int64(answer)
}
