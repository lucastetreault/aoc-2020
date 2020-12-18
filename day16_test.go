package main

import (
	"strconv"
	"strings"
	"testing"
)

type rng struct {
	name string
	high int64
	low  int64
}

type ticket struct {
	potentialFields map[int][]string
	values          []int64
}

func day16part2(strs []string) int64 {

	rngs := make([]rng, 0)
	tickets := make([]ticket, 0)
	fields := make([]string, 0)
	myTicket := ticket{}

	phase := 0

	for _, s := range strs {
		if len(s) == 0 {
			phase++
			continue
		}

		if phase == 0 {
			x := strings.Split(s, ":")
			nums := strings.Split(x[1], " or ")
			fields = append(fields, x[0])
			for _, n := range nums {
				y := strings.Split(n, "-")

				r := rng{name: x[0]}
				r.low, _ = strconv.ParseInt(strings.Trim(y[0], " "), 10, 64)
				r.high, _ = strconv.ParseInt(strings.Trim(y[1], " "), 10, 64)

				rngs = append(rngs, r)
			}
		}
		if phase == 1 {
			if s == "your ticket:" {
				continue
			}

			raw := strings.Split(s, ",")
			myTicket = ticket{
				potentialFields: make(map[int][]string),
				values:          make([]int64, len(raw)),
			}
			for i, n := range raw {
				num, _ := strconv.ParseInt(n, 10, 64)
				myTicket.values[i] = num
			}
		}
		if phase == 2 {
			if s == "nearby tickets:" {
				continue
			}
			validTicket := true
			raw := strings.Split(s, ",")
			t := ticket{
				potentialFields: make(map[int][]string),
				values:          make([]int64, len(raw)),
			}
			for i, n := range raw {
				num, _ := strconv.ParseInt(n, 10, 64)
				t.values[i] = num
				valid := false
				for _, r := range rngs {
					if num >= r.low && num <= r.high {
						valid = true
						if _, ok := t.potentialFields[i]; !ok {
							t.potentialFields[i] = make([]string, 0)
						}
						t.potentialFields[i] = append(t.potentialFields[i], r.name)
					}
				}
				if !valid {
					validTicket = false
					break
				}
			}
			if validTicket {
				tickets = append(tickets, t)
			}
		}
	}

	res := int64(1)

	for _, f := range fields {
		if !strings.HasPrefix(f, "departure") {
			continue
		}
		d := make(map[int]int)
		for _, t := range tickets {
			for k, v := range t.potentialFields {
				for _, s := range v {
					if s == f {
						d[k]++
					}
				}
			}
		}
		for k, v := range d {
			if v == len(tickets) {
				res *= myTicket.values[k]
			}
		}
	}

	return res
}

//func day16part2(strs []string) int64 {
//	return -1
//}

func TestAOCDay16(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int64
		expected int64
	}{
		//{name: "part1sample", input: []string{
		//	"class: 1-3 or 5-7",
		//	"row: 6-11 or 33-44",
		//	"seat: 13-40 or 45-50",
		//	"",
		//	"your ticket:",
		//	"7,1,14",
		//	"",
		//	"nearby tickets:",
		//	"7,3,47",
		//	"40,4,50",
		//	"55,2,20",
		//	"38,6,12",
		//}, fn: day16part1, expected: int64(71)},
		//{name: "part1", input: readStringsInput("day16"), fn: day16part1, expected: int64(-1)},
		{name: "part2sample", input: []string{
			"departure class: 0-1 or 4-19",
			"row: 0-5 or 8-19",
			"departure seat: 0-13 or 16-19",
			"",
			"your ticket:",
			"11,12,13",
			"",
			"nearby tickets:",
			"3,9,18",
			"15,1,5",
			"5,14,9",
		}, fn: day16part2, expected: int64(11 * 13)},
		//{name: "part2", input: readStringsInput("day16"), fn: day16part2, expected: 3219837697833},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
