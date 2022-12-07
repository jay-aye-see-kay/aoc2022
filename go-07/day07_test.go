package main

import "testing"

func TestDay1(t *testing.T) {
	t.Run("pathUp", func(t *testing.T) {
		actual := pathUp("/foo/bar/")
		expected := "/foo/"
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 1 sample", func(t *testing.T) {
		actual := part1(sample)
		expected := 95437
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 1", func(t *testing.T) {
		actual := part1(readInput())
		expected := 1077191
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		actual := part2(readInput())
		expected := 5649896
		if actual != expected {
			t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
		}
	})
}

var sample = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
