package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	is := make([]int, 0)
	for _, a := range args {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}

	fmt.Println("Part 1:", solvePart1(is))
	fmt.Println("Part 2:", solvePart2(is))
}

func solvePart1(is []int) (sum int) {
	for _, i := range is {
		sum += i
	}
	return sum
}

func solvePart2(is []int) (duplicate int) {
	sums := make(map[int]struct{})
	sums[0] = struct{}{}
	sum := 0

	for _, i := range is {
		sum += i
		sums[sum] = struct{}{}
	}

	for {
		for _, i := range is {
			sum += i

			if _, ok := sums[sum]; ok {
				return sum
			} else {
				sums[sum] = struct{}{}
			}
		}
	}
}
