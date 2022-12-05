package main

import "testing"

func TestDay1(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		actual := part1()
		expected := "FCVRLMVQP"
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		actual := part2()
		expected := "RWLWGJGFD"
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})
}
