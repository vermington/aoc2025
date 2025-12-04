package main

import (
	"bufio"
	"fmt"
	"math"
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

var pow10 = func() []int64 {
	out := make([]int64, 19)
	out[0] = 1
	for i := 1; i < len(out); i++ {
		out[i] = out[i-1] * 10
	}
	return out
}()

func atoi64(s string) int64 {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		panic(fmt.Errorf("atoi(%q): %w", s, err))
	}
	return n
}

func repeatPattern(pattern int64, repeats int, base int64) (int64, bool) {
	res := int64(0)
	for i := 0; i < repeats; i++ {
		if res > (math.MaxInt64-pattern)/base {
			return 0, false
		}
		res = res*base + pattern
	}
	return res, true
}

func repeatedNumbers(lo, hi int64, minRepeats, maxRepeats int) []int64 {
	if lo > hi {
		return nil
	}
	seen := make(map[int64]struct{})
	results := []int64{}

	minDigits := len(strconv.FormatInt(lo, 10))
	if minDigits < 2 {
		minDigits = 2
	}
	maxDigits := len(strconv.FormatInt(hi, 10))

	for digits := minDigits; digits <= maxDigits; digits++ {
		if digits >= len(pow10) {
			break
		}
		lower := lo
		if lower < pow10[digits-1] {
			lower = pow10[digits-1]
		}
		upper := hi
		upperCap := pow10[digits] - 1
		if upper > upperCap {
			upper = upperCap
		}
		if lower > upper {
			continue
		}

		for patternLen := 1; patternLen <= digits/2; patternLen++ {
			if digits%patternLen != 0 {
				continue
			}
			repeats := digits / patternLen
			if repeats < minRepeats {
				continue
			}
			if maxRepeats > 0 && repeats > maxRepeats {
				continue
			}
			base := pow10[patternLen]
			startPattern := pow10[patternLen-1]
			endPattern := base - 1
			for pattern := startPattern; pattern <= endPattern; pattern++ {
				candidate, ok := repeatPattern(pattern, repeats, base)
				if !ok || candidate > upper {
					break
				}
				if candidate < lower {
					continue
				}
				if _, ok := seen[candidate]; ok {
					continue
				}
				seen[candidate] = struct{}{}
				results = append(results, candidate)
			}
		}
	}
	return results
}

func parseRange(tuple string) (int64, int64) {
	parts := strings.SplitN(strings.TrimSpace(tuple), "-", 2)
	if len(parts) != 2 {
		panic(fmt.Errorf("bad range %q", tuple))
	}
	lo := atoi64(parts[0])
	hi := atoi64(parts[1])
	return lo, hi
}

// --- your puzzle logic ---

func part1(input string) any {
	var sum int64
	for _, tuple := range strings.Split(input, ",") {
		if strings.TrimSpace(tuple) == "" {
			continue
		}
		lo, hi := parseRange(tuple)
		for _, val := range repeatedNumbers(lo, hi, 2, 2) {
			sum += val
		}
	}
	return sum
}

func part2(input string) any {
	var sum int64
	for _, tuple := range strings.Split(input, ",") {
		if strings.TrimSpace(tuple) == "" {
			continue
		}
		lo, hi := parseRange(tuple)
		for _, val := range repeatedNumbers(lo, hi, 2, 0) {
			sum += val
		}
	}
	return sum
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
