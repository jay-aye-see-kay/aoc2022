package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1: %+v\n", part1(readInput()))
	fmt.Printf("part 2: %+v\n", part2(readInput()))
}

func readInput() string {
	content, err := ioutil.ReadFile("../inputs/day07.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

type StrSet map[string]bool

func (s *StrSet) push(str string) {
	(*s)[str] = true
}

func parseCommands(input string) (StrSet, map[string]int) {
	currentDirFull := "/"
	dirSet := StrSet{}           // list of dirs
  dirSet.push("/")
	fileDict := map[string]int{} // name -> size

	for _, line := range strings.Split(input, "\n") {
		if line[:4] == "$ cd" {
			currentDirShort := line[5:]
			if currentDirShort == ".." {
				currentDirFull = pathUp(currentDirFull)
			} else if currentDirShort == "/" {
				currentDirFull = "/"
			} else {
				currentDirFull += currentDirShort + "/"
			}

		} else if line[:4] == "$ ls" {
			// ignore these lines, we can infer when it was run when we start getting results
		} else if line[:3] == "dir" {
			dirName := strings.Split(line, " ")[1]
			dirSet.push(currentDirFull + dirName + "/")
		} else {
			// we've found a file
			s := strings.Split(line, " ")
			name := s[1]
			size, _ := strconv.Atoi(s[0])
			fileDict[currentDirFull+name] = size
		}
	}
	return dirSet, fileDict
}

func pathUp(path string) string {
	// remove the trailing slash
	path = path[:len(path)-1]
	// find the last slash index
	idx := strings.LastIndex(path, "/")
	return path[:idx+1]
}

func part1(input string) int {
	dirSet, fileDict := parseCommands(input)

	sizes := make(map[string]int)
	for dirPath := range dirSet {
		for filePath, size := range fileDict {
			if len(dirPath) < len(filePath) && filePath[:len(dirPath)] == dirPath {
				sizes[dirPath] += size
			}
		}
	}

	total := 0
	for _, size := range sizes {
		if size <= 100000 {
			total += size
		}
	}

	return total
}

func part2(input string) int {
	dirSet, fileDict := parseCommands(input)

	sizes := make(map[string]int)
	for dirPath := range dirSet {
		for filePath, size := range fileDict {
			if len(dirPath) < len(filePath) && filePath[:len(dirPath)] == dirPath {
				sizes[dirPath] += size
			}
		}
	}

  totalUsage := sizes["/"]
  totalAvailable := 70000000
  requiredSpace := 30000000
  availableSpace := totalAvailable - totalUsage
  requiredToFree := requiredSpace - availableSpace

  toFree := totalUsage
  for _, dirSize := range sizes {
    if dirSize > requiredToFree && dirSize < toFree {
      toFree = dirSize
    }
  }

	return toFree
}
