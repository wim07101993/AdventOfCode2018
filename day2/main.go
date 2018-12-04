package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	fmt.Println("Part 1:", solvePart1(args))
	fmt.Println("Part 2:", solvePart2(args))
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

func solvePart2(ids []string) string {
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if b, s := hasOnlyOneLetterDifferent(ids[i], ids[j]); b {
				return s
			}
		}
	}
	return ""
}

func hasOnlyOneLetterDifferent(id1 string, id2 string) (bool, string) {
	foundDiff := false
	equalLetters := ""
	for i := 0; i < len(id1); i++ {
		if id1[i] != id2[i] {
			if foundDiff {
				return false, ""
			} else {
				foundDiff = true

				builder := strings.Builder{}
				for j := 0; j < len(id1); j++ {
					if i != j {
						builder.WriteRune(rune(id1[j]))
					}
				}
				equalLetters = builder.String()
			}
		}
	}

	return foundDiff, equalLetters
}
