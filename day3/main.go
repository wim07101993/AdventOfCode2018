package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id       int
	x        int
	y        int
	width    int
	height   int
	overlaps bool
}

func main() {
	cs := toClaims(os.Args[1:])

	fmt.Println("Part 1:", solvePart1(cs))
	fmt.Println("Part 2:", solvePart2(cs))
}

func toClaims(strs []string) []claim {
	cs := make([]claim, 0)

	for i := 0; i < len(strs); i += 4 {
		id, _ := strconv.Atoi(strs[i][1:])

		sc := strings.Split(strs[i+2][:len(strs[i+2])-1], ",")
		x, _ := strconv.Atoi(sc[0])
		y, _ := strconv.Atoi(sc[1])

		ss := strings.Split(strs[i+3], "x")
		w, _ := strconv.Atoi(ss[0])
		h, _ := strconv.Atoi(ss[1])

		cs = append(cs, claim{
			id:     id,
			x:      x,
			y:      y,
			width:  w,
			height: h,
		})
	}

	return cs
}

func solvePart1(cs []claim) int {
	arr := [2000][2000]int{}
	count := 0

	for _, c := range cs {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if arr[i][j] == 2 {
					continue
				} else if arr[i][j] == 1 {
					count++
					arr[i][j] = 2
				} else {
					arr[i][j] = 1
				}
			}
		}
	}

	return count
}

func solvePart2(cs []claim) int {
	arr := [2000][2000]int{}
	count := 0

	for _, c := range cs {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if arr[i][j] == 2 {
					continue
				} else if arr[i][j] == 1 {
					count++
					arr[i][j] = 2
				} else {
					arr[i][j] = 1
				}
			}
		}
	}

ClaimsLoop:
	for _, c := range cs {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if arr[i][j] != 1 {
					continue ClaimsLoop
				}
			}
		}

		return c.id
	}
	return -1
}
