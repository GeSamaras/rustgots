package main

import (
	"fmt"
	"strings"
)

type Thing = int

const (
	Tree Thing = iota
	Snow
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	treeCount := 0
	for row, line := range strings.Split(getInput(), "\n") {
		if string(line[row*3%len(line)]) == "#" {
			treeCount++
		}
	}
	fmt.Println(treeCount)
}