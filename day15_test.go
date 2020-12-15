package main

import (
	"testing"
)

func day15part1(input []int) int64 {

	memory := make(map[int]int)

	for i, n := range input[:len(input)-1] {
		memory[n] = i + 1
	}

	next := input[len(input)-1]

	for turn := len(input); turn < 2020; turn++ {
		if last, ok := memory[next]; !ok {
			memory[next] = turn
			next = 0
		} else {
			memory[next] = turn
			next = turn - last
		}
	}

	return int64(next)
}

func day15part2(input []int) int64 {
	memory := make(map[int64]int64)

	for i, n := range input[:len(input)-1] {
		memory[int64(n)] = int64(i + 1)
	}

	next := int64(input[len(input)-1])

	for turn := int64(len(input)); turn < 30000000; turn++ {
		if last, ok := memory[next]; !ok {
			memory[next] = turn
			next = 0
		} else {
			memory[next] = turn
			next = turn - last
		}
	}

	return next
}

func TestAOCDay15(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func([]int) int64
		expected int64
	}{
		//{name: "part1sample", input: []int{0, 3, 6}, fn: day15part1, expected: 1},
		{name: "part1sample", input: []int{1, 3, 2}, fn: day15part1, expected: 1},
		{name: "part1sample", input: []int{2, 1, 3}, fn: day15part1, expected: 10},
		{name: "part1sample", input: []int{1, 2, 3}, fn: day15part1, expected: 27},
		{name: "part1sample", input: []int{2, 3, 1}, fn: day15part1, expected: 78},
		{name: "part1sample", input: []int{3, 2, 1}, fn: day15part1, expected: 438},
		{name: "part1sample", input: []int{3, 1, 2}, fn: day15part1, expected: 1836},
		{name: "part1", input: []int{5, 2, 8, 16, 18, 0, 1}, fn: day15part1, expected: 517},
		{name: "part2sample", input: []int{0, 3, 6}, fn: day15part2, expected: 175594},
		{name: "part2sample", input: []int{1, 3, 2}, fn: day15part2, expected: 2578},
		{name: "part2sample", input: []int{2, 1, 3}, fn: day15part2, expected: 3544142},
		{name: "part2sample", input: []int{1, 2, 3}, fn: day15part2, expected: 261214},
		{name: "part2sample", input: []int{2, 3, 1}, fn: day15part2, expected: 6895259},
		{name: "part2sample", input: []int{3, 2, 1}, fn: day15part2, expected: 18},
		{name: "part2sample", input: []int{3, 1, 2}, fn: day15part2, expected: 362},
		{name: "part2", input: []int{5, 2, 8, 16, 18, 0, 1}, fn: day15part2, expected: 1047739},
		//{name: "part2sample", input: []string{
		//	"mask = 000000000000000000000000000000X1001X",
		//	"mem[42] = 100",
		//	"mask = 00000000000000000000000000000000X0XX",
		//	"mem[26] = 1",
		//}, fn: day15part2, expected: int64(208)},
		//{name: "part2", input: readStringsInput("day15"), fn: day15part2, expected: 3219837697833},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
