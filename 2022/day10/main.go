package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name  string
	param int
}

type CPU struct {
	register int
	cycle    int
	stack    []*Instruction
}

func main() {
	inputFileName := "input.txt"
	data, _ := os.ReadFile(inputFileName)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	cpu := CPU{register: 1, cycle: 1, stack: []*Instruction{}}

	for _, line := range lines {
		if line == "noop" {
			cpu.stack = append(cpu.stack, &Instruction{name: "noop"})
		}
		if strings.HasPrefix(line, "addx") {
			splat := strings.Split(line, " ")
			amount, _ := strconv.Atoi(splat[1])
			cpu.stack = append(cpu.stack, &Instruction{name: "sleep"})
			cpu.stack = append(cpu.stack, &Instruction{name: "addx", param: amount})
		}
	}

	answer1 := 0
	history := []int{}
	for _, instruction := range cpu.stack {
		if cpu.cycle == 20 || cpu.cycle == 60 || cpu.cycle == 100 || cpu.cycle == 140 || cpu.cycle == 180 || cpu.cycle == 220 {
			history = append(history, cpu.register)
			answer1 += cpu.register * cpu.cycle
		}

		cpu.process(*instruction)
		cpu.cycle++
	}

	fmt.Println(history)
	fmt.Println("answer1", answer1) //15680
}

func (c *CPU) process(instruction Instruction) {
	if instruction.name == "addx" {
		c.register += instruction.param
	}
}
