package main

import "testing"

func BenchmarkDay2Part1(b *testing.B) {
	data := dataDay2()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day2part1(data)
	}
}

func BenchmarkDay2Part2(b *testing.B) {
	data := dataDay2()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day2part1(data)
	}
}
