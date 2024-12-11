package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghjm/advent_utils"
)

type instruction struct {
	direction string
	distance  int
}

type data struct {
	instructions []instruction
}

func run() error {
	d := data{}
	err := utils.OpenAndReadLines("input2.txt", func(s string) error {
		v := strings.Split(s, " ")
		if len(v) != 2 {
			return fmt.Errorf("invalid input")
		}
		d.instructions = append(d.instructions, instruction{
			direction: v[0],
			distance:  utils.MustAtoi(v[1]),
		})
		return nil
	})
	if err != nil {
		return err
	}
	pX := 0
	depth := 0
	for _, instr := range d.instructions {
		switch instr.direction {
		case "forward":
			pX += instr.distance
		case "down":
			depth += instr.distance
		case "up":
			depth -= instr.distance
		}
	}
	fmt.Printf("Part 1: %d\n", pX*depth)
	pX = 0
	depth = 0
	aim := 0
	for _, instr := range d.instructions {
		switch instr.direction {
		case "forward":
			pX += instr.distance
			depth += aim * instr.distance
		case "down":
			aim += instr.distance
		case "up":
			aim -= instr.distance
		}
	}
	fmt.Printf("Part 1: %d\n", pX*depth)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
