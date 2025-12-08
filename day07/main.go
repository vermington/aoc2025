package main

import (
	"2025/charmatrix"
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

func countPaths(grid charmatrix.Matrix, rows [][]int, r int, path []int) int {
	if r == len(rows) {
		for i := 0; i < len(path)-1; i++ {
			abs := path[i+1] - path[i]
			if abs < 0 {
				abs = -abs
			}
			if abs > 1 {
				return 0
			}
			fmt.Print(abs, i+1, path[i+1], grid.At(charmatrix.Pos{R: 2 * (i + 1), C: path[i+1]}))
			fmt.Print("\n")
			//			if abs == 1 && grid.At(charmatrix.Pos{R: i+1, C: path[i+1]}) != '^' {
			//				return 0
			//			}
			//			if abs == 0 && grid.At(charmatrix.Pos{R: i+1, C: path[i+1]}) == '^' {
			//				return 0
			//			}
		}
		fmt.Print(path)
		fmt.Print("\n")
		fmt.Print("\n")
		return 1 // reached a full path
	}
	total := 0
	for _, v := range rows[r] {
		total += countPaths(grid, rows, r+1, append(path, v))
	}
	return total
}

// --- your puzzle logic ---

func part1(grid charmatrix.Matrix) (int, int) {
	count := 0
	lastCount := 0
	rayList := make([]int, 0)
	listList := make([][]int, 0)
	for c := 0; c < grid.Cols(); c++ {
		if grid.At(charmatrix.Pos{R: 0, C: c}) == 'S' {
			rayList = append(rayList, c)
			break
		}
	}
	for r := 1; r < grid.Rows(); r++ {
		checkList := rayList
		rayList = make([]int, 0)
		for _, c := range checkList {
			if grid.At(charmatrix.Pos{R: r, C: c}) == '^' {
				count++
				rayList = append(rayList, c-1)
				rayList = append(rayList, c+1)
			} else {
				rayList = append(rayList, c)
			}
		}
		sort.Ints(rayList)
		rayList = slices.Compact(rayList)

		if count != lastCount {
			lastCount = count
			listList = append(listList, rayList)
			fmt.Print(rayList)
			fmt.Print("\n")
		} else {
			rayList = checkList
		}

	}

	paths := countPaths(grid, listList, 0, nil)

	return count, paths
}

// --- entrypoint ---

func main() {
	in1 := readFile("example.input")
	grid := charmatrix.FromStrings(lines(in1))

	res1, res2 := part1(grid)
	fmt.Printf("Part 1: %v\n", res1)

	fmt.Printf("Part 2: %v\n", res2)
}
