package Challenges

import ()

/*
--- Day 17: Conway Cubes ---
As your flight slowly drifts through the sky, the Elves at the Mythical Information Bureau at the North Pole contact you.
They'd like some help debugging a malfunctioning experimental energy source aboard one of their super-secret imaging satellites.

The experimental energy source is based on cutting-edge technology: a set of Conway Cubes contained in a pocket dimension!
When you hear it's having problems, you can't help but agree to take a look.

The pocket dimension contains an infinite 3-dimensional grid. At every integer 3-dimensional coordinate (x,y,z), there exists
a single cube which is either active or inactive.

In the initial state of the pocket dimension, almost all cubes start inactive. The only exception to this is a small flat
region of cubes (your puzzle input); the cubes in this region start in the specified active (#) or inactive (.) state.

The energy source then proceeds to boot up by executing six cycles.

Each cube only ever considers its neighbors: any of the 26 other cubes where any of their coordinates differ by at most 1.
For example, given the cube at x=1,y=2,z=3, its neighbors include the cube at x=2,y=2,z=2, the cube at x=0,y=2,z=3, and so on.

During a cycle, all cubes simultaneously change their state according to the following rules:

If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes
inactive.
If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
The engineers responsible for this experimental energy source would like you to simulate the pocket dimension and determine what
the configuration of cubes should be at the end of the six-cycle boot process.

For example, consider the following initial state:

.#.
..#
###
Even though the pocket dimension is 3-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular,
	this initial state defines a 3x3x1 region of the 3-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown
layer-by-layer at each given z coordinate (and the frame of view follows the active cells in each cycle):

Before any cycles:

z=0
.#.
..#
###


After 1 cycle:

z=-1
#..
..#
.#.

z=0
#.#
.##
.#.

z=1
#..
..#
.#.


After 2 cycles:

z=-2
.....
.....
..#..
.....
.....

z=-1
..#..
.#..#
....#
.#...
.....

z=0
##...
##...
#....
....#
.###.

z=1
..#..
.#..#
....#
.#...
.....

z=2
.....
.....
..#..
.....
.....


After 3 cycles:

z=-2
.......
.......
..##...
..###..
.......
.......
.......

z=-1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=0
...#...
.......
#......
.......
.....##
.##.#..
...#...

z=1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=2
.......
.......
..##...
..###..
.......
.......
.......
After the full six-cycle boot process completes, 112 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles. How many cubes are left in the active state after the sixth cycle?

--- Part Two ---
For some reason, your simulated results don't match what the experimental energy source engineers expected. Apparently, the pocket
dimension actually has four spatial dimensions, not three.

The pocket dimension contains an infinite 4-dimensional grid. At every integer 4-dimensional coordinate (x,y,z,w), there exists a single
cube (really, a hypercube) which is still either active or inactive.

Each cube only ever considers its neighbors: any of the 80 other cubes where any of their coordinates differ by at most 1. For example,
given the cube at x=1,y=2,z=3,w=4, its neighbors include the cube at x=2,y=2,z=3,w=3, the cube at x=0,y=2,z=3,w=4, and so on.

The initial state of the pocket dimension still consists of a small flat region of cubes. Furthermore, the same rules for cycle updating
still apply: during each cycle, consider the number of active neighbors of each cube.

For example, consider the same initial state as in the example above. Even though the pocket dimension is 4-dimensional, this initial state
represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1x1 region of the 4-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer
at each given z and w coordinate:

Before any cycles:

z=0, w=0
.#.
..#
###


After 1 cycle:

z=-1, w=-1
#..
..#
.#.

z=0, w=-1
#..
..#
.#.

z=1, w=-1
#..
..#
.#.

z=-1, w=0
#..
..#
.#.

z=0, w=0
#.#
.##
.#.

z=1, w=0
#..
..#
.#.

z=-1, w=1
#..
..#
.#.

z=0, w=1
#..
..#
.#.

z=1, w=1
#..
..#
.#.


After 2 cycles:

z=-2, w=-2
.....
.....
..#..
.....
.....

z=-1, w=-2
.....
.....
.....
.....
.....

z=0, w=-2
###..
##.##
#...#
.#..#
.###.

z=1, w=-2
.....
.....
.....
.....
.....

z=2, w=-2
.....
.....
..#..
.....
.....

z=-2, w=-1
.....
.....
.....
.....
.....

z=-1, w=-1
.....
.....
.....
.....
.....

z=0, w=-1
.....
.....
.....
.....
.....

z=1, w=-1
.....
.....
.....
.....
.....

z=2, w=-1
.....
.....
.....
.....
.....

z=-2, w=0
###..
##.##
#...#
.#..#
.###.

z=-1, w=0
.....
.....
.....
.....
.....

z=0, w=0
.....
.....
.....
.....
.....

z=1, w=0
.....
.....
.....
.....
.....

z=2, w=0
###..
##.##
#...#
.#..#
.###.

z=-2, w=1
.....
.....
.....
.....
.....

z=-1, w=1
.....
.....
.....
.....
.....

z=0, w=1
.....
.....
.....
.....
.....

z=1, w=1
.....
.....
.....
.....
.....

z=2, w=1
.....
.....
.....
.....
.....

z=-2, w=2
.....
.....
..#..
.....
.....

z=-1, w=2
.....
.....
.....
.....
.....

z=0, w=2
###..
##.##
#...#
.#..#
.###.

z=1, w=2
.....
.....
.....
.....
.....

z=2, w=2
.....
.....
..#..
.....
.....
After the full six-cycle boot process completes, 848 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles in a 4-dimensional space. How many cubes
are left in the active state after the sixth cycle?
*/

