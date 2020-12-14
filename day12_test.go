package main

import (
	"math"
	"strconv"
	"testing"
)

const N = "N"
const S = "S"
const E = "E"
const W = "W"
const L = "L"
const R = "R"
const F = "F"

var directions = []string{N, E, S, W}

func day12part1(strs []string) int {
	direction := E

	x := 0
	y := 0

	for _, s := range strs {
		n, _ := strconv.Atoi(s[1:])
		switch string(s[0]) {
		case L:
			var cdi int
			for i, d := range directions {
				if d == direction {
					cdi = i
					break
				}
			}
			cdi -= n / 90
			cdi = (4 + cdi) % 4
			direction = directions[cdi]
		case R:
			var cdi int
			for i, d := range directions {
				if d == direction {
					cdi = i
					break
				}
			}
			cdi += n / 90
			cdi %= 4
			direction = directions[cdi]
		case F:
			x, y = move(direction, x, y, n)
		default:
			x, y = move(string(s[0]), x, y, n)
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func day12part2(strs []string) int {
	x := 0
	y := 0

	wx := 10
	wy := 1

	for _, s := range strs {
		n, _ := strconv.Atoi(s[1:])
		switch string(s[0]) {
		case L:
			wx, wy = rotate(wx, wy, n)
		case R:
			wx, wy = rotate(wx, wy, n*-1)
		case F:
			x, y = moveToWaypoint(wx, wy, x, y, n)
		default:
			wx, wy = move(string(s[0]), wx, wy, n)
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func rotate(wx int, wy int, n int) (int, int) {
	return int(math.Round(float64(wx)*math.Cos(float64(n)*(math.Pi/180)) - float64(wy)*math.Sin(float64(n)*(math.Pi/180)))), int(math.Round(float64(wx)*math.Sin(float64(n)*(math.Pi/180)) + float64(wy)*math.Cos(float64(n)*(math.Pi/180))))
}

func moveToWaypoint(wx int, wy int, x int, y int, n int) (int, int) {
	return x + (wx * n), y + (wy * n)
}

func move(s string, x int, y int, n int) (int, int) {
	switch s {
	case N:
		y += n
	case S:
		y -= n
	case E:
		x += n
	case W:
		x -= n
	}
	return x, y
}

func TestAOCDay12(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		}, fn: day12part1, expected: 25},
		{name: "part1", input: readStringsInput("day12"), fn: day12part1, expected: 364},
		{name: "part2sample", input: []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"L90",
			"R90",
			"F11",
		}, fn: day12part2, expected: 286},
		{name: "part1", input: readStringsInput("day12"), fn: day12part2, expected: 39518},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
