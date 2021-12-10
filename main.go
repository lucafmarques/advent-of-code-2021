package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Add day number as first args.")
		os.Exit(1)
	}

	switch args[0] {
	case "1":
		day1()
	case "2":
		day2()
	case "3":
		day3()
	case "4":
		day4()
	case "5":
		day5()
	case "6":
		day6()
	case "7":
		day7("math")
	case "8":
		day8()
	case "9":
		day9()
	default:
		fmt.Println("NOT IMPLEMENTED YET")
	}
}
