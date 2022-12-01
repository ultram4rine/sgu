package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("./example/example.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		uniqueOprtr = make(map[string]bool)
		uniqueOprnd = make(map[string]bool)
		operators   = make([]string, 0)
		operands    = make([]string, 0)
	)

	scanner := bufio.NewScanner(file)
	start := false
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if start {
			lOprtr, lOprnd := parseLine(line)
			for _, oprtr := range lOprtr {
				uniqueOprtr[oprtr] = true
			}
			for _, oprnd := range lOprnd {
				uniqueOprnd[oprnd] = true
			}
			operators = append(operators, lOprtr...)
			operands = append(operands, lOprnd...)
		} else {
			if strings.Contains(line, "func main()") {
				start = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nu1 := len(uniqueOprtr)
	nu2 := len(uniqueOprnd)
	N1 := len(operators)
	N2 := len(operands)

	nu := nu1 + nu2
	fmt.Printf("Program vocabulary: %d\n", nu)

	N := N1 + N2
	fmt.Printf("Program length: %d\n", N)

	fmt.Printf("Calculated estimated program length: %f\n", float64(nu1)*math.Log2(float64(nu1))+float64(nu2)*math.Log2(float64(nu2)))

	V := float64(N) * math.Log2(float64(nu))
	fmt.Printf("Volume: %f\n", V)

	D := (float64(nu1) / 2) * (float64(N2) / float64(nu2))
	fmt.Printf("Difficulty: %f\n", D)

	E := D * V
	fmt.Printf("Effort: %f\n", E)
}

var operatorsDict = []string{
	"break", "default", "func", "interface", "select",
	"case", "defer", "go", "map", "struct",
	"chan", "else", "goto", "package", "switch",
	"const", "fallthrough", "if", "range", "type",
	"continue", "for", "import", "return", "var", // Keywords
	",", ";", // Delimeters
	"|=", "^=", "|", "^", "<<", ">>", // Bit arifmetic
	"+=", "-=", "*=", "/=", "%=", ">=", "<=", "++", "--", "+", "-", "*", "/", "%", ">", "<", // Arifmetic
	":=", "=", // Assignment
}

var garbageDict = []string{
	"(", ")", "[", "]", "{", "}", "_",
}

func parseLine(line string) (operators, operands []string) {
	// Firstly move functions to operators slice.
	reg := regexp.MustCompile(`([a-zA-Z.]+)\(([^\)]*)\)(\.[^\)]*\))?`)
	if reg.MatchString(line) {
		operators = append(operators, reg.ReplaceAllString(line, "$1")+"()")
		line = reg.ReplaceAllString(line, "")
	}

	words := strings.Fields(line)

	// Move operators to operators slice.
	for _, op := range operatorsDict {
		for i, word := range words {
			if op == word {
				operators = append(operators, op)
				words = append(words[:i], words[i+1:]...)
			} else {
				// Operator and word can be concatenated
				if strings.Contains(word, op) {
					operators = append(operators, op)
					words[i] = strings.Replace(words[i], op, "", 1)
				}
			}
		}
	}

	// Remove garbage.
	for _, g := range garbageDict {
		for i, word := range words {
			if g == word {
				words = append(words[:i], words[i+1:]...)
			}
		}
	}

	operands = words

	return operators, operands
}
