//+build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		// f[z][y][x]
		type field [23][23][23]bool
		var f field
		for y, line := range strings.Split(input, "\n") {
			for x, col := range line {
				if col == '#' {
					f[11][7+y][7+x] = true
				}
			}
		}

		for i := 0; i < 6; i++ {
			var g field
			for z := range f {
				for y := range f[z] {
					for x := range f[z][y] {
						n := 0
						for zz := z - 1; zz <= z+1; zz++ {
							for yy := y - 1; yy <= y+1; yy++ {
								for xx := x - 1; xx <= x+1; xx++ {
									if zz >= 0 && zz < 23 &&
										yy >= 0 && yy < 23 &&
										xx >= 0 && xx < 23 &&
										!(zz == z && yy == y && xx == x) &&
										f[zz][yy][xx] {
										n++
									}
								}
							}
						}
						g[z][y][x] = (f[z][y][x] && (n == 2 || n == 3)) ||
							(!f[z][y][x] && n == 3)
					}
				}
			}
			f = g
		}

		active := 0
		for z := range f {
			for y := range f[z] {
				for x := range f[z][y] {
					if f[z][y][x] {
						active++
					}
				}
			}
		}
		fmt.Println("3D active after 6 cycles", active)
	}

	{
		// f[w][z][y][x]
		type field [23][23][23][23]bool
		var f field
		for y, line := range strings.Split(input, "\n") {
			for x, col := range line {
				if col == '#' {
					f[11][11][7+y][7+x] = true
				}
			}
		}

		for i := 0; i < 6; i++ {
			var g field
			for w := range f {
				for z := range f[w] {
					for y := range f[w][z] {
						for x := range f[w][z][y] {
							n := 0
							for ww := w - 1; ww <= w+1; ww++ {
								for zz := z - 1; zz <= z+1; zz++ {
									for yy := y - 1; yy <= y+1; yy++ {
										for xx := x - 1; xx <= x+1; xx++ {
											if ww >= 0 && ww < 23 &&
												zz >= 0 && zz < 23 &&
												yy >= 0 && yy < 23 &&
												xx >= 0 && xx < 23 &&
												!(ww == w && zz == z && yy == y && xx == x) &&
												f[ww][zz][yy][xx] {
												n++
											}
										}
									}
								}
							}
							g[w][z][y][x] = (f[w][z][y][x] && (n == 2 || n == 3)) ||
								(!f[w][z][y][x] && n == 3)
						}
					}
				}
			}
			f = g
		}

		active := 0
		for w := range f {
			for z := range f[w] {
				for y := range f[w][z] {
					for x := range f[w][z][y] {
						if f[w][z][y][x] {
							active++
						}
					}
				}
			}
		}
		fmt.Println("4D active after 6 cycles", active)
	}
}

const input = `.##.##..
..###.##
.##....#
###..##.
#.###.##
.#.#..#.
.......#
.#..#..#`
