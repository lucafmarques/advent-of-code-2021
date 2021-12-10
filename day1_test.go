package main

import "testing"

func BenchmarkDay1Part1(b *testing.B) {
	data := dataDay1()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day1part1(data)
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	data := dataDay1()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day1part1(data)
	}
}
