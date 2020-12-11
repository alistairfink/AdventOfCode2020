package Challenges

import ()

/*
--- Day 11: Seating System ---
Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island
where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody
else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the
best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example,
the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow
a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions
	immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
This process continues for three more rounds:

#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##
#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##
#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##
At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change
state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

--- Part Two ---
As soon as people start to arrive, you realize your mistake. People don't just care about adjacent seats - they care about the first
seat they can see in each of those eight directions!

Now, instead of considering just the eight immediately adjacent seats, consider the first seat in each of those eight directions.
For example, the empty seat below would see eight occupied seats:

.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....
The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:

.............
.L.L.#.#.#.#.
.............
The empty seat below would see no occupied seats:

.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.
Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an occupied seat to become
empty (rather than four or more from the previous rules). The other rules still apply: empty seats that see no occupied seats become
occupied, seats matching no rule don't change, and floor never changes.

Given the same starting layout as above, these new rules cause the seating area to shift around as follows:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
Again, at this point, people stop shifting around and the seating area reaches equilibrium. Once this occurs, you count 26 occupied seats.

Given the new visibility method and the rule change for occupied seats becoming empty, once equilibrium is reached, how many seats end up occupied?
*/

func Challenge11(useTestData bool) {
	var data []string
	if useTestData {
		data = challenge11TestData
	} else {
		data = challenge11InputData
	}

	println("Part 1")
	challenge11Part1(data)
	println("Part 2")
	challenge11Part2(data)
}

func challenge11Part1(data []string) {
	counter := 0
	changes := true
	result := 0
	for changes {
		changes = false
		dataCopy := make([]string, len(data))
		result = 0
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				if data[i][j] == '.' {
					dataCopy[i] += "."
					continue
				}

				surroundingOccupied := challegne11Part1CheckSurrounding(data, i, j)
				if data[i][j] == 'L' {
					if surroundingOccupied == 0 {
						dataCopy[i] += "#"
						changes = true
						result++
					} else {
						dataCopy[i] += "L"
					}
				} else {
					if surroundingOccupied >= 4 {
						dataCopy[i] += "L"
						changes = true
					} else {
						dataCopy[i] += "#"
						result++
					}
				}
			}
		}

		data = dataCopy
		counter++
	}

	println(result)
}

func challenge11Part2(data []string) {
	counter := 0
	changes := true
	result := 0
	for changes {
		changes = false
		dataCopy := make([]string, len(data))
		result = 0
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				if data[i][j] == '.' {
					dataCopy[i] += "."
					continue
				}

				surroundingOccupied := challegne11Part2CheckSurrounding(data, i, j)
				if data[i][j] == 'L' {
					if surroundingOccupied == 0 {
						dataCopy[i] += "#"
						changes = true
						result++
					} else {
						dataCopy[i] += "L"
					}
				} else {
					if surroundingOccupied >= 5 {
						dataCopy[i] += "L"
						changes = true
					} else {
						dataCopy[i] += "#"
						result++
					}
				}
			}
		}

		data = dataCopy
		counter++
	}

	println(result)
}

func challegne11Part1CheckSurrounding(data []string, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) &&
				x+i >= 0 &&
				x+i < len(data) &&
				y+j >= 0 &&
				y+j < len(data[x]) &&
				data[x+i][y+j] == '#' {
				count++
			}
		}
	}

	return count
}

func challegne11Part2CheckSurrounding(data []string, x, y int) int {
	pairs := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{0, -1},
		[]int{1, -1},
	}
	count := 0
	for i, pair := range pairs {
		left, right := false, false
		for !left || !right {
			if !left &&
				x+pair[0] >= 0 &&
				x+pair[0] < len(data) &&
				y+pair[1] >= 0 &&
				y+pair[1] < len(data[x]) {
				if data[x+pair[0]][y+pair[1]] == '#' {
					count++
					left = true
				} else if data[x+pair[0]][y+pair[1]] == 'L' {
					left = true
				}
			} else {
				left = true
			}

			if !right &&
				x-pair[0] >= 0 &&
				x-pair[0] < len(data) &&
				y-pair[1] >= 0 &&
				y-pair[1] < len(data[x]) {
				if data[x-pair[0]][y-pair[1]] == '#' {
					count++
					right = true
				} else if data[x-pair[0]][y-pair[1]] == 'L' {
					right = true
				}
			} else {
				right = true
			}

			if pair[0] < 0 {
				pairs[i][0]--
			} else if pair[0] > 0 {
				pairs[i][0]++
			}

			if pair[1] < 0 {
				pairs[i][1]--
			} else if pair[1] > 0 {
				pairs[i][1]++
			}
		}
	}

	return count
}

var challenge11InputData []string = []string{
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL",
	".......L.L.....LL..L....LLL...L...........L.....L...LL...L.....L..LL..L..L.L...L........L.L",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLL.LL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLL..LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"L.LLL..L...L.L.......L....L..LLL........LLL..L..L....LLLLLLL.......L.L...L........LL.L..L..",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	".L..LLLLL.....L..L...L....L..LLL.L.....L......L..LL......L.L...LLL..LLLL......LL.L.L...L...",
	"LLLLLLL.LL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LL.LLLLLLL.LLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLL..LLLLLLL.LLLLLL",
	"LLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLL.L",
	"LLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLL.L",
	"LLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"L..........LLL....LL...............L.....L.......L......L..L.L...L..LL.......L..L...L....LL",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL..LLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLL.LLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLL.LLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLL",
	"L...L..L..L.L...L..L.L.....L....L..L...L....L.LL..L.L...L.L.LL....LL..L..L......L.L...L..L.",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLL.L.LLLLL.LLLLL..LLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLL.L.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL.LL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"....L....L.L..LL.L...LLL.L.L..L........L..L..L....L..LL..L.L....L......LL.L.......L.LL....L",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"..L..L..LLL..L...L.......LL..L..LLLL.L....L.LLL.L...L...LL...LL.L.L......L.......L....LL.L.",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LL.LL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLL..LLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"L...LLLLLLLL..L....L.L....L...LL..L.LL..L.L.L.L..L....L.....LL.........LL.L.....L.LL..L...L",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.L.LLLLLLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLL.L.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.L.LLLLLLLLLLLLLLL.LLLLL.L.LLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLL.LLL.LLLLL.LLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLL..LLLLLL",
	".L..LL....L..L.LL.......L.L.L.........L...L.L...LL......LL..L...L..L.L.L..L........L....L.L",
	"LLLLLLL.LLLLLLLL.LLLLLLL..LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.L.LLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"..LL..L.LL..L........L.L..L.....L.L..L.L.....LL.LL...L....L..LL.L...LLLLL.L.L.....L.......L",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLL.LL",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLLLLL",
	"L.L..........L.L...LL...L....L..L.L..L.........LL...L..L...L...L.L.LLL.L.....LL....L......L",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLL..LLLLLLLLLLLLLL.LLLLLLLLLLLLL.LL.LLLLLLLLLLL.LLLLLL",
	"LLLLLL.LL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL..LLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLL.LL.LL.LLLLLLLLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLL",
	"LLLLLL.LLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLL.LLLLLL",
	"LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLL",
}

var challenge11TestData []string = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}
