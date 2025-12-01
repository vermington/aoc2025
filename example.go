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

// --- your puzzle logic ---

func part1(input string) any {
  // TODO: implement
  // example: sum integers, one per line
  sum := 0
  for _, ln := range lines(input) {
    sum += atoi(ln)
  }
  return sum
}

func part2(input string) any {
  // TODO: implement
  return "not implemented"
}

// --- entrypoint ---

func main() {
  in1 := readFile("1.input")
  res1 := part1(in1)
  fmt.Printf("Part 1: %v\n", res1)

  in2 := readFile("2.input")
  res2 := part2(in2)
  fmt.Printf("Part 2: %v\n", res2)
}