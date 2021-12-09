package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data = func() (data [][]int) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var lineData []int
		line := strings.Split(scanner.Text(), "")
		for _, c := range line {
			n, _ := strconv.Atoi(c)
			lineData = append(lineData, n)
		}
		data = append(data, lineData)
	}

	return
}()

type Coord struct {
	y, x int
}

func main() {
	fmt.Printf("Part 1: %v\n", part1(data))
	fmt.Printf("Part 1: %v\n", part2(data))
}

func part1(data [][]int) int {
	var total int
	lengthy := len(data)
	lengthx := len(data[0])

	for y := range data {
		for x := range data[y] {
			if lookaround(data, x, y, lengthx, lengthy) {
				total += data[y][x] + 1
			}
		}
	}

	return total
}

func part2(data [][]int) int {
	var total int = 1
	var basins []int
	lengthy := len(data)
	lengthx := len(data[0])

	for y := range data {
		for x := range data[y] {
			if lookaround(data, x, y, lengthx, lengthy) {
				size := findbasin(data, x, y)
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	for _, v := range basins[len(basins)-3:] {
		total *= v
	}

	return total
}

func findbasin(data [][]int, x, y int) int {
	var f func(int, int, int, rune) int
	walked := map[Coord]bool{}

	start, ly, lx := data[y][x], len(data), len(data[0])

	f = func(x, y, v int, dir rune) int {
		coord := Coord{y, x}

		if x < 0 || y < 0 || x == lx || y == ly {
			return 0
		}

		curr := data[y][x]
		flow := curr-v > 0

		if _, ok := walked[coord]; ok || !flow || curr == 9 {
			return 0
		}

		walked[coord] = true

		switch dir {
		case 'u':
			return 1 + f(x, y-1, curr, 'u') + f(x-1, y, curr, 'l') + f(x+1, y, curr, 'r')
		case 'd':
			return 1 + f(x, y+1, curr, 'd') + f(x-1, y, curr, 'l') + f(x+1, y, curr, 'r')
		case 'l':
			return 1 + f(x-1, y, curr, 'l') + f(x, y-1, curr, 'u') + f(x, y+1, curr, 'd')
		case 'r':
			return 1 + f(x+1, y, curr, 'r') + f(x, y-1, curr, 'u') + f(x, y+1, curr, 'd')
		}

		return 0
	}

	return 1 + f(x, y-1, start, 'u') + f(x, y+1, start, 'd') + f(x-1, y, start, 'l') + f(x+1, y, start, 'r')
}

func lookaround(data [][]int, x, y, lx, ly int) bool {
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
