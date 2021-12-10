package main

import "testing"

func BenchmarkDay7Part1brute(b *testing.B) {
	data := dataDay7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7part1brute(data)
	}
}

func BenchmarkDay7Part2brute(b *testing.B) {
	data := dataDay7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7part2brute(data)
	}
}

func BenchmarkDay7Part1math(b *testing.B) {
	data := dataDay7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7part1(data)
	}
}

func BenchmarkDay7Part2math(b *testing.B) {
	data := dataDay7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7part2(data)
	}
}
