package main

import "testing"

func BenchmarkDay8Part1(b *testing.B) {
	data := dataDay8()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day8part1(data)
	}
}

func BenchmarkDay8Part2(b *testing.B) {
	data := dataDay8()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day8part1(data)
	}
}
