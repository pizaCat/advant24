package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type equation struct {
	result          uint64
	reverse_numbers []uint64
}

func equationFromString(line string) equation {
	items := strings.Split(line, ": ")
	result, _ := strconv.ParseUint(items[0], 10, 64)
	items = strings.Split(items[1], " ")
	numbers := make([]uint64, len(items))
	for i, n := range items {
		num, _ := strconv.ParseUint(n, 10, 64)
		numbers[i] = num
	}

	slices.Reverse(numbers)
	return equation{result, numbers}
}

func (e equation) isSolvable(use_concat bool) bool {
	if len(e.reverse_numbers) == 1 {
		return (e.result) == e.reverse_numbers[0]
	}

	return e.solveMult(use_concat) ||
		e.solveAdd(use_concat) ||
		(use_concat && e.solveConcat(use_concat))
}

func (e equation) solveMult(use_concat bool) bool {
	return e.result%e.reverse_numbers[0] == 0 &&
		equation{e.result / e.reverse_numbers[0], e.reverse_numbers[1:]}.isSolvable(use_concat)
}

func (e equation) solveAdd(use_concat bool) bool {
	return equation{e.result - e.reverse_numbers[0], e.reverse_numbers[1:]}.isSolvable(use_concat)
}

func (e equation) solveConcat(use_concat bool) bool {
	s_res := strconv.FormatUint(e.result, 10)
	s_num := strconv.FormatUint(e.reverse_numbers[0], 10)
	if strings.HasSuffix(s_res, s_num) {
		s_res = s_res[:len(s_res)-len(s_num)]
		res, _ := strconv.ParseUint(s_res, 10, 64)
		return equation{res, e.reverse_numbers[1:]}.isSolvable(use_concat)
	}
	return false
}

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	lines := strings.Split(strings.ReplaceAll(string(buf), "\r\n", "\n"), "\n")

	result_p1 := uint64(0)
	result_p2 := uint64(0)
	for _, line := range lines {
		eq := equationFromString(line)
		if eq.isSolvable(false) {
			result_p1 = result_p1 + eq.result
		}
		if eq.isSolvable(true) {
			result_p2 = result_p2 + eq.result
		}
	}
	fmt.Println("Part1:", result_p1)
	fmt.Println("Part2:", result_p2)
}
