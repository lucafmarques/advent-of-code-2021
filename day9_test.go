package main

import "testing"

func BenchmarkDay9Part1(b *testing.B) {
	data := dataDay9()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day9part1(data)
	}
}

func BenchmarkDay9Part2(b *testing.B) {
	data := dataDay9()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day9part2(data)
	}
}
