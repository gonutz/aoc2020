//+build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("trees for 3,1", treesForSlope(3, 1))
	fmt.Println(
		"product of trees for all slopes:", 1*
			treesForSlope(1, 1)*
			treesForSlope(3, 1)*
			treesForSlope(5, 1)*
			treesForSlope(7, 1)*
			treesForSlope(1, 2),
	)
}

func treesForSlope(dx, dy int) int {
	lines := strings.Split(input, "\n")
	x := 0
	y := 0
	trees := 0
	for y < len(lines) {
		line := lines[y]
		if line[x%len(line)] == '#' {
			trees++
		}
		x += dx
		y += dy
	}
	return trees
}

const input = `...........#..#.#.###....#.....
...#..#...........#.#...#......
#.....#..#........#...#..##....
..#...##.#.....#.............#.
#.#..#......#.....#....#.......
.....#......#..#....#.....#....
.......#.#..............#......
.....#...#..........##...#.....
...#....#.#...#.#........#...#.
..#.........###.......##...##..
.#....#...........#........#..#
..#.............##.............
..#.##.#....#................#.
.....##.#.......#....#...#.....
......#.#....##................
..#..........###..#..#.#..#....
....#..............#....##..#.#
.#.........#.#....#.#.#....#...
..#.....#......##.#....#.......
..#.#....#..#.#...##....###....
...#......##...#........#.#..#.
.##.#.......##....#............
...##..#.#............#...#.#..
.##...##.#..#..................
..#......##......#......##.....
.....##...#..#...#.........#...
.##.#.....#..#..#.##....##....#
..#.#......#.......##..........
......................#......##
##.#...#.................#.#.#.
......#.#..........#.....##.#..
#.#......#.....#...........#...
.....#...#.......#..#..#.#...#.
...........#......#.#...#......
....##...##...........#......#.
.........#.##..................
......#...#....#......##.##...#
......#...#.#########......#...
.......#.#...#.......#..#......
............#...#...#.###...##.
...........#..........#...#....
...#..#.#................#.#..#
..#....#.....#.............#.#.
....##......#........#....#....
........##...............#....#
........#..#...#..............#
...#....#.#...#..#...#....#.#.#
.........#.......#....##.......
#.##..............#.#........##
......................###......
.........#....##..##....#.#.#..
.#...##.#.#...#....##..#.....#.
....................#.#......#.
.#..#.......................#..
..#..#.............#..#....#...
...#...#..#...#...#.#..#.##....
........#....#...#....#........
.#.....#........#.........#...#
...#......#..#..#..####..#.....
#....#..............#.##.......
.#....#.............##...#.....
....#...#.##........##......#..
##....#...#.......#..#........#
....##........................#
..................#..#.........
..#....#........#..#.......#...
#...#..#....#...##...........#.
.........#..#..#.#.##..........
....#.#..#.....#..#.#.#.#..#.##
##................#.##.....#...
.#.....###..#..#..#.....#....##
...#...........#..........####.
.#.....#....#......#.##..#.#...
..#...##....#................#.
........#.......#......#.#.....
....#.#.#.#....#...#......#..#.
...........#......#..#.........
###...##......##.#..#....##....
##....##.........#..#....###...
#.#.....#....#......#...#..##..
#....##.#..............#.##....
.#........#.#.........#...#....
......................#......#.
........#..#..##.....#..#.#....
..#...###.................#..#.
...#...#............#..........
.##.......#..#.........#....#..
.#..............#....#....##...
...............##..#.#.......##
.#.....#....#...#..#.......#..#
#..#.............#....#......#.
.....#.#......#.........###..#.
.#...#.#...............#....#..
#......#.............#.........
.#.##.#.####...#..#.##..#.....#
.....#......#..#...#.......#...
#........###...#.....#..#.....#
....#.#.....#...#..........#...
...#.#.......#.#......#..#.##..
..#..........#.#..#.......#.#..
#...#.#..............#...###.#.
...#..#...#............###.....
..#..#...#.#............#..#...
.###.#.....................#..#
....#....#..#.....##.##........
#....#....#.#..#.........#.....
.#.....##....#............##..#
#....#.#...#...#..#.#......#...
#.....##.....##.....##.#...##..
...##...#..#..####.#........#..
.........#...#......##..##..#..
..#.....###.#..#..####.#.......
.......#......#...#..##....#...
.#.....#.#.#..#....#...#..##...
..........#.#...#...#.#..#.....
....#.....#........#.....##..#.
..#.#.##.........#...##.....##.
.........#.##....#............#
............##.....#.......#.#.
......#.....#...#..#..###......
##.....#.......#...##.#....#...
...........##......#..##...#.#.
..#.#.#.#...#.......#....#...#.
#.............#.....#.#...###..
##....#.......#.....#..##.#....
...#.......#....#.........##...
......#.......#......##.##.....
..#......#...#.#........#......
....#.#....#.#.##......#.....#.
#......#.........#..#....#.....
........#..#....##.....#.......
#......##....#.##....#......#..
..#.......#............##.....#
...........#...#...........#...
#.......#...#....#....#...#.#.#
..###..#........#........#.....
..#..###...........#.#......#..
.#...........#......#..........
.#.......#.....#.......#.......
..#......##.#............#....#
#..........#.....#.#..#........
.....#...##.##.......#......#..
..........#.....#........#.#.#.
....#......#.....#......#.#....
.........#.#.#..#...##....#...#
.........#.......#...##...#.#..
.##........#...............#...
.......#....#...........##.....
.........###......#.........#.#
......#.......#...#..........#.
...#.#..........##......#...#..
#.......#.#..........#.........
................#..#......#..##
.....#...#....#.#.....#........
#.....#....#...........#....#..
#....#.#..#...##....#...##.#...
...#.....#......#.#....#..#..#.
..#................#...#.#..##.
..........#..............#..#.#
.....##.....#..#.###...........
....#.#......#.#...........#...
.#....#.#.........##.#....#....
.#.#........#........###....#..
##.#................#...#..#...
.......#......##..#.....#..#.#.
...#............#......###...##
#.#...........#.........#......
.....#.#.#.................#...
....#..............#...#.#.....
...#.#.....##..#...............
.#..##...#....##.....###..#....
...............#...#...#.#.###.
.###....#.....#...#.#......#...
...#..#.....#.......#..##.#....
...........#..#....##..#...#...
...#...#..........#.......##.#.
............#.#.......#........
....#.........#.....#..........
...#.###.##..#...##..####..#..#
.#.#...#..#...................#
.....#..#.....##..#............
....#......#...##..#.##........
...#...............#..#.....##.
...#......#.........#.#..##....
.#....#.##.......#......#......
....#.......#....#..........#..
#.#.#....###.#.#.............#.
..##..###........#.#..#.#..#...
......#.#............##.#...###
.........#.#....#####.....##...
............##......#.#..#.....
...#.....#.....###....#........
##..........####.##.#.#........
....................##.....##.#
#.#............#........#......
....#...##.....#......#....#...
...###.#..##..................#
..###......#..............#.#.#
.#...#...........#....#........
....#............#..#..#.#.....
...#.........#.#.........#.####
..#...#...#...#...........#....
...............#.#......##..#..
#....#.#.......#.#..#......#..#
........#.#....#..#...#..#...#.
...#..#.......#...........#....
...........#.......#...........
.#......#................#.....
....#.#.#...#......#..#...#....
................#.#.#....#.....
.........................##..#.
.#...........#............##...
#...............#.....##.#.#..#
.........#.......###....#.....#
....##...#...#.....#..#........
........#.....#..#.#.#...#..#..
......#.......#.#.........#.#..
#......#............#...#....#.
#..##...#..#................#..
.##...#...#.....#.##.......#..#
.......#.##........##..##......
##.#..##...............#.....#.
......#....#..##...#......###.#
#........##..#....#.#......#...
.#......##.#...#.#...#.........
.#.#...#..#.............#......
.##..........#..........#......
.#.....#.....#..............#.#
..#.........#..#.#.....#.#....#
..#.##..............##...#..###
....................#..........
......###..#..#...........#....
..#..........#.......#...#.....
...#......#......#.............
....##..............#.#.....#..
........#.#......#..#........##
.............#...#.#.........##
...###...#..........##.......#.
.#..........#...##..#.#.....#..
##...#.........#...............
......#....#....#.....#.....#..
..........#....#...#...#..#...#
...##....#.#.#..#...##.........
#......#.#...##.###...#....#...
##.......##.#......##..#...#...
......#.............#.##.....##
#.......###....####.#...##....#
..#...#..#.......#..........#..
#.....#..#..#..#.##...###...#..
.....##.#..#..#..#.#....#...#..
..#...#..................##....
....#.#........##..............
#...#.......##...#...#.#.......
..#...#........##....#.#.......
..........###...###...#......#.
#.....#..###...##...##..#..#..#
..#.....##.....#.......##..#.#.
........#........#.........#...
.................#....#.......#
.......#...#.....#...#.#.......
....##...............#...##...#
.##...#................#...#...
.............#.................
.#..#....#....#.#....#.........
.#.#..#..........#.......#.....
.....##.....##...#..#..........
#...#.#.........#......#..#....
........#....#...#....#.#.##...
....#..#........#...#...#......
.#..#.....#.#...#.........#....
.#..#..#...........#..#....#...
....###.............#..#.......
#......#..#..##..........#.#...
#..#..#.##..#...#.#.#..........
....###......#.##.....#....#...
.##..#...#......##.#...........
..#..#.......#.....#.##....#.#.
.......#.#.#........#....##....
..##...#....#...............###
#..##..#...........#.#....##...
...##..#.....................#.
###......#....#....###..#...##.
.........##............#..#...#
..#..........#...#.#.#......#.#
.......#.....##..##......#.##..
#..........#.....##.#..........
#.......#.#...#...#....#.......
#...#.....##.......#.#..#.#.#..
.........#.#.#..#..#...#.###...
.................##...#....#...
###.......#..........##...#....
#.#..#.........#....##.#.......
......#.#.....#........#.......
.......#.#........#......#.#..#
..............#..#...##....#..#
#...........#...##.....#..#.#..
..#....#..#.#.#...#..#....#.#..
...##.#.....#..#...##..#.....#.
..#.#................#........#
......#...#.............#......
.##............#....#...#..#...
....#...#...........#.......#..
.###..#.......#.............#.#
.#.#....#.#...........#.#......
...#.........#.........#..#....
...#..........#.#.....#.#......
.....#........#....##......#...
..#.#.#......#..#.#......#....#
.#.#..#................#.#.....
.#.#.........##...#.......#.#.#
#..#.....#...#..#...........#..
..##......####......#..#....###
#.....###....#.#........#..#..#
..##.#...#.#..##..........#..#.
#.........#.#.............#...#
...#.#...#...#.#.#....##.......
##.##...#.....#...#...........#
....#........#.#.....#.........
.................##..#..##...##
.....##....#...#...#.....#..#..
....#...#........#............#
..#...........##....#...#...##.
.....#......#.........#..##.#..`
