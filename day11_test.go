package main

import (
	"testing"
)

const empty = "L"
const occupied = "#"

func day11part1(strs []string) int {

	seats := make([][]rune, len(strs))
	for i, s := range strs {
		seats[i] = make([]rune, len(s))
		for j, r := range s {
			seats[i][j] = r
		}
	}
	changed := true
	for {
		if !changed {
			break
		}
		changed = false
		clone := make([][]rune, len(seats))
		for i := range seats {
			clone[i] = make([]rune, len(seats[i]))
			for j := range seats[i] {
				clone[i][j] = seats[i][j]
				if string(seats[i][j]) == empty && allAdjacentSeatsEmpty(seats, i, j) {
					clone[i][j] = rune(occupied[0])
					changed = true
				}
				if string(seats[i][j]) == occupied && fourOrMoreAdjacentSeatsOccupied(seats, i, j) {
					clone[i][j] = rune(empty[0])
					changed = true
				}
			}
		}
		seats = clone
	}

	occupiedCount := 0
	for i := range seats {
		for j := range seats[i] {
			if string(seats[i][j]) == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func day11part2(strs []string) int {
	seats := make([][]rune, len(strs))
	for i, s := range strs {
		seats[i] = make([]rune, len(s))
		for j, r := range s {
			seats[i][j] = r
		}
	}
	changed := true
	for {
		if !changed {
			break
		}
		changed = false
		clone := make([][]rune, len(seats))
		for i := range seats {
			clone[i] = make([]rune, len(seats[i]))
			for j := range seats[i] {
				clone[i][j] = seats[i][j]
				if string(seats[i][j]) == empty && allAdjacentSeatsVisibleEmpty(seats, i, j) {
					clone[i][j] = rune(occupied[0])
					changed = true
				}
				if string(seats[i][j]) == occupied && fiveOrMoreAdjacentSeatsOccupied(seats, i, j) {
					clone[i][j] = rune(empty[0])
					changed = true
				}
			}
		}
		seats = clone
	}

	occupiedCount := 0
	for i := range seats {
		for j := range seats[i] {
			if string(seats[i][j]) == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func allAdjacentSeatsEmpty(seats [][]rune, i int, j int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < len(seats) && j+y >= 0 && j+y < len(seats[0]) {
				if string(seats[i+x][j+y]) == occupied {
					return false
				}
			}
		}
	}
	return true
}

func allAdjacentSeatsVisibleEmpty(seats [][]rune, i int, j int) bool {

	// right
	for x := 1; x+i < len(seats); x++ {
		if string(seats[i+x][j]) == occupied {
			return false
		}
		if string(seats[i+x][j]) == empty {
			break
		}
	}

	// left
	for x := 1; i-x >= 0; x++ {
		if string(seats[i-x][j]) == occupied {
			return false
		}
		if string(seats[i-x][j]) == empty {
			break
		}
	}

	// up
	for y := 1; j+y < len(seats[i]); y++ {
		if string(seats[i][j+y]) == occupied {
			return false
		}
		if string(seats[i][j+y]) == empty {
			break
		}
	}

	// down
	for y := 1; j-y >= 0; y++ {
		if string(seats[i][j-y]) == occupied {
			return false
		}
		if string(seats[i][j-y]) == empty {
			break
		}
	}

	// up right
	for x := 1; x+i < len(seats) && x+j < len(seats[i]); x++ {
		if string(seats[i+x][j+x]) == occupied {
			return false
		}
		if string(seats[i+x][j+x]) == empty {
			break
		}
	}

	// down left
	for x := 1; i-x >= 0 && j-x >= 0; x++ {
		if string(seats[i-x][j-x]) == occupied {
			return false
		}
		if string(seats[i-x][j-x]) == empty {
			break
		}
	}

	// up left
	for x := 1; i-x >= 0 && x+j < len(seats[i]); x++ {
		if string(seats[i-x][j+x]) == occupied {
			return false
		}
		if string(seats[i-x][j+x]) == empty {
			break
		}
	}

	// down right
	for x := 1; x+i < len(seats) && j-x >= 0; x++ {
		if string(seats[i+x][j-x]) == occupied {
			return false
		}
		if string(seats[i+x][j-x]) == empty {
			break
		}
	}

	return true
}

func fourOrMoreAdjacentSeatsOccupied(seats [][]rune, i int, j int) bool {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < len(seats) && j+y >= 0 && j+y < len(seats[0]) {
				if string(seats[i+x][j+y]) == occupied {
					count++
				}
			}
		}
	}
	return count >= 4
}

func fiveOrMoreAdjacentSeatsOccupied(seats [][]rune, i int, j int) bool {
	count := 0

	for x := 1; x+i < len(seats); x++ {
		if string(seats[i+x][j]) == occupied {
			count++
			break
		}
		if string(seats[i+x][j]) == empty {
			break
		}
	}
	for x := 1; i-x >= 0; x++ {
		if string(seats[i-x][j]) == occupied {
			count++
			break
		}
		if string(seats[i-x][j]) == empty {
			break
		}
	}
	for y := 1; j+y < len(seats[i]); y++ {
		if string(seats[i][j+y]) == occupied {
			count++
			break
		}
		if string(seats[i][j+y]) == empty {
			break
		}
	}
	for y := 1; j-y >= 0; y++ {
		if string(seats[i][j-y]) == occupied {
			count++
			break

		}
		if string(seats[i][j-y]) == empty {
			break
		}
	}
	for x := 1; x+i < len(seats) && x+j < len(seats[i]); x++ {
		if string(seats[i+x][j+x]) == occupied {
			count++
			break
		}
		if string(seats[i+x][j+x]) == empty {
			break
		}
	}
	for x := 1; i-x >= 0 && j-x >= 0; x++ {
		if string(seats[i-x][j-x]) == occupied {
			count++
			break
		}
		if string(seats[i-x][j-x]) == empty {
			break
		}
	}
	for x := 1; i-x >= 0 && x+j < len(seats[i]); x++ {
		if string(seats[i-x][j+x]) == occupied {
			count++
			break
		}
		if string(seats[i-x][j+x]) == empty {
			break
		}
	}
	for x := 1; x+i < len(seats) && j-x >= 0; x++ {
		if string(seats[i+x][j-x]) == occupied {
			count++
			break
		}
		if string(seats[i+x][j-x]) == empty {
			break
		}
	}

	return count >= 5
}

func TestAOCDay11(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{
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
		}, fn: day11part1, expected: 37},
		{name: "part1", input: readStringsInput("day11"), fn: day11part1, expected: 2166},
		{name: "part2sample", input: []string{
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
		}, fn: day11part2, expected: 26},
		{name: "part2", input: readStringsInput("day11"), fn: day11part2, expected: 1955},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
