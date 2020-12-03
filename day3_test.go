package main

import (
	"testing"
)

func CountTreesWithSingleSlope(input []string) int64 {
	return countTrees(input, slopeFunc(3, 1))
}

func CountTreesAllSlopes(input []string) int64 {
	return countTrees(input, slopeFunc(1, 1)) *
		countTrees(input, slopeFunc(3, 1)) *
		countTrees(input, slopeFunc(5, 1)) *
		countTrees(input, slopeFunc(7, 1)) *
		countTrees(input, slopeFunc(1, 2))
}

func countTrees(input []string, s func(input []string, j int, i int) (int, int)) int64 {
	treeCount := int64(0)
	j := 0
	for i := 0; i < len(input)-1; {
		j, i = s(input, j, i)
		if string(input[i][j]) == "#" {
			treeCount++
		}
	}
	return treeCount
}

func slopeFunc(right int, down int) func(input []string, j int, i int) (int, int) {
	return func(input []string, j int, i int) (int, int) {
		j += right
		for ; j >= len(input[i]); {
			doubleHillSize(input)
		}
		i += down
		return j, i
	}
}

func doubleHillSize(input []string) {
	for i, s := range input {
		input[i] = s + s
	}
}

func TestAOCDay3(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}, fn: CountTreesWithSingleSlope, expected: 7},
		{name: "part1", input: readStringsInput("day3"), fn: CountTreesWithSingleSlope, expected: 214},
		{name: "part2sample", input: []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}, fn: CountTreesAllSlopes, expected: 336},
		{name: "part2", input: readStringsInput("day3"), fn: CountTreesAllSlopes, expected: 8336352024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
