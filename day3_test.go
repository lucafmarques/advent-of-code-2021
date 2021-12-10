package main

import "testing"

func BenchmarkDay3Part1(b *testing.B) {
	data := dataDay3()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day3part1(data)
	}
}

func BenchmarkDay3Part2(b *testing.B) {
	data := dataDay3()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day3part1(data)
	}
}
