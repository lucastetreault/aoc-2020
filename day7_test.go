package main

import (
	"strconv"
	"strings"
	"testing"
)

type Bag struct {
	Colour      string
	ContainedIn map[string]int
	Contains    map[string]int
}

func Day7Part1(input []string) int {
	bags := parseInput(input)

	queue := bags["shiny gold"].ContainedIn

	options := make(map[string]int)
	for k, v := range queue {
		options[k] = v
	}

	count := 0
	for ; ; {
		if len(queue) == 0 {
			break
		}

		for next := range queue {
			count++
			for k, v := range bags[next].ContainedIn {
				if _, ok := options[k]; !ok {
					options[k] = v
					queue[k] = v
				}
			}
			delete(queue, next)
		}

	}

	return count
}

func Day7Part2(input []string) int {
	bags := parseInput(input)

	queue := make([]string, 0)

	for k, v := range bags["shiny gold"].Contains {
		for i := 0; i < v; i++ {
			queue = append(queue, k)
		}
	}

	count := 0
	for ; ; {
		if len(queue) == 0 {
			break
		}
		count++
		for k, v := range bags[queue[0]].Contains {
			for i := 0; i < v; i++ {
				queue = append(queue, k)
			}
		}
		queue = queue[1:]
	}

	return count
}

func parseInput(input []string) map[string]Bag {
	bags := make(map[string]Bag)

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
			bag.Contains[b.Colour] = count

			if _, ok := b.ContainedIn[bag.Colour]; ok {
				b.ContainedIn[bag.Colour] = b.ContainedIn[bag.Colour] + count
			} else {
				b.ContainedIn[bag.Colour] = b.ContainedIn[bag.Colour] + count
			}
		}
	}
	return bags
}

func getOrMakeBag(bags map[string]Bag, colour string) Bag {
	var bag Bag
	var ok bool
	if bag, ok = bags[colour]; !ok {
		bags[colour] = Bag{Colour: colour, ContainedIn: make(map[string]int), Contains: make(map[string]int)}
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
		{name: "part2", input: readStringsInput("day7"), fn: Day7Part2, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
