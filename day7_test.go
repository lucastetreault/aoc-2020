package main

import (
	"strconv"
	"strings"
	"testing"
)

type Bag struct {
	Colour      string
	ContainedIn []*Bag
	Contains    []*Bag
}

func Day7Part1(input []string) int {
	bags := parseInput(input)

	queue := bags["shiny gold"].ContainedIn

	options := make(map[string]int)

	for i := 0; i < len(queue); i++ {
		options[queue[i].Colour] = 1
		queue = append(queue, queue[i].ContainedIn...)
	}

	return len(options)
}

func Day7Part2(input []string) int {
	bags := parseInput(input)

	queue := bags["shiny gold"].Contains

	for i := 0; i < len(queue); i++ {
		queue = append(queue, queue[i].Contains...)
	}

	return len(queue)
}

func parseInput(input []string) map[string]*Bag {
	bags := make(map[string]*Bag)

	for _, i := range input {
		i = strings.Replace(i, "bags", "", -1)
		i = strings.Replace(i, "bag", "", -1)
		i = strings.Replace(i, ".", "", -1)
		arr := strings.Split(i, "contain")
		colour := strings.Trim(arr[0], " ")
		bag := getOrMakeBag(bags, colour)
		contains := strings.Split(arr[1], ",")
		for _, c := range contains {
			b := getOrMakeBag(bags, strings.Trim(c, " ")[2:])
			count, _ := strconv.Atoi(strings.Trim(c, " ")[:1])
			for i := 0; i < count; i++ {
				bag.Contains = append(bag.Contains, b)
			}

			for i := 0; i < count; i++ {
				b.ContainedIn = append(b.ContainedIn, bag)
			}
		}
	}
	return bags
}

func getOrMakeBag(bags map[string]*Bag, colour string) *Bag {
	var bag *Bag
	var ok bool
	if bag, ok = bags[colour]; !ok {
		bags[colour] = &Bag{Colour: colour, ContainedIn: make([]*Bag, 0), Contains: make([]*Bag, 0)}
		bag = bags[colour]
	}
	return bag
}

func TestAOCDay7(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		}, fn: Day7Part1, expected: 4},
		{name: "part1", input: readStringsInput("day7"), fn: Day7Part1, expected: 355},
		{name: "part2sample", input: []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		}, fn: Day7Part2, expected: 32},
		{name: "part2sample", input: []string{
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		}, fn: Day7Part2, expected: 126},
		{name: "part2", input: readStringsInput("day7"), fn: Day7Part2, expected: 5312},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
