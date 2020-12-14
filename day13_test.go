package main

import (
	"strconv"
	"strings"
	"testing"
)

func day13part1(strs []string) int64 {
	t, _ := strconv.Atoi(strs[0])
	originalT := t
	raw := strings.Split(strs[1], ",")

	var ids = make([]int, 0)
	for _, r := range raw {
		if r == "x" {
			continue
		}
		n, _ := strconv.Atoi(r)
		ids = append(ids, n)
	}

	earlist := -1
	earlistId := 0

	for i := len(ids); i > 0; {
		for _, id := range ids {
			if earlist != -1 {
				i--
			}
			if t%id == 0 && (earlist == -1 || t < earlist) {
				earlist = t
				earlistId = id

			}
		}
		t++
	}

	return int64((earlist - originalT) * earlistId)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func day13part2(strs []string) int64 {
	raw := strings.Split(strs[1], ",")
	var ids = make(map[int64]int64)

	var whatever = make([]int64, 0)

	for i, r := range raw {
		if r == "x" {
			continue
		}
		id, _ := strconv.Atoi(r)
		ids[int64(i)] = int64(id)
		whatever = append(whatever, int64(id-i))
	}

	t := int64(0)

	search := ids[0]

	for ; ; t += search {
		valid := true

		n := int64(0)

		for k, v := range ids {
			if (t+k)%v != 0 {
				valid = false
				break
			}
			if n == 0 {
				n = v
			} else {
				n = LCM(n, v)
				if n > search {
					search = n
				}
			}
		}
		if valid {
			break
		}
	}
	return t
}

func TestAOCDay13(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			"939",
			"7,13,x,x,59,x,31,19",
		}, fn: day13part1, expected: int64(295)},
		{name: "part1", input: readStringsInput("day13"), fn: day13part1, expected: int64(3865)},
		{name: "part2sample", input: []string{
			"939",
			"7,13,x,x,59,x,31,19",
		}, fn: day13part2, expected: int64(1068781)},
		{name: "part2sample2", input: []string{
			"939",
			"17,x,13,19",
		}, fn: day13part2, expected: int64(3417)},
		{name: "part2sample2", input: []string{
			"939",
			"67,7,59,61",
		}, fn: day13part2, expected: int64(754018)},
		{name: "part2sample2", input: []string{
			"939",
			"67,x,7,59,61",
		}, fn: day13part2, expected: int64(779210)},
		{name: "part2sample2", input: []string{
			"939",
			"67,7,x,59,61",
		}, fn: day13part2, expected: int64(1261476)},
		{name: "part2sample2", input: []string{
			"939",
			"1789,37,47,1889",
		}, fn: day13part2, expected: int64(1202161486)},
		{name: "part2", input: readStringsInput("day13"), fn: day13part2, expected: 415579909629976},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
