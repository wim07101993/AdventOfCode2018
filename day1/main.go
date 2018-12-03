package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	sum := 0

	for _, a := range args {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		sum += i
	}

	fmt.Println(sum)
}
