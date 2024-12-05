package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	input := string(buf)
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	// build grid
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	count := 0
	for y, row := range grid {
		for x := range row {
			count = count_words(x, y, grid, "XMAS", count)
		}
	}

	println("Part1: ", count)

	count = 0
	for y, row := range grid {
		for x := range row {
			if is_x_mas(x, y, grid) {
				count++
			}
		}
	}

	println("Part2: ", count)
}

func count_words(x int, y int, grid [][]byte, search_word string, count int) int {
	// rows
	count = count_words_in_dir(x, y, -1, 0, grid, search_word, count)
	count = count_words_in_dir(x, y, 1, 0, grid, search_word, count)
	// columns
	count = count_words_in_dir(x, y, 0, -1, grid, search_word, count)
	count = count_words_in_dir(x, y, 0, 1, grid, search_word, count)
	// diags /
	count = count_words_in_dir(x, y, -1, -1, grid, search_word, count)
	count = count_words_in_dir(x, y, 1, 1, grid, search_word, count)
	// diags \
	count = count_words_in_dir(x, y, -1, 1, grid, search_word, count)
	count = count_words_in_dir(x, y, 1, -1, grid, search_word, count)
	return count
}

func count_words_in_dir(x int, y int, x_dir int, y_dir int, grid [][]byte, search_word string, count int) int {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		// out of bounds
		return count
	}

	if grid[y][x] == search_word[0] {
		if len(search_word) == 1 {
			// found a match!
			return count + 1
		}
		search_word = search_word[1:]
		return count_words_in_dir(x+x_dir, y+y_dir, x_dir, y_dir, grid, search_word, count)
	}
	// no match
	return count
}

func is_x_mas(x int, y int, grid [][]byte) bool {
	if grid[y][x] != 'A' || y < 1 || y >= (len(grid)-1) || x < 1 || x >= (len(grid[y])-1) {
		return false
	}

	r := regexp.MustCompile(`(MAS)|(SAM)`)

	s1 := []byte{grid[y-1][x-1], grid[y][x], grid[y+1][x+1]}
	s2 := []byte{grid[y+1][x-1], grid[y][x], grid[y-1][x+1]}

	return r.Match(s1) && r.Match(s2)
}
