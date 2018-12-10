package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := readFile()
	guards := parseInput(input)

	x := solvePart1(guards)
	fmt.Println("Part1:", x)
}

func readFile() []string {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)
	return lines
}

func parseInput(input []string) guards {
	const timeFormat = "2006-01-02 15:04"
	dateRegex := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}")
	idRegex := regexp.MustCompile("#\\d+")

	guards := guards{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		sMoment := dateRegex.FindString(line)
		moment, _ := time.Parse(timeFormat, sMoment)

		if strings.Contains(line, "wakes up") {
			guards.last().
				shifts.last().
				sleepPeriods.last().
				end = moment
		} else if strings.Contains(line, "falls asleep") {
			lastShift := guards.last().
				shifts.last()

			lastShift.sleepPeriods = append(lastShift.sleepPeriods, sleepPeriod{start: moment})
		} else {
			sId := idRegex.FindString(line)
			id, _ := strconv.Atoi(sId[1:])

			if i > 0 {
				guards.last().
					shifts.last().
					end = moment
			}

			guard := guard{
				id: id,
				shifts: shifts{
					shift{
						start:        moment,
						sleepPeriods: sleepPeriods{},
					},
				},
			}

			guards = append(guards, guard)
		}
	}

	return guards
}

func solvePart1(guards []guard) int {
	return -1
}
