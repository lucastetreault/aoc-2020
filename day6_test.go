package main

import (
	"testing"
)

func SumAnswers(strings []string) int {
	sum := 0
	m := make(map[string]string)
	for _, s := range strings {
		if len(s) == 0 {
			sum += len(m)
			m = make(map[string]string)
			continue
		}
		for _, c := range s {
			m[string(c)] = string(c)
		}
	}
	return sum
}

func SumGroupAnswers(strings []string) int {
	sum := 0
	groupSize := 0
	m := make(map[string]int)
	for _, s := range strings {
		if len(s) == 0 {

			for _, v := range m {
				if v == groupSize {
					sum++
				}
			}

			m = make(map[string]int)
			groupSize = 0
			continue
		} else {
			groupSize++
		}
		for _, c := range s {
			if v, ok := m[string(c)]; ok {
				m[string(c)] = v + 1
			} else {
				m[string(c)] = 1
			}
		}
	}
	return sum
}

func TestAOCDay6(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
			"",
		}, fn: SumAnswers, expected: 11},
		//{name: "part1sample", input: []string{"FFFBBBFRRR"}, fn: HighestSeatId, expected: 119},
		//{name: "part1sample", input: []string{"BBFFBBFRLL"}, fn: HighestSeatId, expected: 820},
		{name: "part1", input: readStringsInput("day6"), fn: SumAnswers, expected: 6273},
		{name: "part2sample", input: []string{
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
			"",
		}, fn: SumGroupAnswers, expected: 6},
		{name: "part2", input: readStringsInput("day6"), fn: SumGroupAnswers, expected: 3254},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
