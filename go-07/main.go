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

// need to parse input into a list of commands
// then parse those into a tree
// then walk the tree summing file sizes
// - what kind of DS is easiest to walk and sum?
// - a recursive walk + sum is probably best

type FileNode struct {
	name string
	kind string // "dir" | "file"

	// dir only
	nodes []FileNode

	// file only
	size int
}

// dir name -> size
type DirSizes map[string]int

type FileDetails struct {
	name string
	size int
}

type DirList map[string][]FileDetails

func parseCommands(input string) DirList {

	currentDirShort := "/"
	currentDirFull := "/"

	// hold list of current file detail until we change
	tmpFileDetails := []FileDetails{}

	// list of dirs, and what's in them
	dirs := make(DirList)

	for _, line := range strings.Split(input, "\n") {
		if line[:4] == "$ cd" {
			// record previous dir info
			if len(tmpFileDetails) > 0 {
				dirs[currentDirFull] = tmpFileDetails
			}

			// reset previous dir info
			currentDirShort = line[5:]
			tmpFileDetails = []FileDetails{}

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
			// ignore dirs we don't go into?
		} else {
			// we've found a file
			s := strings.Split(line, " ")
			name := s[1]
			size, err := strconv.Atoi(s[0])
			if err != nil {
				panic("couldn't parse file size")
			}
			tmpFileDetails = append(tmpFileDetails, FileDetails{name, size})
		}
	}

	// add any remaining tmp data into dirs
	if len(tmpFileDetails) > 0 {
		dirs[currentDirFull] = tmpFileDetails
	}

	fmt.Println("")
	return dirs
}

func findChildDirs(parentDir string, dirList DirList) DirList {
	filteredList := make(DirList)
	for path, data := range dirList {
		if len(parentDir) < len(path) && path[:len(parentDir)] == parentDir {
			filteredList[path] = data
		}
	}
	return filteredList
}

func sumDirList(dirList DirList) int {
	total := 0
	for _, dirData := range dirList {
		for _, fileData := range dirData {
			total += fileData.size
		}
	}
	return total
}

func pathUp(path string) string {
	// remove the trailing slash
	path = path[:len(path)-1]
	// find the last slash index
	idx := strings.LastIndex(path, "/")
	return path[:idx+1]
}

// too high: 112680374
// too low:    1027500
func part1(input string) int {
	dirs := parseCommands(input)

	sizes := make(map[string]int)
	for path, data := range dirs {
		dirPlusChildren := findChildDirs(path, dirs)
		dirPlusChildren[path] = data
		sizes[path] = sumDirList(dirPlusChildren)
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
	return 0
}

func walkTree(tree *FileNode, dirSizes *DirSizes) *DirSizes {
	totalSize := 0

	for _, node := range tree.nodes {
		if tree.kind == "file" {
			totalSize += node.size
		} else if tree.kind == "dir" {
			newTree := walkTree(tree, dirSizes)
			totalSize += (*newTree)[node.name]
		} else {
			panic("ahh")
		}
	}

	(*dirSizes)[tree.name] = totalSize
	return dirSizes
}
