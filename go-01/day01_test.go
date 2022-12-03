package main

import "testing"

func TestDay1(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		actual := part1()
		expected := 71934
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		actual := part2()
		expected := 211447
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})
}

func BenchmarkDay1(b *testing.B) {
	b.Run("part 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1()
		}
	})

	b.Run("part 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2()
		}
	})
}
