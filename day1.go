package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dataDay1 = func() (dataDay1 []int) {
	file, _ := os.Open("data/day1.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		dataDay1 = append(dataDay1, n)
	}

	return
}

func day1() {
	data := dataDay1()
	fmt.Printf("Part 1: %v\n", day1part1(data))
	fmt.Printf("Part 2: %v\n", day1part2(data))
}

func day1part1(data []int) int {
	return solve(data, 1)
}

func day1part2(data []int) int {
	return solve(data, 3)
}

func solve(data []int, gap int) int {
	var total int

	size := len(data)
	for i := range data {
		if i+gap+1 > size {
			gap = size - i
		}
		slice1 := data[i : i+gap]
		slice2 := data[i+1 : i+gap+1]

		sum1 := sumNumbers(slice1...)
		sum2 := sumNumbers(slice2...)

		if (sum2 - sum1) > 0 {
			total += 1
		}
	}

	return total
}

func sumNumbers(n ...int) (sum int) {
	for _, v := range n {
		sum += v
	}

	return sum
}
