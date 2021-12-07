package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Coord struct {
	p1 Point
	p2 Point
}

type Map [999][999]int

func (m Map) String() string {
	var output string
	for _, row := range m {
		output = fmt.Sprintf("%v\n%v", output, row)
	}

	return output
}

var data = func() (data []Coord) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), " -> ")

		point1 := strings.Split(line[0], ",")
		p1x, _ := strconv.Atoi(point1[0])
		p1y, _ := strconv.Atoi(point1[1])

		point2 := strings.Split(line[1], ",")
		p2x, _ := strconv.Atoi(point2[0])
		p2y, _ := strconv.Atoi(point2[1])

		coord := Coord{
			Point{
				x: p1x,
				y: p1y,
			},
			Point{
				x: p2x,
				y: p2y,
			},
		}

		data = append(data, coord)
	}

	return
}()

func main() {
	fmt.Printf("Part 1: %v\n", part1(data))
	fmt.Printf("Part 2: %v\n", part2(data))
}

func part1(coord []Coord) int {
	var total int
	var table Map

	for _, c := range coord {
		if c.p1.x == c.p2.x {
			for i := range walk(c.p1.y, c.p2.y) {
				if table[i][c.p1.x]++; table[i][c.p1.x] == 2 {
					total += 1
				}
			}
		} else if c.p1.y == c.p2.y {
			for i := range walk(c.p1.x, c.p2.x) {
				if table[c.p1.y][i]++; table[c.p1.y][i] == 2 {
					total += 1
				}
			}
		}
	}

	return total
}

func part2(coord []Coord) int {
	var total int
	var table Map

	for _, c := range coord {
		if abs(c.p1.x-c.p2.x) == abs(c.p1.y-c.p2.y) {
			s, f := closestCenter(c.p1, c.p2)

			var dirx, diry int
			if s.y <= f.y {
				diry = 1
			} else {
				diry = -1
			}

			if s.x <= f.x {
				dirx = 1
			} else {
				dirx = -1
			}

			for i := 0; i <= abs(c.p1.y-c.p2.y); i++ {
				if table[s.y+(diry*i)][s.x+(dirx*i)]++; table[s.y+(diry*i)][s.x+(dirx*i)] == 2 {
					total += 1
				}
			}
		}

		if c.p1.x == c.p2.x {
			for i := range walk(c.p1.y, c.p2.y) {
				if table[i][c.p1.x]++; table[i][c.p1.x] == 2 {
					total += 1
				}
			}
		} else if c.p1.y == c.p2.y {
			for i := range walk(c.p1.x, c.p2.x) {
				if table[c.p1.y][i]++; table[c.p1.y][i] == 2 {
					total += 1
				}
			}
		}
	}

	return total
}

func closestCenter(p1 Point, p2 Point) (Point, Point) {
	d1 := math.Sqrt(float64(p1.x*p1.x + p1.y*p1.y))
	d2 := math.Sqrt(float64(p2.x*p2.x + p2.y*p2.y))

	if d1 < d2 {
		return p1, p2
	}

	return p2, p1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func walk(n, m int) <-chan int {
	ch := make(chan int)

	var source, dest int
	if n > m {
		source = m
		dest = n
	} else {
		source = n
		dest = m
	}

	go func() {
		for i := source; i <= dest; i++ {
			ch <- i
		}

		close(ch)
	}()

	return ch
}
