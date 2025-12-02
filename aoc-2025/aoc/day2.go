package aoc

import (
	"strconv"
	"strings"

	"github.com/pepperonirollz/aoc-2025/utils"
)

func Day2Part1(path string) int64 {
	data, _ := utils.Parse(path)
	splitLines := strings.Split(data[0], ",")
	sum := int64(0)
	for _, v := range splitLines {
		startEnd := strings.Split(v, "-")
		start := startEnd[0]
		end := startEnd[1]
		startInt, _ := strconv.ParseInt(start, 10, 64)
		endInt, _ := strconv.ParseInt(end, 10, 64)

		for j := startInt; j <= endInt; j++ {
			if hasRepeatedString(strconv.FormatInt(int64(j), 10)) {
				sum += j
			}
		}
	}
	return int64(sum)
}

func Day2Part2(path string) int64 {
	data, _ := utils.Parse(path)
	splitLines := strings.Split(data[0], ",")
	sum := int64(0)
	for _, v := range splitLines {
		startEnd := strings.Split(v, "-")
		start := startEnd[0]
		end := startEnd[1]
		startInt, _ := strconv.ParseInt(start, 10, 64)
		endInt, _ := strconv.ParseInt(end, 10, 64)

		for j := startInt; j <= endInt; j++ {
			if isRepeatedPattern(strconv.FormatInt(int64(j), 10)) {
				sum += j
			}
		}
	}
	return int64(sum)
}

func hasRepeatedString(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]
	return firstHalf == secondHalf
}

func isRepeatedPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	for length := 1; length <= n/2; length++ {
		if n%length == 0 {
			pattern := s[:length]
			if strings.Repeat(pattern, n/length) == s {
				return true
			}
		}
	}
	return false
}
