package main

import (
	"sort"
	"testing"
)

//1, 2, 3 jolts lower
// device is rated for 3 jolts higher
// if adapters are 3, 6, 9 we can handle 12
// outlet is 0

func day10part1(adapters []int64) int64 {
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	adapters = append([]int64{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	oneJoltDiff := int64(0)
	threeJoltDiff := int64(0)
	for i, a := range adapters[:len(adapters)-1] {
		if adapters[i+1]-a == 1 {
			oneJoltDiff++
		} else if adapters[i+1]-a == 3 {
			threeJoltDiff++
		}
	}
	return oneJoltDiff * threeJoltDiff
}

var opsCount = 0
var valids map[int64]int64

func day10part2(adapters []int64) int64 {
	valids = make(map[int64]int64)
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	adapters = append([]int64{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	opsCount = 0
	return validateJoltageChain(adapters, adapters[len(adapters)-1])
}

func validateJoltageChain(adapters []int64, i int64) int64 {
	if i == 0 {
		return 1
	}
	found := false
	for _, a := range adapters {
		if a == i {
			found = true
			break
		}
	}
	if !found {
		return 0
	}

	if _, ok := valids[i-1]; !ok {
		valids[i-1] = validateJoltageChain(adapters, i-1)
	}
	if _, ok := valids[i-2]; !ok {
		valids[i-2] = validateJoltageChain(adapters, i-2)
	}
	if _, ok := valids[i-3]; !ok {
		valids[i-3] = validateJoltageChain(adapters, i-3)
	}

	return valids[i-1] + valids[i-2] + valids[i-3]
}

func TestAOCDay10(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		fn       func([]int64) int64
		expected int64
	}{
		{name: "part1sample", input: []int64{
			1,
			4,
			5,
			6,
			7,
			10,
			11,
			12,
			15,
			16,
			19,
		}, fn: day10part1, expected: 35},
		{name: "part1sample2", input: []int64{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}, fn: day10part1, expected: 22 * 10},
		{name: "part1", input: readInt64Input("day10"), fn: day10part1, expected: 2080},
		{name: "part2sample", input: []int64{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		}, fn: day10part2, expected: 8},
		{name: "part1sample2", input: []int64{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}, fn: day10part2, expected: 19208},
		{name: "part2", input: readInt64Input("day10"), fn: day10part2, expected: 6908379398144},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
