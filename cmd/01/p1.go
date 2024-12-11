package main

import (
	"fmt"
	"os"

	"github.com/ghjm/advent_utils"
)

type data struct {
	values []int
}

func run() error {
	d := data{}
	err := utils.OpenAndReadLines("input1.txt", func(s string) error {
		d.values = append(d.values, utils.MustAtoi(s))
		return nil
	})
	if err != nil {
		return err
	}
	count := 0
	for i := 1; i < len(d.values); i++ {
		if d.values[i] > d.values[i-1] {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
	count = 0
	for i := 3; i < len(d.values); i++ {
		if d.values[i]+d.values[i-1]+d.values[i-2] > d.values[i-1]+d.values[i-2]+d.values[i-3] {
			count++
		}
	}
	fmt.Printf("Part 2: %d\n", count)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
