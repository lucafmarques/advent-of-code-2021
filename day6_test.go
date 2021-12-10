package main

import "testing"

func BenchmarkDay6Part1(b *testing.B) {
	data := dataDay6()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day6part1(data)
	}
}

func BenchmarkDay6Part2(b *testing.B) {
	data := dataDay6()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day6part1(data)
	}
}