func Challenge17(useTestData bool) {
	var data []string
	if useTestData {
		data = challenge17TestData
	} else {
		data = challenge17InputData
	}

	println("Part 1")
	challenge17Part1(data)
	println("Part 2")
	challenge17Part2(data)
}

func challenge17Part1(data []string) {
	threeDSpace := challenge17FormatData(data)
	for i := 0; i < 6; i++ {
		newSpace := challenge17CreateBlankSpace(threeDSpace)
		for x := 0; x < len(threeDSpace); x++ {
			for y := 0; y < len(threeDSpace[x]); y++ {
				for z := 0; z < len(threeDSpace[x][y]); z++ {
					count := challenge17CheckAdjacent(threeDSpace, x, y, z)
					if threeDSpace[x][y][z] == '#' {
						if count != 2 && count != 3 {
							newSpace[x][y][z] = '.'
						} else {
							newSpace[x][y][z] = '#'
						}
					} else {
						if count == 3 {
							newSpace[x][y][z] = '#'
						} else {
							newSpace[x][y][z] = '.'
						}
					}
				}
			}
		}

		threeDSpace = newSpace
	}

	finalCount := 0
	for x := 0; x < len(threeDSpace); x++ {
		for y := 0; y < len(threeDSpace[x]); y++ {
			for z := 0; z < len(threeDSpace[x][y]); z++ {
				if threeDSpace[x][y][z] == '#' {
					finalCount++
				}
			}
		}
	}

	println(finalCount)
}

func challenge17Part2(data []string) {
	fourDSpace := challenge17FormatData4D(data)
	for i := 0; i < 6; i++ {
		newSpace := challenge17CreateBlankSpace4D(fourDSpace)
		for x := 0; x < len(fourDSpace); x++ {
			for y := 0; y < len(fourDSpace[x]); y++ {
				for z := 0; z < len(fourDSpace[x][y]); z++ {
					for t := 0; t < len(fourDSpace[x][y][z]); t++ {
						count := challenge17CheckAdjacent4D(fourDSpace, x, y, z, t)
						if fourDSpace[x][y][z][t] == '#' {
							if count != 2 && count != 3 {
								newSpace[x][y][z][t] = '.'
							} else {
								newSpace[x][y][z][t] = '#'
							}
						} else {
							if count == 3 {
								newSpace[x][y][z][t] = '#'
							} else {
								newSpace[x][y][z][t] = '.'
							}

						}
					}
				}
			}
		}

		fourDSpace = newSpace
	}

	finalCount := 0
	for x := 0; x < len(fourDSpace); x++ {
		for y := 0; y < len(fourDSpace[x]); y++ {
			for z := 0; z < len(fourDSpace[x][y]); z++ {
				for t := 0; t < len(fourDSpace[x][y][z]); t++ {
					if fourDSpace[x][y][z][t] == '#' {
						finalCount++
					}
				}
			}
		}
	}

	println(finalCount)
}

