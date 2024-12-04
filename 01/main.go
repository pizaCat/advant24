package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var list1 = []int{}
	var list2 = []int{}
	for _, line := range lines {
		vals := strings.Split(line, " ")
		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	sort.Ints(list1)
	sort.Ints(list2)
	var result int = 0
	for index, cur := range list1 {
		if cur < list2[index] {
			result = result + list2[index] - cur
		} else {
			result = result + cur - list2[index]
		}
	}
	fmt.Println("Part1: ", result)

	result = 0
	for _, cur1 := range list1 {
		for _, cur2 := range list2 {
			if cur1 == cur2 {
				result = result + cur1
			}
		}
	}
	fmt.Println("Part2: ", result)
}
