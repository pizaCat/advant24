package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var result int = 0

	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var reports = [][]int{}
	for _, line := range lines {
		reports = append(reports, stoi(line))
	}

	for _, report := range reports {
		if safe(report) {
			result = result + 1
		}
	}
	fmt.Println("Part1: ", result)

	result = 0
	for _, report := range reports {
		if safe(report) {
			result = result + 1
		} else {
			for i := 0; i < len(report); i++ {
				var copy []int
				for j := 0; j < len(report); j++ {
					if i != j {
						copy = append(copy, report[j])
					}
				}
				if safe(copy) {
					result = result + 1
					break
				}
			}
		}
	}
	fmt.Println("Part1: ", result)
}

func stoi(line string) []int {
	r := regexp.MustCompile(`[0-9]+`)
	matches := r.FindAllString(line, -1)
	var ints = []int{}
	for _, i := range matches {
		j, _ := strconv.Atoi(i)
		ints = append(ints, j)
	}
	return ints
}

func safe(report []int) bool {
	if report[0] < report[1] {
		// inc
		for index, val := range report {
			if index > 0 && (val > (report[index-1]+3) || val <= report[index-1]) {
				return false
			}
		}
		return true
	} else if report[0] > report[1] {
		// dec
		for index, val := range report {
			if index > 0 && (val < (report[index-1]-3) || val >= report[index-1]) {
				return false
			}
		}
		return true
	}
	return false
}
