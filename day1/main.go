package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var data = func() (data []int) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		data = append(data, n)
	}

	return
}()

func main() {
	fmt.Printf("Part 1: %v\n", solve(data, 1))
	fmt.Printf("Part 2: %v\n", solve(data, 3))
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
