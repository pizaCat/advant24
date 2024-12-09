package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	sections := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	rule_lines := strings.Split(sections[0], "\n")
	update_lines := strings.Split(sections[1], "\n")

	rules := [][]int{}
	for _, line := range rule_lines {
		nums := strings.Split(line, "|")
		n0, _ := strconv.Atoi(nums[0])
		n1, _ := strconv.Atoi(nums[1])
		rules = append(rules, []int{n0, n1})
	}

	result := 0
	for _, line := range update_lines {
		update := []int{}
		for _, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			update = append(update, n)
		}
		if in_order(update, rules) {
			middle_i := (len(update) - 1) / 2
			result = result + update[middle_i]
		}
	}

	fmt.Println("Part1: ", result)
}

func in_order(seq []int, rules [][]int) bool {
	for i := 0; i < len(seq)-1; i++ {
		n0 := seq[i]
		n1 := seq[i+1]

		for _, rule := range rules {
			if slices.Contains(rule, n0) && slices.Contains(rule, n1) && rule[0] == n1 {
				return false
			}
		}
	}
	return true
}
