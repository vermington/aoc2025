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

func parseInput(input string) ([][]int, []byte) {
	numbers := make([][]int, 0)
	operators := make([]byte, 0)
	for _, ln := range lines(input, true) {
		list := strings.Fields(ln)
		if list[0][0] < '0' {
			for _, l := range list {
				operators = append(operators, l[0])
			}
		} else {
			for i, l := range list {
				if len(numbers) <= i {
					numbers = append(numbers, []int{})
				}
				numbers[i] = append(numbers[i], atoi(l))
			}
		}
	}
	return numbers, operators
}

func computeSlice(slice []string) int {
	sum := 0
	cols := len(slice)
	width := len(slice[0]) - 1
	op := slice[cols-1][0]
	if op == 42 {
		sum = 1
	} else {
		sum = 0
	}
	num := make([]int, width)
	for i := width - 1; i >= 0; i-- {
		for _, n := range slice[:cols-1] {
			if n[i] != ' ' {
				num[i] = num[i]*10 + int(n[i]-'0')
			}
		}
	}
	for _, n := range num {
		if op == 42 {
			sum *= n
		} else {
			sum += n
		}
	}
	return sum
}

// --- your puzzle logic ---

func part1(numbers [][]int, operators []byte) int {
	total := 0
	sum := 0
	for i, op := range operators {
		if op == 42 {
			sum = 1
		} else {
			sum = 0
		}
		for _, n := range numbers[i] {
			if op == 42 {
				sum *= n
			} else {
				sum += n
			}
		}
		total += sum
	}
	return total
}

func part2(input string) int {
	total := 0
	universe := lines(input, false)
	full := len(universe)
	width := len(universe[0])
	slice := make([]string, full)
	for w := range width {
		flag := true
		for i, ln := range universe {
			slice[i] += string(ln[w])
			if ln[w] != ' ' {
				flag = false
			}
		}
		if flag {
			total += computeSlice(slice)
			slice = make([]string, full)
		}
	}
	return total
}

// --- entrypoint ---

func main() {
	in1 := readFile("real.input")
	nums, ops := parseInput(in1)

	res1 := part1(nums, ops)
	fmt.Printf("Part 1: %v\n", res1)

	in2 := readFile("real.input")
	res2 := part2(in2)
	fmt.Printf("Part 2: %v\n", res2)
}
