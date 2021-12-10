package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var dataDay9 = func() (dataDay9 [][]int) {
	file, _ := os.Open("data/day9.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var lineData []int
		line := strings.Split(scanner.Text(), "")
		for _, c := range line {
			n, _ := strconv.Atoi(c)
			lineData = append(lineData, n)
		}
		dataDay9 = append(dataDay9, lineData)
	}

	return
}

func day9() {
	data := dataDay9()
	fmt.Printf("Part 1: %v\n", day9part1(data))
	fmt.Printf("Part 1: %v\n", day9part2(data))
}

func day9part1(data [][]int) int {
	var total int
	lengthy, lengthx := len(data), len(data[0])

	for y := range data {
		for x := range data[y] {
			if findvent(data, x, y, lengthx, lengthy) {
				total += data[y][x] + 1
			}
		}
	}

	return total
}

func day9part2(data [][]int) int {
	var basins []int
	lengthy, lengthx := len(data), len(data[0])

	for y := range data {
		for x := range data[y] {
			if findvent(data, x, y, lengthx, lengthy) {
				size := findbasinsize(data, x, y)
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	count := len(basins)

	return basins[count-1] * basins[count-2] * basins[count-3]
}

func findbasinsize(data [][]int, x, y int) int {
	var f func(int, int, int) int

	start, ly, lx := data[y][x], len(data), len(data[0])

	f = func(x, y, v int) int {
		if x < 0 || y < 0 || x == lx || y == ly {
			return 0
		}

		curr := data[y][x]
		flow := curr-v > 0

		if !flow || curr >= 9 {
			return 0
		}

		data[y][x] = 10
		return 1 + f(x, y-1, curr) + f(x, y+1, start) + f(x-1, y, curr) + f(x+1, y, curr)
	}

	return 1 + f(x, y-1, start) + f(x, y+1, start) + f(x-1, y, start) + f(x+1, y, start)
}

func findvent(data [][]int, x, y, lx, ly int) bool {
	var min bool = true

	if y == 0 {
		min = min && (data[y][x] < data[y+1][x])
	} else if y == ly-1 {
		min = min && (data[y][x] < data[y-1][x])
	} else {
		min = min && (data[y][x] < data[y-1][x])
		min = min && (data[y][x] < data[y+1][x])
	}

	if !min {
		return false
	}

	if x == 0 {
		min = min && (data[y][x] < data[y][x+1])
	} else if x == lx-1 {
		min = min && (data[y][x] < data[y][x-1])
	} else {
		min = min && (data[y][x] < data[y][x-1])
		min = min && (data[y][x] < data[y][x+1])
	}

	return min
}
