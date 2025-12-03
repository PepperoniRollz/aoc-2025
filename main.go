package main

import (
	"fmt"
	"github.com/pepperonirollz/aoc-2025/aoc"
)

func main() {
	day1Part1Result := aoc.Part1("inputs/day1.txt")
	day1Part2Result := aoc.Part2("inputs/day1.txt")

	fmt.Printf("Day 1 - Part 1: %d\n", day1Part1Result)
	fmt.Printf("Day 1 - Part 2: %d\n", day1Part2Result)

	day2Part1Result := aoc.Day2Part1("inputs/day2.txt")
	day2Part2Result := aoc.Day2Part2("inputs/day2.txt")

	fmt.Printf("Day 2 - Part 1: %d\n", day2Part1Result)
	fmt.Printf("Day 2 - Part 2: %d\n", day2Part2Result)

	day3Part1Result := aoc.Day3Part1("inputs/day3.txt")
	day3Part2Result := aoc.Day3Part2("inputs/day3.txt")

	fmt.Printf("Day 3 - Part 1: %d\n", day3Part1Result)
	fmt.Printf("Day 3 - Part 2: %d\n", day3Part2Result)

}
