package main

import "testing"

func BenchmarkDay5Part1(b *testing.B) {
	data := dataDay5()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day5part1(data)
	}
}

func BenchmarkDay5Part2(b *testing.B) {
	data := dataDay5()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day5part1(data)
	}
}
