package main

import (
	"math"
	"sort"
	"testing"
)

func HighestSeatId(input []string) int {
	max := -1
	for _, s := range input {
		id := ReadRow(s)*8 + ReadColumn(s)
		if id > max {
			max = id
		}
	}
	return max
}

func FindMySeat(input []string) int {
	seatIds := make(sort.IntSlice, len(input))
	for i, s := range input {
		seatIds[i] = ReadRow(s)*8 + ReadColumn(s)
	}
	seatIds.Sort()

	for i, id := range seatIds {
		if id+1 != seatIds[i+1] {
			return id + 1
		}
	}

	return -1
}

func ReadColumn(str string) int {
	min := 0
	max := 8
	for _, s := range str[7:] {
		if string(s) == "L" {
			max = int(math.Floor(float64((min + max) / 2)))
		} else {
			min = int(math.Ceil(float64((min + max) / 2)))
		}
	}
	return min
}

func ReadRow(str string) int {
	min := 0
	max := 127
	for _, s := range str[:7] {
		if string(s) == "F" {
			max = int(math.Floor(float64((min + max) / 2)))
		} else {
			min = int(math.Ceil(float64((min + max) / 2)))
		}
	}
	return max
}

func TestAOCDay5(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{"BFFFBBFRRR"}, fn: HighestSeatId, expected: 567},
		{name: "part1sample", input: []string{"FFFBBBFRRR"}, fn: HighestSeatId, expected: 119},
		{name: "part1sample", input: []string{"BBFFBBFRLL"}, fn: HighestSeatId, expected: 820},
		{name: "part1", input: readStringsInput("day5"), fn: HighestSeatId, expected: 933},
		{name: "part2", input: readStringsInput("day5"), fn: FindMySeat, expected: 711},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
