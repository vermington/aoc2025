package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct{ first, last int }

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

func parseInput(input string) ([]Range, []int) {
	ranges := make([]Range, 0)
	items := make([]int, 0)
	flag := true
	for _, ln := range lines(input) {
		if flag {
			if ln != "" {
				parts := strings.Split(ln, "-")
				a, b := atoi(parts[0]), atoi(parts[1])
				ranges = append(ranges, Range{first: a, last: b})
			} else {
				flag = false
			}
			continue
		} else {
			items = append(items, atoi(ln))
		}
	}
	return ranges, items
}

// --- your puzzle logic ---

func part1(ranges []Range, items []int) int {
	sum := 0
	for _, item := range items {
		for _, limits := range ranges {
			if item >= limits.first && item <= limits.last {
				sum += 1
				break
			}
		}
	}
	return sum
}

func part2(ranges []Range) int {
	sum := 0
	reducedRanges := make([]Range, 0)
	rlen := 0

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].first < ranges[j].first
	})

	for idx, limits := range ranges {
		if idx == 0 {
			reducedRanges = append(reducedRanges, limits)
		} else if limits.first <= reducedRanges[rlen].last {
			if limits.last > reducedRanges[rlen].last {
				reducedRanges[rlen].last = limits.last
			} else {
				continue
			}
		} else {
			reducedRanges = append(reducedRanges, limits)
			rlen++
		}
	}

	for _, limits := range reducedRanges {
		sum += limits.last - limits.first + 1
	}
	return sum
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	ranges, items := parseInput(in1)

	res1 := part1(ranges, items)
	fmt.Printf("Part 1: %v\n", res1)

	res2 := part2(ranges)
	fmt.Printf("Part 2: %v\n", res2)
}
