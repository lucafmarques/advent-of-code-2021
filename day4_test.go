package main

import "testing"

func BenchmarkDay4Part1(b *testing.B) {
	data := dataDay4()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day4part1(data)
	}
}

func BenchmarkDay4Part2(b *testing.B) {
	data := dataDay4()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day4part1(data)
	}
}
