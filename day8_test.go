package main

import (
	"testing"
)

func Day8Part1(instructions []Instruction) int {
	g := Bootloader(instructions)
	acc, _ := g.Boot(DetectBootLoop)
	return acc
}

func Day8Part2(instructions []Instruction) int {
	g := Bootloader(instructions)
	acc, _ := g.Boot(RepairBootLoop)
	return acc
}

func TestAOCDay8(t *testing.T) {
	tests := []struct {
		name     string
		input    []Instruction
		fn       func([]Instruction) int
		expected int
	}{
		{name: "part1sample", input: []Instruction{
			&NopInstruction{Argument: 0},
			&AccInstruction{Argument: +1},
			&JmpInstruction{Argument: +4},
			&AccInstruction{Argument: +3},
			&JmpInstruction{Argument: -3},
			&AccInstruction{Argument: -99},
			&AccInstruction{Argument: +1},
			&JmpInstruction{Argument: -4},
			&AccInstruction{Argument: +6},
		}, fn: Day8Part1, expected: 5},
		{name: "part1", input: readInstructionsInput("day8"), fn: Day8Part1, expected: 1928},
		{name: "part2sample", input: []Instruction{
			&NopInstruction{Argument: 0},
			&AccInstruction{Argument: +1},
			&JmpInstruction{Argument: +4},
			&AccInstruction{Argument: +3},
			&JmpInstruction{Argument: -3},
			&AccInstruction{Argument: -99},
			&AccInstruction{Argument: +1},
			&JmpInstruction{Argument: -4},
			&AccInstruction{Argument: +6},
		}, fn: Day8Part2, expected: 8},
		{name: "part2", input: readInstructionsInput("day8"), fn: Day8Part2, expected: 1319},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
