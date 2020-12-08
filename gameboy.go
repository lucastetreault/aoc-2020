package main

import "errors"

type Gameboy struct {
	instructions []Instruction
	accumulator  int
	idx          int
}

func Bootloader(instructions []Instruction) *Gameboy {
	return &Gameboy{instructions: instructions}
}

func (g *Gameboy) Boot(config ...BootConfig) (int, error) {
	detectBootLoop := false
	repairBootLoop := false

	for _, c := range config {
		switch c {
		case DetectBootLoop:
			detectBootLoop = true
		case RepairBootLoop:
			repairBootLoop = true
		}
	}

	for ; g.idx < len(g.instructions); {
		currentInstruction := g.instructions[g.idx]

		if detectBootLoop && currentInstruction.HasBeenExecuted() {
			return g.accumulator, BootLoopDetectedError
		}
		if repairBootLoop {
			if acc, err := CloneForRepair(g).Boot(DetectBootLoop); err == nil {
				return acc, nil
			}
		}
		currentInstruction.Execute(g)
	}
	return g.accumulator, nil
}

func CloneForRepair(g *Gameboy) Bootable {
	clone := &Gameboy{
		instructions: make([]Instruction, len(g.instructions)),
		accumulator:  g.accumulator,
		idx:          g.idx,
	}
	for idx, instruction := range g.instructions {
		clone.instructions[idx] = instruction.Clone()
	}

	var err error
	if clone.instructions[clone.idx], err = g.instructions[g.idx].Repair(); err != nil {
		return &Unrepairable{}
	}
	return clone
}

var BootLoopDetectedError = errors.New("boot loop detected")
var UnrepairableError = errors.New("unrepairable")

type BootConfig int

const (
	DetectBootLoop = iota
	RepairBootLoop
)

type Bootable interface {
	Boot(config ...BootConfig) (int, error)
}

type Unrepairable struct {
}

func (u *Unrepairable) Boot(_ ...BootConfig) (int, error) {
	return -1, UnrepairableError
}

type Instruction interface {
	Execute(gameboy *Gameboy)
	HasBeenExecuted() bool
	Clone() Instruction
	Repair() (Instruction, error)
}

type AccInstruction struct {
	executed bool
	Argument int
}

func (a *AccInstruction) Execute(g *Gameboy) {
	g.accumulator += a.Argument
	g.idx++
	a.executed = true
	return
}

func (a AccInstruction) HasBeenExecuted() bool {
	return a.executed
}

func (a *AccInstruction) Clone() Instruction {
	return &AccInstruction{
		executed: a.executed,
		Argument: a.Argument,
	}
}

func (a *AccInstruction) Repair() (Instruction, error) {
	return nil, UnrepairableError
}

type JmpInstruction struct {
	executed bool
	Argument int
}

func (a *JmpInstruction) Execute(g *Gameboy) {
	g.idx += a.Argument
	a.executed = true
	return
}

func (a JmpInstruction) HasBeenExecuted() bool {
	return a.executed
}

func (a *JmpInstruction) Clone() Instruction {
	return &JmpInstruction{
		executed: a.executed,
		Argument: a.Argument,
	}
}

func (a *JmpInstruction) Repair() (Instruction, error) {
	return &NopInstruction{
		executed: a.executed,
	}, nil
}

type NopInstruction struct {
	Argument int
	executed bool
}

func (a *NopInstruction) Execute(g *Gameboy) {
	a.executed = true
	g.idx++
	return
}

func (a NopInstruction) HasBeenExecuted() bool {
	return a.executed
}

func (a *NopInstruction) Clone() Instruction {
	return &NopInstruction{
		Argument: a.Argument,
		executed: a.executed,
	}
}

func (a *NopInstruction) Repair() (Instruction, error) {
	return &JmpInstruction{
		Argument: a.Argument,
		executed: a.executed,
	}, nil
}
