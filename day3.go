package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dataDay3 = func() (dataDay3 [][]string) {
	file, _ := os.Open("data/day3.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dataDay3 = append(dataDay3, strings.Split(scanner.Text(), ""))
	}

	return
}

func day3() {
	data := dataDay3()
	fmt.Printf("Part 1: %v\n", day3part1(data))
	fmt.Printf("Part 2: %v\n", day3part2(data))
}

func day3part1(data [][]string) int64 {
	var common []string
	var uncommon []string

	for i := 0; i < len(data[0]); i++ {
		line := verticalLine(data, i)
		count0 := strings.Count(line, "0")
		count1 := strings.Count(line, "1")

		if count0 > count1 {
			common = append(common, "0")
			uncommon = append(uncommon, "1")
		} else {
			common = append(common, "1")
			uncommon = append(uncommon, "0")
		}
	}

	gamma, _ := strconv.ParseInt(strings.Join(common, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(uncommon, ""), 2, 64)

	return gamma * epsilon
}

func day3part2(data [][]string) int64 {
	oxygenBits := parseReport(data, 0, func(c0, c1 int) string {
		if c0 <= c1 {
			return "1"
		}

		return "0"
	})
	co2Bits := parseReport(data, 0, func(c0, c1 int) string {
		if c0 <= c1 {
			return "0"
		}

		return "1"
	})

	oxygen, _ := strconv.ParseInt(strings.Join(oxygenBits, ""), 2, 64)
	co2, _ := strconv.ParseInt(strings.Join(co2Bits, ""), 2, 64)

	return oxygen * co2
}

func parseReport(data [][]string, iter int, fn func(int, int) string) []string {
	var searchSpace [][]string
	line := verticalLine(data, iter)

	target := fn(strings.Count(line, "0"), strings.Count(line, "1"))

	for i := 0; i < len(data); i++ {
		if data[i][iter] == target {
			searchSpace = append(searchSpace, data[i])
		}
	}

	if len(searchSpace) == 1 {
		return searchSpace[0]
	}

	return parseReport(searchSpace, iter+1, fn)
}

func verticalLine(data [][]string, pos int) string {
	var line []string

	for i := 0; i < len(data); i++ {
		line = append(line, data[i][pos])
	}

	return strings.Join(line, "")
}
