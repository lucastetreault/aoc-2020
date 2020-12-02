package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

type Passwords struct {
	min      int
	max      int
	letter   string
	password string
}

func readPasswordInput(day string) []Passwords {
	b, err := ioutil.ReadFile(day + "_input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	raw := strings.Split(string(b), "\n")
	in := make([]Passwords, len(raw))

	for i, r := range raw {
		strs := strings.Split(r, " ")
		nums := strings.Split(strs[0], "-")
		in[i].min, _ = strconv.Atoi(nums[0])
		in[i].max, _ = strconv.Atoi(nums[1])
		in[i].letter = strs[1][:1]
		in[i].password = strs[2]
	}
	return in
}

func CharacterCountPasswordPolicy(input []Passwords) int {
	validCount := 0
	for _, in := range input {
		count := 0
		for _, p := range in.password {
			if string(p) == in.letter {
				count++
			}
		}
		if count >= in.min && count <= in.max {
			validCount++
		}
	}
	return validCount
}

func XORPositionValuePasswordPolicy(input []Passwords) int {
	validCount := 0
	for _, in := range input {
		if (string(in.password[in.min-1]) == in.letter) != (string(in.password[in.max-1]) == in.letter) {
			validCount++
		}
	}
	return validCount
}

func TestAOCDay2(t *testing.T) {
	tests := []struct {
		name     string
		input    []Passwords
		fn       func([]Passwords) int
		expected int
	}{
		{name: "sample", input: []Passwords{
			{min: 1, max: 3, letter: "a", password: "abcde"},
			{min: 1, max: 3, letter: "b", password: "cdefg"},
			{min: 2, max: 9, letter: "c", password: "ccccccccc"},
		}, fn: CharacterCountPasswordPolicy, expected: 2},
		{name: "part1", input: readPasswordInput("day2"), fn: CharacterCountPasswordPolicy, expected: 467},
		{name: "part2", input: readPasswordInput("day2"), fn: XORPositionValuePasswordPolicy, expected: 441},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}

}
