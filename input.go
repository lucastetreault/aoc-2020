package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInt64Input(day string) []int64 {
	b, err := ioutil.ReadFile(day + "_input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	raw := strings.Split(string(b), "\n")
	i64arr := make([]int64, len(raw))
	for i, r := range raw {
		i64arr[i], _ = strconv.ParseInt(r, 10, 64)
	}
	return i64arr
}

func readStringsInput(day string) []string {
	b, err := ioutil.ReadFile(day + "_input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	return strings.Split(string(b), "\n")
}
