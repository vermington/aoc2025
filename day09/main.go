package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func lines(s string, trim bool) []string {
	// Split into non-empty, trimmed lines
	sc := bufio.NewScanner(strings.NewReader(s))
	out := []string{}
	for sc.Scan() {
		ln := sc.Text()
		if trim {
			ln = strings.TrimSpace(ln)
		}
		out = append(out, ln)
	}
	return out
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(fmt.Errorf("atoi(%q): %w", s, err))
	}
	return n
}

func parseInput(input string) [][]int {
	coordinates := make([][]int, 0)
	for _, ln := range lines(input, true) {
		list := strings.Split(ln, ",")
		point := make([]int, 2)
		for i := range 2 {
			point[i] = atoi(list[i])
		}
		coordinates = append(coordinates, point)
	}
	return coordinates
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func area(from []int, to []int) int {
	return (abs(from[0]-to[0]) + 1) * (abs(from[1] - to[1] + 1))
}

func rangeFence(coors [][]int) [][]int {
	newCoors := make([][]int, 0)
	minX := coors[0][0]
	minY := coors[0][1]
	maxX := coors[0][0]
	maxY := coors[0][1]
	for _, c := range coors {
		if c[0] < minX {
			minX = c[0]
		}
		if c[1] < minY {
			minY = c[1]
		}
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}
	for _, c := range coors {
		point := make([]int, 0)
		point = append(point, c[0]-minX)
		point = append(point, c[1]-minY)
		newCoors = append(newCoors, point)
	}
	return newCoors
}

func fenced(coors [][]int, i int, j int) bool {
	xmin := -1
	xmax := -1
	ymin := -1
	ymax := -1
	if coors[i][0] < coors[j][0] {
		xmin = coors[i][0]
		xmax = coors[j][0]
	} else {
		xmin = coors[i][0]
		xmax = coors[j][0]
	}
	if coors[i][1] < coors[j][1] {
		ymin = coors[i][1]
		ymax = coors[j][1]
	} else {
		ymin = coors[i][1]
		ymax = coors[j][1]
	}

	fmt.Print(xmin, xmax, ymin, ymax)
	fmt.Print("\n")
	fmt.Print(area(coors[i], coors[j]))
	fmt.Print("\n")
	for _, c := range coors {
		//		if c[0] == coors[i][0] && c[1] == coors[i][1] {
		//			continue
		//		}
		//		if c[0] == coors[j][0] && c[1] == coors[j][1] {
		//			continue
		//		}
		if c[0] > xmin && c[0] < xmax && c[1] > ymin && c[1] < ymax {
			return false
		}
	}
	return true

}

// --- your puzzle logic ---

func part1(coors [][]int) int {
	max := 0
	for i := 0; i < len(coors); i++ {
		for j := i + 1; j < len(coors); j++ {
			calc := area(coors[i], coors[j])
			if calc > max {
				max = calc
			}
		}
	}
	return max
}

func part2(coors [][]int) int {
	max := 0
	//	coors = rangeFence(coors)
	for i := 0; i < len(coors); i++ {
		for j := i + 1; j < len(coors); j++ {
			if fenced(coors, i, j) {
				calc := area(coors[i], coors[j])
				if calc > max {
					max = calc
				}
			}
		}
	}
	return max
}

// --- entrypoint ---

func main() {
	in1 := readFile("example.input")
	coors := parseInput(in1)

	res1 := part1(coors)
	fmt.Printf("Part 1: %v\n", res1)

	res2 := part2(coors)
	fmt.Printf("Part 2: %v\n", res2)
}
