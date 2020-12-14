package main

import (
	"strconv"
	"strings"
	"testing"
)

func day14part1(strs []string) int64 {

	mem := make([]int64, 100000)

	var mask string
	for _, s := range strs {
		if strings.Contains(s, "mask") {
			mask = reverse(strings.Trim(strings.Split(s, "=")[1], " "))
			continue
		}

		op := strings.Split(s, "=")
		n, _ := strconv.ParseInt(strings.Trim(op[1], " "), 10, 64)

		for i, v := range mask {
			if string(v) == "1" {
				n = n | (1 << (i))
			} else if string(v) == "0" {
				n = n &^ (1 << (i))
			}
		}

		op[0] = strings.Replace(op[0], "mem[", "", -1)
		op[0] = strings.Replace(op[0], "]", "", -1)
		idx, _ := strconv.Atoi(strings.Trim(op[0], " "))
		mem[idx] = n
	}

	sum := int64(0)
	for _, n := range mem {
		sum += n
	}
	return sum
}

func reverse(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func day14part2(strs []string) int64 {

	mem := make(map[int64]int64)

	var mask string
	for _, s := range strs {
		if strings.Contains(s, "mask") {
			mask = reverse(strings.Trim(strings.Split(s, "=")[1], " "))
			continue
		}

		op := strings.Split(s, "=")
		n, _ := strconv.ParseInt(strings.Trim(op[1], " "), 10, 64)

		op[0] = strings.Replace(op[0], "mem[", "", -1)
		op[0] = strings.Replace(op[0], "]", "", -1)
		idx, _ := strconv.ParseInt(strings.Trim(op[0], " "), 10, 64)

		addrs := make([]int64, 1)
		addrs[0] = idx

		for i, v := range mask {
			if string(v) == "0" {
				continue
			} else if string(v) == "1" {
				for a, addr := range addrs {
					addrs[a] = addr | (1 << (i))
				}
			} else if string(v) == "X" {
				for a, addr := range addrs {
					addrs[a] = addr | (1 << (i))
				}
				for _, addr := range addrs {
					a := addr &^ (1 << (i))
					addrs = append(addrs, a)
				}
			}
		}

		for _, addr := range addrs {
			mem[addr] = n
		}
	}

	sum := int64(0)
	for _, n := range mem {
		sum += n
	}
	return sum
}

func TestAOCDay14(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			"mem[8] = 11",
			"mem[7] = 101",
			"mem[8] = 0",
		}, fn: day14part1, expected: int64(165)},
		{name: "part1", input: readStringsInput("day14"), fn: day14part1, expected: int64(16003257187056)},
		{name: "part2sample", input: []string{
			"mask = 000000000000000000000000000000X1001X",
			"mem[42] = 100",
			"mask = 00000000000000000000000000000000X0XX",
			"mem[26] = 1",
		}, fn: day14part2, expected: int64(208)},
		{name: "part2", input: readStringsInput("day14"), fn: day14part2, expected: 3219837697833},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
