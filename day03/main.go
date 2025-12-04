package main

import (
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

func getMax(s string, start int, end int) (int, int) {
	max := s[start]
	joltPos := 0
	for idx, b := range []byte(s[start+1 : len(s)-end]) {
		if b > max {
			joltPos = idx + start + 1
			max = b
		}
	}
	jolts := int(max - '0')
	return jolts, joltPos
}

func getJoltage(s string, l int) int {
	l--
	jolts, pos := getMax(s, 0, l)
	for l > 0 {
		l--
		units := 0
		units, pos = getMax(s, pos+1, l)
		jolts = 10*jolts + units
	}
	return jolts
}

// --- your puzzle logic ---

func part1(input string) any {
	sum := 0
	for _, ln := range lines(input) {
		sum += getJoltage(ln, 2)
	}
	return sum
}

func part2(input string) any {
	sum := 0
	for _, ln := range lines(input) {
		j := getJoltage(ln, 3)
		fmt.Printf("jolts: %v\n", j)
		sum += j
	}
	return sum
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	res1 := part1(in1)
	fmt.Printf("Part 1: %v\n", res1)

	in2 := readFile("example.input")
	res2 := part2(in2)
	fmt.Printf("Part 2: %v\n", res2)
}
