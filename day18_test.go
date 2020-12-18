package main

import (
	"strconv"
	"strings"
	"testing"
)

func day18part1(strs []string) int64 {

	sum := int64(0)

	for _, raw := range strs {
		var idx = 0
		sum += evaluate(raw, &idx)
	}

	return sum
}

func plus(left, right int64) int64 {
	return left + right
}

func times(left, right int64) int64 {
	return left * right
}

func evaluate(s string, idx *int) int64 {
	var left *int64
	var operator func(int64, int64) int64
	for ; *idx < len(s); *idx++ {
		r := s[*idx]
		value := int64(0)
		if string(r) == " " {
			continue
		} else if string(r) == "(" {
			*idx++
			value = evaluate(s, idx)
			if left == nil {
				left = &value
			} else {
				// do the calculation once we figured out
				value = operator(*left, value)
				left = &value
			}
		} else if string(r) == ")" {
			return *left
		} else if string(r) == "+" {
			operator = plus
		} else if string(r) == "*" {
			operator = times
		} else {
			value, _ = strconv.ParseInt(string(r), 10, 64)
			if left == nil {
				left = &value
			} else {
				// do the calculation once we figured out
				value = operator(*left, value)
				left = &value
			}
		}
	}
	return *left
}

func day18part2(strs []string) int64 {

	sum := int64(0)

	for _, raw := range strs {
		raw = strings.Replace(raw, " ", "", -1)
		sum += evaluateAdditionBeforeMultiplication(raw)
	}

	return sum
}

func evaluateAdditionBeforeMultiplication(s string) int64 {
	var left *int64
	var operator func(int64, int64) int64
	for i := 0; i < len(s); i++ {
		r := string(s[i])
		value := int64(0)
		if r == "(" {
			count := 1
			idx := 0
			for i, next := range s[i+1:] {
				idx = i
				next := string(next)
				if next == "(" {
					count++
				} else if next == ")" {
					count--
					if count == 0 {
						break
					}
				}
			}

			str := s[i+1 : i+1+idx]
			value = evaluateAdditionBeforeMultiplication(str)
			i += idx+1
			if left == nil {
				left = &value
			} else {
				// do the calculation once we figured out
				value = operator(*left, value)
				left = &value
			}
		} else if r == ")" {
			return *left
		} else if r == "+" {
			operator = plus
		} else if r == "*" {
			operator = times
			value = evaluateAdditionBeforeMultiplication(s[i+1:])
			if left == nil {
				left = &value
			} else {
				// do the calculation once we figured out
				value = operator(*left, value)
				left = &value
			}
			break
		} else {
			value, _ = strconv.ParseInt(r, 10, 64)
			if left == nil {
				left = &value
			} else {
				// do the calculation once we figured out
				value = operator(*left, value)
				left = &value
			}
		}
	}
	return *left
}

func TestAOCDay18(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			"1 + 2 * 3 + 4 * 5 + 6",
		}, fn: day18part1, expected: int64(71)},
		{name: "part1sample", input: []string{
			"1 + (2 * 3) + (4 * (5 + 6))",
		}, fn: day18part1, expected: int64(51)},
		{name: "part1", input: readStringsInput("day18"), fn: day18part1, expected: int64(24650385570008)},
		{name: "part2sample", input: []string{
			"1 + 2 * 3 + 4 * 5 + 6",
		}, fn: day18part2, expected: int64(231)},
		{name: "part2sample", input: []string{
			"1 + (2 * 3) + (4 * (5 + 6))",
		}, fn: day18part2, expected: int64(51)},
		{name: "part2sample", input: []string{
			"2 * 3 + (4 * 5)",
		}, fn: day18part2, expected: int64(46)},
		{name: "part2sample", input: []string{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
		}, fn: day18part2, expected: int64(1445)},
		{name: "part2sample", input: []string{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		}, fn: day18part2, expected: int64(669060)},
		{name: "part2sample", input: []string{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
		}, fn: day18part2, expected: int64(23340)},
		{name: "part2", input: readStringsInput("day18"), fn: day18part2, expected: 158183007916215},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
