package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type equation struct {
	result  uint64
	numbers []uint64
	reverse []uint64
}

func newEquation(line string) equation {
	items := strings.Split(line, ": ")
	result, _ := strconv.ParseUint(items[0], 10, 64)
	items = strings.Split(items[1], " ")
	numbers := make([]uint64, len(items))
	reverse := make([]uint64, len(items))
	for i, n := range items {
		num, _ := strconv.ParseUint(n, 10, 64)
		numbers[i] = num
		reverse[i] = num
	}

	slices.Reverse(reverse)

	return equation{result, numbers, reverse}
}

func (e equation) solve() ([]byte, bool) {
	return e.innerSolve([]byte{})
}
func (e equation) innerSolve(ops []byte) ([]byte, bool) {
	if len(e.reverse) == 1 {
		return ops, (e.result) == e.reverse[0]
	}

	// more numbers left
	sum := uint64(0)
	for _, n := range e.reverse[1:] {
		if n != 1 {
			// 1's don't count because they are cancelled when multiplying!!!
			sum = sum + n
		}
	}

	res := e.result / e.reverse[0]
	mod := e.result % e.reverse[0]
	if res >= sum && mod == 0 {
		ops2 := make([]byte, len(ops))
		copy(ops2, ops)
		ops2 = append(ops2, '*')
		ops2, solved := equation{res, e.numbers, e.reverse[1:]}.innerSolve(ops2)
		if solved {
			return ops2, solved
		}
	}

	res = e.result - e.reverse[0]
	if res >= sum {
		ops = append(ops, '+')
		ops, solved := equation{res, e.numbers, e.reverse[1:]}.innerSolve(ops)
		if solved {
			return ops, solved
		}
	}
	return ops, false
}

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	lines := strings.Split(strings.ReplaceAll(string(buf), "\r\n", "\n"), "\n")

	result := uint64(0)
	for _, line := range lines {
		eq := newEquation(line)
		_, solved := eq.solve()
		if solved {
			result = result + eq.result
		}
	}
	fmt.Println("Part1:", result)
}
