package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const up byte = '^'
const down byte = 'v'
const left byte = '<'
const right byte = '>'

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	// build grid (y,x)
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	guard_xy := []int{-1, -1}

	// find guard
	for i, row := range grid {
		if slices.Contains(row, '^') {
			guard_xy[1] = i
			guard_xy[0] = slices.Index(row, '^')
			break
		}
	}

	// while guard on map
	for within_bounds(guard_xy, grid) {
		// for _, row := range grid {
		// 	for _, c := range row {
		// 		fmt.Print(string(c))
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()

		dest_xy := []int{guard_xy[0], guard_xy[1]}
		turn_dest := up
		guard_char := grid[guard_xy[1]][guard_xy[0]]

		if guard_char == up {
			// move up
			dest_xy[1] = dest_xy[1] - 1
			turn_dest = right
		} else if guard_char == down {
			// move down
			dest_xy[1] = dest_xy[1] + 1
			turn_dest = left
		} else if guard_char == left {
			// move left
			dest_xy[0] = dest_xy[0] - 1
			turn_dest = up
		} else if guard_char == right {
			// move right
			dest_xy[0] = dest_xy[0] + 1
			turn_dest = down
		} else {
			panic("unrecognized guard!?!?")
		}

		if !within_bounds(dest_xy, grid) {
			// set current guard pos to visited
			grid[guard_xy[1]][guard_xy[0]] = 'X'
			break
		}

		if grid[dest_xy[1]][dest_xy[0]] == '#' {
			// dest is an obstacle, turn clockwise
			grid[guard_xy[1]][guard_xy[0]] = byte(turn_dest)
		} else {
			// move to destination
			grid[dest_xy[1]][dest_xy[0]] = grid[guard_xy[1]][guard_xy[0]]
			grid[guard_xy[1]][guard_xy[0]] = 'X'
			guard_xy = dest_xy
		}
	}

	result := 0
	for _, row := range grid {
		for _, val := range row {
			if val == 'X' {
				result++
			}
		}
	}

	fmt.Println("Part1:", result)
}

func within_bounds(pos_xy []int, grid [][]byte) bool {
	return pos_xy[0] >= 0 && pos_xy[0] < len(grid[0]) && pos_xy[1] >= 0 && pos_xy[1] < len(grid)
}