func challenge17FormatData(data []string) [][][]byte {
	for i := 0; i < len(data); i++ {
		data[i] = "......" + data[i] + "......"
	}

	plane := [][]byte{}
	plainRow := ""
	for i := 0; i < len(data[0]); i++ {
		plainRow += "."
	}

	for i := 0; i < 6; i++ {
		plane = append(plane, []byte(plainRow))
	}

	for i := 0; i < len(data); i++ {
		plane = append(plane, []byte(data[i]))
	}

	for i := 0; i < 6; i++ {
		plane = append(plane, []byte(plainRow))
	}

	space3D := [][][]byte{}
	for i := 0; i < 6; i++ {
		emptyPlane := make([][]byte, len(plane))
		for i := 0; i < len(plane); i++ {
			emptyPlane[i] = []byte(plainRow)
		}

		space3D = append(space3D, emptyPlane)
	}

	space3D = append(space3D, plane)

	for i := 0; i < 6; i++ {
		emptyPlane := make([][]byte, len(plane))
		for i := 0; i < len(plane); i++ {
			emptyPlane[i] = []byte(plainRow)
		}

		space3D = append(space3D, emptyPlane)
	}

	return space3D
}

func challenge17FormatData4D(data []string) [][][][]byte {
	middle := challenge17FormatData(data)
	result := [][][][]byte{}
	for i := 0; i < 6; i++ {
		result = append(result, challenge17CreateBlankSpace(middle))
	}

	result = append(result, middle)
	for i := 0; i < 6; i++ {
		result = append(result, challenge17CreateBlankSpace(middle))
	}

	return result
}

func challenge17CreateBlankSpace(input [][][]byte) [][][]byte {
	blankRow := ""
	for i := 0; i < len(input[0][0]); i++ {
		blankRow += "."
	}

	output := [][][]byte{}
	for _, plane := range input {
		tempPlane := [][]byte{}
		for i := 0; i < len(plane); i++ {
			tempPlane = append(tempPlane, []byte(blankRow))
		}

		output = append(output, tempPlane)
	}

	return output
}

func challenge17CreateBlankSpace4D(input [][][][]byte) [][][][]byte {
	output := [][][][]byte{}
	for _, space := range input {
		output = append(output, challenge17CreateBlankSpace(space))
	}

	return output
}

func challenge17CheckAdjacent(input [][][]byte, x, y, z int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				} else if x+i >= 0 && x+i < len(input) &&
					y+j >= 0 && y+j < len(input[x]) &&
					z+k >= 0 && z+k < len(input[x][y]) &&
					input[x+i][y+j][z+k] == '#' {
					count++
				}
			}
		}
	}

	return count
}

func challenge17CheckAdjacent4D(input [][][][]byte, x, y, z, t int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					} else if x+i >= 0 && x+i < len(input) &&
						y+j >= 0 && y+j < len(input[x]) &&
						z+k >= 0 && z+k < len(input[x][y]) &&
						t+l >= 0 && t+l < len(input[x][y][z]) &&
						input[x+i][y+j][z+k][t+l] == '#' {
						count++
					}
				}
				if i == 0 && j == 0 && k == 0 {
					continue
				}
			}
		}
	}

	return count
}

var challenge17InputData []string = []string{
	"...#..#.",
	"..##.##.",
	"..#.....",
	"....#...",
	"#.##...#",
	"####..##",
	"...##.#.",
	"#.#.#...",
}

var challenge17TestData []string = []string{
	".#.",
	"..#",
	"###",
}
