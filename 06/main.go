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
const obstacle byte = '#'

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	// build og_grid (y,x)
	og_grid := [][]byte{}
	for _, line := range lines {
		og_grid = append(og_grid, []byte(line))
	}

	grid := copy_grid(og_grid)
	move_guard(grid)

	obstacle_ideas := [][]int{}
	result := 0
	for r, row := range grid {
		for c, val := range row {
			if val == up || val == down || val == left || val == right {
				result++
				// grab these spots to use them in part 2!
				obstacle_ideas = append(obstacle_ideas, []int{r, c})
			}
		}
	}

	fmt.Println("Part1:", result)

	result = 0
	for _, rc := range obstacle_ideas {
		i := rc[0]
		j := rc[1]
		c := og_grid[i][j]
		if c == '.' {
			// make a grid copy
			grid2 := copy_grid(og_grid)
			// replace current pos with obstacle
			grid2[i][j] = obstacle

			if move_guard(grid2) {
				result++
			}

			print(".")
		}
	}
	println()
	fmt.Println("Part2:", result)
}

func copy_grid(grid [][]byte) [][]byte {
	g2 := make([][]byte, len(grid))
	for i, row := range grid {
		g2[i] = append(g2[i], row[:]...)
	}
	return g2
}

func move_guard(grid [][]byte) bool {
	guard_xy := []int{-1, -1}

	// find guard
	for i, row := range grid {
		if slices.Contains(row, '^') {
			guard_xy[1] = i
			guard_xy[0] = slices.Index(row, '^')
			break
		}
	}

	turn_dest := up
	hist := []string{}
	// while guard on map
	for within_bounds(guard_xy, grid) {
		dest_xy := []int{guard_xy[0], guard_xy[1]}
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
			break
		}

		if grid[dest_xy[1]][dest_xy[0]] == obstacle {
			// dest is an obstacle, turn clockwise
			grid[guard_xy[1]][guard_xy[0]] = byte(turn_dest)
		} else {
			curr_move := string(guard_xy[0]) + string(guard_xy[1]) + string(grid[guard_xy[1]][guard_xy[0]])
			// check if destination is already traveled in the same direction
			if slices.Contains(hist, curr_move) {
				// infinite loop detected!
				return true
			}
			// register move
			hist = append(hist, curr_move)
			// move to destination
			grid[dest_xy[1]][dest_xy[0]] = grid[guard_xy[1]][guard_xy[0]]
			guard_xy = dest_xy
		}
	}
	return false
}

func within_bounds(pos_xy []int, grid [][]byte) bool {
	return pos_xy[0] >= 0 && pos_xy[0] < len(grid[0]) && pos_xy[1] >= 0 && pos_xy[1] < len(grid)
}
