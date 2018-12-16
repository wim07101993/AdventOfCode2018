package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := readFile()
	fmt.Println("Part1:", solvePart1(input))
}

func readFile() string {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func solvePart1(input string) int {
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] &&
			unicode.ToUpper(rune(input[i])) == unicode.ToUpper(rune(input[i-1])) {
			input = input[:i-1] + input[i+1:]
			i -= 2
			if i < 0 {
				i = 0
			}
		}
	}
	return len(input)
}
