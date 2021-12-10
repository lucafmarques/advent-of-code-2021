package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var dataDay7 = func() (dataDay7 []int) {
	file, _ := os.Open("data/day7.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	for _, s := range strings.Split(scanner.Text(), ",") {
		n, _ := strconv.Atoi(s)
		dataDay7 = append(dataDay7, n)
	}

	return
}

func day7(kind string) {
	data := dataDay7()
	switch kind {
	case "brute-force":
		brute(data)
	default:
		maths(data)
	}
}

func brute(data []int) {
	fmt.Println("Brute Force Solution")
	pos, fuel := day7part1brute(data)
	fmt.Printf("Part 1: %v %v\n", pos, fuel)
	pos, fuel = day7part2brute(data)
	fmt.Printf("Part 2: %v %v\n", pos, fuel)
}

func maths(data []int) {
	fmt.Println("Maths Solution")
	pos, fuel := day7part1(data)
	fmt.Printf("Part 1: %v %v\n", pos, fuel)
	pos, fuel = day7part2(data)
	fmt.Printf("Part 2: %v %v\n", pos, fuel)
}

func day7part1brute(data []int) (int, int) {
	cost := func(n, m int) int {
		return abs(n - m)
	}

	positions := calculatePositions(data, cost)

	fn := func(n, m int) bool {
		return n < m
	}

	return find(positions, fn)
}

func day7part2brute(data []int) (int, int) {
	cost := func(n, m int) int {
		diff := abs(n - m)
		return (diff * (diff + 1)) / 2
	}

	positions := calculatePositions(data, cost)

	fn := func(n, m int) bool {
		return n < m
	}

	return find(positions, fn)
}

func calculatePositions(data []int, cost func(n, m int) int) map[int]int {
	var positions map[int]int = map[int]int{}

	for _, p := range data {
		var fuel int
		if _, ok := positions[p]; ok {
			continue
		}

		for _, c := range data {
			fuel += cost(p, c)
		}

		positions[p] = fuel
	}

	return positions
}

func find(posMap map[int]int, fn func(n, m int) bool) (int, int) {
	var pos, fuel int = 0, posMap[0]

	for k, v := range posMap {
		if fn(v, fuel) {
			pos = k
			fuel = v
		}
	}

	return pos, fuel
}

func day7part1(data []int) (int, int) {
	cost := func(n, m int) int {
		return abs(n - m)
	}

	sorted := data
	sort.Ints(sorted)

	median := len(sorted) / 2

	return data[median], calculateFuelCost(sorted, data[median], cost)
}

func day7part2(data []int) (int, int) {
	cost := func(n, m int) int {
		diff := abs(n - m)
		return (diff * (diff + 1)) / 2
	}

	mean := func() int {
		var total int
		for _, v := range data {
			total += v
		}

		return total / len(data)
	}()

	return mean, calculateFuelCost(data, mean, cost)
}

func calculateFuelCost(data []int, target int, cost func(int, int) int) int {
	var fuel int

	for _, v := range data {
		fuel += cost(target, v)
	}

	return fuel
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
