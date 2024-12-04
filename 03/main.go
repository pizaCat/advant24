package main

import (
	"os"
	"regexp"
	"strconv"
)

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	text := string(buf)

	println("part1: ", process_do(text))

	r := regexp.MustCompile(`(?s)(^|do\(\))(.*?)(don't\(\)|$)`)
	matches := r.FindAllString(text, -1)
	res := 0
	for _, match := range matches {
		//println(match)
		res = res + process_do(match)
	}
	println("part2: ", res)
}

func process_do(input string) int {
	r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	matches := r.FindAllString(input, -1)

	res := 0
	for _, match := range matches {
		r := regexp.MustCompile(`[0-9]+`)
		nums := r.FindAllString(match, -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		//println(match, "->", a, "x", b, "=", a*b)
		res = res + (a * b)
	}
	//println()
	return res
}
