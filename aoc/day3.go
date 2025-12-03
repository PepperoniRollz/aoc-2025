package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pepperonirollz/aoc-2025/utils"
)

func Day3Part1(path string) int64 {
	data, _ := utils.Parse(path)
	sum := int64(0)
	for _, line := range data {
		max := -1

		for i := 0; i < len(line); i++ {
			for j := i + 1; j <= len(line)-1; j++ {
				substr := string(line[i]) + string(line[j])
				num, _ := strconv.ParseInt(substr, 10, 64)
				if num > int64(max) {
					max = int(num)
				}
			}
		}
		sum += int64(max)
	}
	return int64(sum)
}

func Day3Part2(path string) int64 {
	data, _ := utils.Parse(path)
	sum := int64(0)
	for _, line := range data {
		joltage := make([]int64, 12)
		index := 0
		max := int64(0)
		maxIndex := 0
		for index < 12 {
			for i := maxIndex; i <= len(line)-1; i++ {
				parsed, _ := strconv.ParseInt(string(line[i]), 10, 64)
				if parsed > max && len(line)-i >= 12-index {
					max = parsed
					maxIndex = i
				}
			}
			joltage[index] = max
			index++
			maxIndex++
			max = 0
		}
		var builder strings.Builder
		for _, v := range joltage {
			builder.WriteString(strconv.FormatInt(v, 10))
		}

		joltageStr := builder.String()

		joltageAsInt, _ := strconv.ParseInt(joltageStr, 10, 64)
		sum += joltageAsInt
	}
	return sum
}
