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
	result_p2 := 0
	for _, line := range update_lines {
		update := []int{}
		for _, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			update = append(update, n)
		}
		if in_order(update, rules) {
			middle_i := (len(update) - 1) / 2
			result = result + update[middle_i]
		} else {
			update = order_update(update, rules)
			middle_i := (len(update) - 1) / 2
			result_p2 = result_p2 + update[middle_i]
		}
	}

	fmt.Println("Part1: ", result)
	fmt.Println("Part1: ", result_p2)
}

func in_order(update []int, rules [][]int) bool {
	for i := 0; i < len(update)-1; i++ {
		n0 := update[i]
		n1 := update[i+1]

		for _, rule := range rules {
			if slices.Contains(rule, n0) && slices.Contains(rule, n1) && rule[0] == n1 {
				return false
			}
		}
	}
	return true
}

func order_update(update []int, rules [][]int) []int {
	for ordered := false; !ordered; ordered = in_order(update, rules) {
		for i := 0; i < len(update)-1; i++ {
			n0 := update[i]
			n1 := update[i+1]

			for _, rule := range rules {
				if slices.Contains(rule, n0) && slices.Contains(rule, n1) && rule[0] == n1 {
					// flip n0 and n1
					update = append(update[:i], update[i+1:]...)
					update = slices.Insert(update, i+1, n0)
				}
			}
		}
	}
	return update
}
