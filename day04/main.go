package main

import (
	"2025/charmatrix"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --- file & parsing helpers ---

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		// Fail loudly so you notice missing files
		panic(fmt.Errorf("read %s: %w", path, err))
	}
	return string(b)
}

func lines(s string) []string {
	// Split into non-empty, trimmed lines
	sc := bufio.NewScanner(strings.NewReader(s))
	out := []string{}
	for sc.Scan() {
		ln := strings.TrimSpace(sc.Text())
		if ln != "" {
			out = append(out, ln)
		}
	}
	return out
}

func countNeighbours(grid charmatrix.Matrix, pos charmatrix.Pos, limit int) bool {
	nbrs := grid.Neighbours(pos, charmatrix.EightDirs)
	count := 0
	total := len(nbrs)

	for i, nb := range nbrs {
		if grid.At(nb) == '@' {
			count++
			if count == limit {
				return false // already at or above the limit
			}
		}
		if remaining := total - i - 1; count+remaining < limit {
			return true // can’t reach the limit anymore
		}
	}
	return true // never hit the limit
}

// --- your puzzle logic ---

func part1(input string) (int, charmatrix.Matrix) {
	sum := 0
	grid := charmatrix.FromStrings(lines(input))
	clone := grid
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Cols(); c++ {
			pos := charmatrix.Pos{R: r, C: c}
			if grid.At(pos) != '@' {
				continue
			}
			if countNeighbours(grid, pos, 4) {
				sum += 1
				clone.Set(pos, '.')
			}
		}
	}
	return sum, clone
}

func part2(grid charmatrix.Matrix) (int, charmatrix.Matrix) {
	sum := 0
	clone := grid
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Cols(); c++ {
			pos := charmatrix.Pos{R: r, C: c}
			if grid.At(pos) != '@' {
				continue
			}
			if countNeighbours(grid, pos, 4) {
				sum += 1
				clone.Set(pos, '.')
			}
		}
	}
	return sum, clone
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	res1, grid := part1(in1)
	fmt.Printf("Part 1: %v\n", res1)

	res2 := res1
	for res1 > 0 {
		res1, grid = part2(grid)
		res2 += res1
	}

	fmt.Printf("Part 2: %v\n", res2)
}
