package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInstructionsInput(day string) []Instruction {
	b, err := ioutil.ReadFile(day + "_input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	lines := strings.Split(string(b), "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		arr := strings.Split(line, " ")

		arg, _ := strconv.Atoi(arr[1])

		switch arr[0] {
		case "acc":
			instructions[i] = &AccInstruction{
				Argument: arg,
			}
			break
		case "jmp":
			instructions[i] = &JmpInstruction{
				Argument: arg,
			}
			break
		case "nop":
			instructions[i] = &NopInstruction{}
			break
		}
	}

	return instructions
}
