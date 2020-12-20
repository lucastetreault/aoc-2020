package main

import (
	"testing"
)

const active = "#"
const inactive = "."

func day17part1(strs []string) int64 {
	universe := make(map[int]map[int]map[int]string)

	for x, s := range strs {
		for y, r := range s {
			if string(r) == active {
				activate(universe, x, y, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		universe = cycle(universe)
	}

	count := int64(0)
	for _, x := range universe {
		for _, y := range x {
			count += int64(len(y))
		}
	}

	return count
}

func activate4d(universe map[int]map[int]map[int]map[int]string, x, y, z, w int) {
	if _, ok := universe[x]; !ok {
		universe[x] = make(map[int]map[int]map[int]string)
	}
	if _, ok := universe[x][y]; !ok {
		universe[x][y] = make(map[int]map[int]string)
	}
	if _, ok := universe[x][y][z]; !ok {
		universe[x][y][z] = make(map[int]string)
	}
	universe[x][y][z][w] = active
}

func state4d(universe map[int]map[int]map[int]map[int]string, x, y, z, w int) string {
	if _, ok := universe[x]; !ok {
		return inactive
	}
	if _, ok := universe[x][y]; !ok {
		return inactive
	}
	if _, ok := universe[x][y][z]; !ok {
		return inactive
	}
	if _, ok := universe[x][y][z][w]; !ok {
		return inactive
	}

	return active
}

func incrementActiveNeighbourCount4d(universe map[int]map[int]map[int]map[int]int, x, y, z, w int) {
	if _, ok := universe[x]; !ok {
		universe[x] = make(map[int]map[int]map[int]int)
	}
	if _, ok := universe[x][y]; !ok {
		universe[x][y] = make(map[int]map[int]int)
	}
	if _, ok := universe[x][y][z]; !ok {
		universe[x][y][z] = make(map[int]int)
	}
	universe[x][y][z][w]++
}

func cycle4d(universe map[int]map[int]map[int]map[int]string) map[int]map[int]map[int]map[int]string {
	activeNeighbourCount := make(map[int]map[int]map[int]map[int]int)

	for x := range universe {
		for y := range universe[x] {
			for z := range universe[x][y] {
				for w := range universe[x][y][z] {

					for i := -1; i <= 1; i++ {
						for j := -1; j <= 1; j++ {
							for k := -1; k <= 1; k++ {
								for l := -1; l <= 1; l++ {
									if i == 0 && j == 0 && k == 0 && l == 0 {
										continue
									}
									incrementActiveNeighbourCount4d(activeNeighbourCount, x+i, y+j, z+k, w+l)
								}
							}
						}
					}
				}
			}
		}
	}
	clone := make(map[int]map[int]map[int]map[int]string)

	for x := range activeNeighbourCount {
		for y := range activeNeighbourCount[x] {
			for z := range activeNeighbourCount[x][y] {
				for w := range activeNeighbourCount[x][y][z] {
					if state4d(universe, x, y, z, w) == active && (activeNeighbourCount[x][y][z][w] == 2 || activeNeighbourCount[x][y][z][w] == 3) {
						activate4d(clone, x, y, z, w)
					} else if state4d(universe, x, y, z, w) == inactive && activeNeighbourCount[x][y][z][w] == 3 {
						activate4d(clone, x, y, z, w)
					}
				}
			}
		}
	}

	return clone
}
func day17part2(strs []string) int64 {
	universe := make(map[int]map[int]map[int]map[int]string)

	for x, s := range strs {
		for y, r := range s {
			if string(r) == active {
				activate4d(universe, x, y, 0, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		universe = cycle4d(universe)
	}

	count := int64(0)
	for _, x := range universe {
		for _, y := range x {
			for _, z := range y {
				count += int64(len(z))
			}
		}
	}

	return count
}

func activate(universe map[int]map[int]map[int]string, x, y, z int) {
	if _, ok := universe[x]; !ok {
		universe[x] = make(map[int]map[int]string)
	}
	if _, ok := universe[x][y]; !ok {
		universe[x][y] = make(map[int]string)
	}
	universe[x][y][z] = active
}

func state(universe map[int]map[int]map[int]string, x, y, z int) string {
	if _, ok := universe[x]; !ok {
		return inactive
	}
	if _, ok := universe[x][y]; !ok {
		return inactive
	}

	if _, ok := universe[x][y][z]; !ok {
		return inactive
	}

	return active
}

func incrementActiveNeighbourCount(universe map[int]map[int]map[int]int, x, y, z int) {
	if _, ok := universe[x]; !ok {
		universe[x] = make(map[int]map[int]int)
	}
	if _, ok := universe[x][y]; !ok {
		universe[x][y] = make(map[int]int)
	}
	universe[x][y][z]++
}

func cycle(universe map[int]map[int]map[int]string) map[int]map[int]map[int]string {
	activeNeighbourCount := make(map[int]map[int]map[int]int)

	for x := range universe {
		for y := range universe[x] {
			for z := range universe[x][y] {

				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						for k := -1; k <= 1; k++ {
							if i == 0 && j == 0 && k == 0 {
								continue
							}
							incrementActiveNeighbourCount(activeNeighbourCount, x+i, y+j, z+k)
						}
					}
				}
			}
		}
	}

	clone := make(map[int]map[int]map[int]string)

	for x := range activeNeighbourCount {
		for y := range activeNeighbourCount[x] {
			for z := range activeNeighbourCount[x][y] {
				if state(universe, x, y, z) == active && (activeNeighbourCount[x][y][z] == 2 || activeNeighbourCount[x][y][z] == 3) {
					activate(clone, x, y, z)
				} else if state(universe, x, y, z) == inactive && activeNeighbourCount[x][y][z] == 3 {
					activate(clone, x, y, z)
				}
			}
		}
	}

	return clone
}

func TestAOCDay17(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			".#.",
			"..#",
			"###",
		}, fn: day17part1, expected: int64(112)},
		{name: "part1", input: readStringsInput("day17"), fn: day17part1, expected: int64(291)},
		{name: "part2sample", input: []string{
			".#.",
			"..#",
			"###",
		}, fn: day17part2, expected: int64(848)},
		{name: "part1", input: readStringsInput("day17"), fn: day17part2, expected: int64(1524)},
		//{name: "part2sample", input: []string{
		//	"1 + 2 * 3 + 4 * 5 + 6",
		//}, fn: day17part2, expected: int64(231)},
		//{name: "part2", input: readStringsInput("day17"), fn: day17part2, expected: 158173007916215},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
