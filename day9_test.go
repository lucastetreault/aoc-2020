package main

import (
	"sort"
	"testing"
)

func part1(size int) func(nums []int64) int64 {
	return func(nums []int64) int64 {
		for i, n := range nums[size:] {
			if !isValid(nums[i:i+size], n) {
				return n
			}
		}
		return -1
	}
}

func part2(size int) func(nums []int64) int64 {
	return func(nums []int64) int64 {
		p1 := part1(size)(nums)

		arr := make([]int64, 0)
		sum := int64(0)
		for _, n := range nums {
			for ; sum > p1; {
				sum -= arr[0]
				arr = arr[1:]
			}
			if sum < p1 {
				arr = append(arr, n)
				sum += n
			}
			if sum == p1 {
				sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
				return arr[0] + arr[len(arr)-1]
			}
		}
		return -1
	}
}

func isValid(preamble []int64, num int64) bool {
	for _, a := range preamble {
		for _, b := range preamble {
			if a+b == num {
				return true
			}
		}
	}
	return false
}

func TestAOCDay9(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		fn       func([]int64) int64
		expected int64
	}{
		{name: "part1sample", input: []int64{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}, fn: part1(5), expected: 127},
		{name: "part1", input: readInt64Input("day9"), fn: part1(25), expected: 36845998},
		{name: "part2sample", input: []int64{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}, fn: part2(5), expected: 62},
		{name: "part2", input: readInt64Input("day9"), fn: part2(25), expected: 4830226},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
