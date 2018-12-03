package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println("Part 1:", solvePart1(args))
}

func solvePart1(ids []string) int {
	doubles := 0
	triples := 0

	for _, id := range ids {
		double := false
		triple := false

		letters := make(map[rune]int)
		for _, r := range id {
			letters[r]++
		}

		for _, n := range letters {
			if !triple && n == 3 {
				triple = true
				triples++
			} else if !double && n == 2 {
				double = true
				doubles++
			}

			if triple && double {
				break
			}
		}
	}

	return doubles * triples
}
