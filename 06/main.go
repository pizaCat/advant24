package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const OBSTACLE byte = '#'

type Orientation byte

const (
	UP    Orientation = '^'
	DOWN  Orientation = 'v'
	LEFT  Orientation = '<'
	RIGHT Orientation = '>'
)

type grid struct {
	width  int
	height int
	data   []byte
	guard  guard
}

func newGrid(input string) grid {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	width := len(lines[0])
	height := len(lines)
	data := make([]byte, width*height)

	for y, line := range lines {
		for x, val := range []byte(line) {
			data[x+(y*width)] = val
		}
	}

	pos := slices.Index(data, byte(UP))
	if pos == -1 {
		panic("Could not find the guard! Send help!")
	}
	x := pos % width
	y := pos / width
	o := UP
	guard := guard{false, []breadcrumb{{x, y, o}}}

	return grid{width, height, data, guard}
}

func (g grid) clone() grid {
	data := make([]byte, len(g.data))
	copy(data, g.data)
	return grid{g.width, g.height, data, g.guard.clone()}
}

func (g grid) withinBounds(pos breadcrumb) bool {
	return pos.x >= 0 && pos.x < g.width && pos.y >= 0 && pos.y < g.height
}

func (g grid) guardWithinBounds() bool {
	return g.withinBounds(g.guard.pos())
}

// func (g grid) print() {
// 	for i, val := range g.data {
// 		if i%g.width == 0 {
// 			fmt.Println()
// 		}
// 		fmt.Print(string(val))
// 	}
// 	fmt.Println()
// }

func (g grid) getVal(x int, y int) byte {
	if !g.withinBounds(breadcrumb{x, y, UP}) {
		panic("Out of bounds!")
	}
	return g.data[x+(y*g.width)]
}

func (g *grid) setVal(x int, y int, val byte) {
	i := x + (y * g.width)
	if i >= 0 && i < len(g.data) {
		g.data[i] = val
	}
}

func (g grid) uniquePathTiles() []breadcrumb {
	tiles := make(map[int]breadcrumb)
	for _, pos := range g.guard.path {
		if g.withinBounds(pos) {
			tiles[pos.x+(pos.y*g.width)] = pos
		}
	}
	values := make([]breadcrumb, len(tiles))

	i := 0
	for k := range tiles {
		values[i] = tiles[k]
		i++
	}
	return values
}

func (g *grid) move_guard() guard {
	for g.guardWithinBounds() && !g.guard.looping {
		move_dest := breadcrumb{g.guard.x(), g.guard.y(), g.guard.getOrientation()}
		switch move_dest.o {
		case UP:
			move_dest.y--
		case DOWN:
			move_dest.y++
		case LEFT:
			move_dest.x--
		case RIGHT:
			move_dest.x++
		default:
			panic("unrecognized guard!?!?")
		}

		if !g.withinBounds(move_dest) {
			return g.guard
		}

		if g.getVal(move_dest.x, move_dest.y) == OBSTACLE {
			// dest is an obstacle, turn clockwise
			g.guard.turnRight()
		} else {
			g.guard.move(move_dest)
		}
	}
	return g.guard
}

type breadcrumb struct {
	x int
	y int
	o Orientation
}

type guard struct {
	looping bool
	path    []breadcrumb
}

func (g guard) clone() guard {
	path := make([]breadcrumb, len(g.path))
	copy(path, g.path)
	return guard{g.looping, path}
}

func (g guard) x() int {
	return g.path[len(g.path)-1].x
}

func (g guard) y() int {
	return g.path[len(g.path)-1].y
}

func (g guard) getOrientation() Orientation {
	return g.path[len(g.path)-1].o
}

func (g guard) pos() breadcrumb {
	return g.path[len(g.path)-1]
}

func (g *guard) move(b breadcrumb) {
	if slices.Contains(g.path, b) {
		g.looping = true
	}
	g.path = append(g.path, b)
}

func (g *guard) turn(o Orientation) {
	g.move(breadcrumb{g.x(), g.y(), o})
}

func (g *guard) turnRight() {
	switch g.getOrientation() {
	case UP:
		g.turn(RIGHT)
	case RIGHT:
		g.turn(DOWN)
	case DOWN:
		g.turn(LEFT)
	case LEFT:
		g.turn(UP)
	default:
		panic("Imposter! GUAAAAARDS!")
	}
}

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	og_grid := newGrid(string(buf))

	grid_copy := og_grid.clone()
	grid_copy.move_guard()

	guard_path := grid_copy.uniquePathTiles()
	fmt.Println("Part1:", len(guard_path))

	result := 0
	for _, pos := range guard_path {
		if og_grid.getVal(pos.x, pos.y) == '.' {
			grid_copy = og_grid.clone()
			grid_copy.setVal(pos.x, pos.y, OBSTACLE)

			if grid_copy.move_guard().looping {
				result++
			}
		}
	}

	fmt.Println("Part2:", result)
}
