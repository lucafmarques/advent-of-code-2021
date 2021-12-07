package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var data = func() (data []int) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	for _, s := range strings.Split(scanner.Text(), ",") {
		n, _ := strconv.Atoi(s)
		data = append(data, n)
	}

	return
}()

type Simulation struct {
	counts *[9]int
	period int
}

func main() {
	fmt.Printf("Part 1: %v\n", part1(data))
	fmt.Printf("Part 2: %v\n", part2(data))
}

func part1(data []int) int {
	simulation := NewSimulation(data, 80)
	return simulation.simulate()
}

func part2(data []int) int {
	simulation := NewSimulation(data, 256)
	return simulation.simulate()
}

func (s *Simulation) sum() (total int) {
	for i := range s.counts {
		total += s.counts[i]
	}

	return
}

func (s *Simulation) simulate() int {
	for day := 1; day <= s.period; day++ {
		var iteration [9]int
		var reproduced int

		for pos := range s.counts {
			if pos == 0 {
				reproduced = s.counts[0]
				continue
			}

			iteration[pos-1] = s.counts[pos]
		}

		iteration[6] += reproduced
		iteration[8] = reproduced

		s.counts = &iteration
	}

	return s.sum()
}

func NewSimulation(data []int, days int) Simulation {
	s := Simulation{
		counts: &[9]int{},
		period: days,
	}

	for i := range data {
		s.counts[data[i]] += 1
	}

	return s
}
