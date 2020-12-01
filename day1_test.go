package main

import (
	"testing"
)

func findTwoNumbersThatSumTo2020(input []int64) int64 {
	for i, a := range input {
		for _, b := range input[i+1:] {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return 0
}

func findThreeNumbersThatSumTo2020(input []int64) int64 {
	for i, a := range input {
		for j, b := range input[i+1:] {
			for _, c := range input[i+j+1:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return 0
}

func TestAOCDay1(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		fn       func([]int64) int64
		expected int64
	}{
		{name: "sample", input: []int64{1721, 979, 366, 299, 675, 1456}, fn: findTwoNumbersThatSumTo2020, expected: 514579},
		{name: "part1", input: readInt64Input("day1"), fn: findTwoNumbersThatSumTo2020, expected: 902451},
		{name: "part2", input: readInt64Input("day1"), fn: findThreeNumbersThatSumTo2020, expected: 85555470},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}

}
