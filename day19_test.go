package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

type rule struct {
	num     int
	raw     string
	options []string
}

func day19part1(strs []string) int64 {
	phase := 0
	rules := make([]*rule, 0)
	inputs := make([]string, 0)
	for _, s := range strs {
		if s == "" {
			phase++
			continue
		}

		if phase == 0 {
			x := strings.Split(s, ":")
			r := rule{raw: strings.Trim(x[1], " ")}
			r.num, _ = strconv.Atoi(x[0])
			rules = append(rules, &r)
		} else {
			inputs = append(inputs, s)
		}
	}

	// sort descending
	sort.Slice(rules, func(i, j int) bool { return rules[j].num < rules[i].num })

	for _, r := range rules {
		if strings.Contains(r.raw, "\"") {
			r.options = []string{strings.Trim(strings.Trim(r.raw, " "), "\"")}
		} else {
			opts := strings.Split(r.raw, "|")
			r.options = make([]string, 0)
			for _, o := range opts {
				s := []string{""}
				o = strings.Trim(o, " ")
				for _, x := range strings.Split(o, " ") {
					i, _ := strconv.Atoi(x)
					for j, t := range s {
						for i, ref := range rules[len(rules)-1-i].options {
							if i == 0 {
								s[j] += ref
							} else {
								s = append(s, t+ref)
							}
						}
					}
				}
				r.options = append(r.options, s...)
			}
		}
	}

	// sort ascending
	sort.Slice(rules, func(i, j int) bool { return rules[i].num < rules[j].num })

	count := 0
	for _, input := range inputs {
		for _, opt := range rules[0].options {
			if input == opt {
				count++
				continue
			}
		}
	}

	return int64(count)
}

func day19part2(strs []string) int64 {
	return -1
}

func TestAOCDay19(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		{name: "part1sample", input: []string{
			"0: 4 1 5",
			"1: 2 3 | 3 2",
			"2: 4 4 | 5 5",
			"3: 4 5 | 5 4",
			"4: \"a\"",
			"5: \"b\"",
			"",
			"ababbb",
			"bababa",
			"abbbab",
			"aaabbb",
			"aaaabbb",
		}, fn: day19part1, expected: int64(2)},
		{name: "part1", input: readStringsInput("day19"), fn: day19part1, expected: int64(-1)},
		//{name: "part2sample", input: []string{
		//	"1 + 2 * 3 + 4 * 5 + 6",
		//}, fn: day19part2, expected: int64(231)},
		//{name: "part2", input: readStringsInput("day19"), fn: day19part2, expected: 158193007916215},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
