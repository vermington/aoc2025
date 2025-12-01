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

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(fmt.Errorf("atoi(%q): %w", s, err))
	}
	return n
}

func lineProc(s string) (first string, rest int) {
	return s[:1], atoi(s[1:])
}

// --- your puzzle logic ---

func part1(input string) any {
	mul := 0
	pos := 50
	zeroes := 0
	for _, ln := range lines(input) {
		dir, delta := lineProc(ln)
		if dir == "R" {
			mul = 1
		} else {
			mul = -1
		}
		pos += mul * delta
		if pos%100 == 0 {
			zeroes += 1
		}
	}
	return zeroes
}

func part2(input string) any {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mul := 0
	pos := 50
	zeroes := 0
	prevZero := 0
	for _, ln := range lines(input) {
		dir, delta := lineProc(ln)

		// Identify direction of travel
		if dir == "R" {
			mul = 1
		} else {
			mul = -1
		}

		// Normalise delta
		zeroes += delta / 100
		delta = delta % 100

		// Actual position change
		pos += mul * delta
		fmt.Fprintln(f, "Task:", dir, delta)
		fmt.Fprintln(f, "Live pos (add):", pos)

		if pos%100 == 0 {
			// Identify if zero
			zeroes += 1
			pos = 0
			prevZero = 1
			fmt.Fprintln(f, "Live zeroes (0):", zeroes)
		} else if pos > 100 {
			// Normalise if > 100
			zeroes += 1 - prevZero
			pos = pos % 100
			prevZero = 0
			fmt.Fprintln(f, "Live zeroes (>100):", zeroes)
		} else if pos < 0 {
			// Normalise if < 100
			zeroes += 1 - prevZero
			pos = 100 + pos%100
			prevZero = 0
			fmt.Fprintln(f, "Live zeroes (<100):", zeroes)
		} else {
			// Reset prevZero if no boundary crossing
			prevZero = 0
		}
	}
	return zeroes
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	res1 := part1(in1)
	fmt.Printf("Part 1: %v\n", res1)

	in2 := readFile("real.input")
	res2 := part2(in2)
	fmt.Printf("Part 2: %v\n", res2)
}
