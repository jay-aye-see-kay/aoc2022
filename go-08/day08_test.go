package main

import "testing"

func TestDay1(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		actual := parse(smallSample)
		expected := Grid{[]int{3, 0, 3}, []int{2, 5, 5}, []int{6, 5, 3}}
		if actual[0][0] != expected[0][0] || actual[2][2] != expected[2][2] {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("is visible edge", func(t *testing.T) {
		actual := parse(smallSample).isVisible(0, 0)
		expected := true
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("is visible center", func(t *testing.T) {
		actual := parse(smallSample).isVisible(1, 1)
		expected := true
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 1", func(t *testing.T) {
		actual := part1()
		expected := 0
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		actual := part2()
		expected := 0
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})
}

var smallSample = `303
255
653`

var sample = `30373
25512
65332
33549
35390`
