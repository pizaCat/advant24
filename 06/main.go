package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type tileType byte

const (
	UP          tileType = '^'
	DOWN        tileType = 'v'
	LEFT        tileType = '<'
	RIGHT       tileType = '>'
	EMPTY_SPACE tileType = '.'
	OBSTACLE    tileType = '#'
)

type grid struct {
	width    int
	height   int
	guardPos position
	tiles    map[position]tileType
}

func newGrid(input string) (grid, guard) {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	width := len(lines[0])
	height := len(lines)
	tiles := make(map[position]tileType)

	guard_p := position{-1, -1}
	guard_o := UP

	for y, line := range lines {
		for x, val := range []tileType(line) {
			switch val {
			case UP, DOWN, LEFT, RIGHT:
				guard_p.x = x
				guard_p.y = y
				guard_o = val
			}
			tiles[position{x, y}] = val
		}
	}

	if guard_p.x == -1 {
		panic("Could not find the guard! Send help!")
	}

	grid := grid{width, height, guard_p, tiles}
	return grid, guard{false, []tile{{guard_p, guard_o}}, &grid}
}

func (g grid) withinBounds(pos position) bool {
	return pos.x >= 0 && pos.x < g.width && pos.y >= 0 && pos.y < g.height
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

func (g grid) getVal(p position) tileType {
	return g.tiles[p]
}

func (g guard) getUniqueTiles() []tile {
	tiles := make(map[position]tile)
	for _, tile := range g.path {
		if g.grid.withinBounds(tile.p) {
			tiles[tile.p] = tile
		}
	}
	values := make([]tile, len(tiles))

	i := 0
	for k := range tiles {
		values[i] = tiles[k]
		i++
	}
	return values
}

type position struct {
	x int
	y int
}
type tile struct {
	p position
	t tileType
}

type guard struct {
	looping bool
	path    []tile
	grid    *grid
}

func (g guard) o() tileType {
	return g.path[len(g.path)-1].t
}

func (g guard) p() position {
	return g.path[len(g.path)-1].p
}

func (g *guard) moveTo(p position) {
	tile := tile{p, g.o()}
	if slices.Contains(g.path, tile) {
		g.looping = true
	}
	g.path = append(g.path, tile)
}

func (g *guard) moveAndTurnTo(b tile) {
	if slices.Contains(g.path, b) {
		g.looping = true
	}
	g.path = append(g.path, b)
}

func (g *guard) turn(o tileType) {
	g.moveAndTurnTo(tile{g.p(), o})
}

func (g *guard) turnRight() {
	switch g.o() {
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

func (g guard) withinBounds() bool {
	return g.grid.withinBounds(g.p())
}

func (guard guard) move(extraObstacle position) guard {
	for guard.withinBounds() && !guard.looping {
		move_dest := guard.p()
		switch guard.o() {
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

		if !guard.grid.withinBounds(move_dest) {
			return guard
		}

		if guard.grid.getVal(move_dest) == OBSTACLE ||
			(move_dest.x == extraObstacle.x && move_dest.y == extraObstacle.y) {
			// dest is an obstacle, turn clockwise
			guard.turnRight()
		} else {
			guard.moveTo(move_dest)
		}
	}
	return guard
}

func main() {
	buf, _ := os.ReadFile("RAW_INPUT.txt")
	grid, guard := newGrid(string(buf))

	guard_path := guard.move(position{-1, -1}).getUniqueTiles()
	fmt.Println("Part1:", len(guard_path))

	potential_obstacles := []position{}
	for _, breadcrumb := range guard_path {
		if grid.getVal(breadcrumb.p) == EMPTY_SPACE {
			potential_obstacles = append(potential_obstacles, breadcrumb.p)
		}
	}
	result := 0
	c := make(chan bool, len(potential_obstacles))
	for _, pos := range potential_obstacles {
		go doesGuardLoop(guard, pos, c)
	}

	for range potential_obstacles {
		if <-c {
			result++
		}
	}
	fmt.Println("Part2:", result)
}

func doesGuardLoop(g guard, pos position, c chan bool) {
	c <- g.move(pos).looping
}
