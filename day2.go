package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dataDay2 = func() (dataDay2 []string) {
	file, _ := os.Open("data/day2.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dataDay2 = append(dataDay2, scanner.Text())
	}

	return
}

type Position struct {
	Horizontal int
	Vertical   int
	Aim        int
}

func day2part1(data []string) int {
	pos := Position{}
	for _, v := range data {
		command := strings.Split(v, " ")
		value, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			pos.Horizontal += value
		case "down":
			pos.Vertical += value
		case "up":
			pos.Vertical -= value
		}
	}

	return pos.Vertical * pos.Horizontal
}

func day2part2(data []string) int {
	pos := Position{}
	for _, v := range data {
		command := strings.Split(v, " ")
		value, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			pos.Horizontal += value
			pos.Vertical += pos.Aim * value
		case "down":
			pos.Aim += value
		case "up":
			pos.Aim -= value
		}
	}

	return pos.Vertical * pos.Horizontal
}

func day2() {
	data := dataDay2()
	fmt.Printf("Part 1: %v\n", day2part1(data))
	fmt.Printf("Part 2: %v\n", day2part2(data))
}
