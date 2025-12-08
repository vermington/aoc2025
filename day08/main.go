package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		point := make([]int, 3)
		for i := range 3 {
			point[i] = atoi(list[i])
		}
		coordinates = append(coordinates, point)
	}
	return coordinates
}

func dist(from []int, to []int) int {
	return (from[0]-to[0])*(from[0]-to[0]) +
		(from[1]-to[1])*(from[1]-to[1]) +
		(from[2]-to[2])*(from[2]-to[2])

}

func contains(array []int, check int) bool {
	for _, a := range array {
		if check == a {
			return true
		}
	}
	return false
}

// --- your puzzle logic ---

func part1(coors [][]int) (int, int) {
	distList := make([]int, 0)
	fromList := make([]int, 0)
	toList := make([]int, 0)
	mul := 1
	for i := 0; i < len(coors); i++ {
		for j := i + 1; j < len(coors); j++ {
			distList = append(distList, dist(coors[i], coors[j]))
			fromList = append(fromList, i)
			toList = append(toList, j)
		}
	}

	idx := make([]int, len(distList))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return distList[idx[i]] < distList[idx[j]]
	})

	circuits := make([][]int, 0)
	savedCircuits := make([][]int, 0)
	for n, i := range idx {
		if n == 1000 {
			savedCircuits = make([][]int, len(circuits))
			for i, c := range circuits {
				savedCircuits[i] = append([]int(nil), c...)
			}
		}
		fromIdx := -1
		toIdx := -1
		for p, c1 := range circuits {
			if contains(c1, fromList[i]) {
				fromIdx = p
			}
			if contains(c1, toList[i]) {
				toIdx = p
			}
		}
		if fromIdx == -1 && toIdx == -1 {
			circ := make([]int, 2)
			circ[0], circ[1] = fromList[i], toList[i]
			circuits = append(circuits, circ)
		} else if fromIdx == -1 {
			circuits[toIdx] = append(circuits[toIdx], fromList[i])
		} else if toIdx == -1 {
			circuits[fromIdx] = append(circuits[fromIdx], toList[i])
		} else if toIdx != fromIdx {
			fromCirc := circuits[fromIdx]
			toCirc := circuits[toIdx]
			newCircuits := make([][]int, 0)
			newCircuits = append(newCircuits, append(append([]int(nil), fromCirc...), toCirc...))
			for p, c := range circuits {
				if p != fromIdx && p != toIdx {
					newCircuits = append(newCircuits, c)
				}
			}
			circuits = newCircuits
		}
		if len(circuits[0]) == 1000 {
			mul = coors[fromList[i]][0] * coors[toList[i]][0]
			break
		}
	}

	lenList := make([]int, 0)
	for _, c := range savedCircuits {
		lenList = append(lenList, len(c))
	}
	sort.Slice(lenList, func(i, j int) bool { return lenList[i] > lenList[j] })

	return lenList[0] * lenList[1] * lenList[2], mul
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	coors := parseInput(in1)

	res1, res2 := part1(coors)
	fmt.Printf("Part 1: %v\n", res1)

	fmt.Printf("Part 2: %v\n", res2)
}
