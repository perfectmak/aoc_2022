package day5

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	stack []string
}

func (s *Stack) Push(values ...string) {
	s.stack = append(s.stack, values...)
}

func (s *Stack) Pop() string {
	if s.Empty() {
		return ""
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return top
}

func (s *Stack) PopN(n int) []string {
	if s.Empty() || n == 0 {
		return []string{}
	}

	top := s.stack[len(s.stack)-n:]
	s.stack = s.stack[:len(s.stack)-n]

	return top
}

func (s *Stack) Top() string {
	if s.Empty() {
		return ""
	}

	return s.stack[len(s.stack)-1]
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Reverse() {
	newStack := Stack{}
	for !s.Empty() {
		newStack.Push(s.Pop())
	}
	s.stack = newStack.stack
}

func Solve() (string, error) {
	return Solve2()
}

func Solve1() (string, error) {
	//inputFile, err := os.Open("day5/sample.txt")
	inputFile, err := os.Open("day5/input.txt")
	defer inputFile.Close()
	if err != nil {
		return "", err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return "", err
	}

	exp := regexp.MustCompile("( {3}) {1}")
	stacks := make([]Stack, 10)
	maxStackPos := 0
	for inputReader.Scan() {
		line := inputReader.Text()
		if line == "" {
			// stack process next are move instructions
			break
		}

		if string(line[0:2]) == " 1" {
			continue
		}

		cleanedStr := strings.ReplaceAll(exp.ReplaceAllString(line, "X"), " ", "")
		stackPos := 0
		for _, value := range cleanedStr {
			if stackPos > maxStackPos {
				maxStackPos = stackPos
			}

			if value == '[' || value == ']' {
				continue
			}

			if value == 'X' {
				stackPos++
				continue
			}

			//fmt.Printf("Pushing Values: %s; Pos %d\n", string(value), stackPos)
			stacks[stackPos].Push(string(value))
			stackPos++
		}
	}

	for i := 0; i <= maxStackPos; i++ {
		stacks[i].Reverse()
	}

	//fmt.Printf("Stacks: %v\n", stacks)
	for inputReader.Scan() {
		line := inputReader.Text()
		parts := strings.Split(line, " ")
		moveX, _ := strconv.ParseInt(parts[1], 10, 32)
		fromS, _ := strconv.ParseInt(parts[3], 10, 32)
		toS, _ := strconv.ParseInt(parts[5], 10, 32)

		for i := int64(0); i < moveX; i++ {
			value := stacks[fromS-1].Pop()
			stacks[toS-1].Push(value)
		}
		//fmt.Printf("Stacks: %v\n", stacks)
	}

	if inputReader.Err() != nil {
		return "", inputReader.Err()
	}

	result := ""
	for i := 0; i <= maxStackPos; i++ {
		result += stacks[i].Top()
	}

	return result, nil
}

func Solve2() (string, error) {
	//inputFile, err := os.Open("day5/sample.txt")
	inputFile, err := os.Open("day5/input.txt")
	defer inputFile.Close()
	if err != nil {
		return "", err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return "", err
	}

	exp := regexp.MustCompile("( {3}) {1}")
	stacks := make([]Stack, 10)
	maxStackPos := 0
	for inputReader.Scan() {
		line := inputReader.Text()
		if line == "" {
			// stack process next are move instructions
			break
		}

		if string(line[0:2]) == " 1" {
			continue
		}

		cleanedStr := strings.ReplaceAll(exp.ReplaceAllString(line, "X"), " ", "")
		stackPos := 0
		for _, value := range cleanedStr {
			if stackPos > maxStackPos {
				maxStackPos = stackPos
			}

			if value == '[' || value == ']' {
				continue
			}

			if value == 'X' {
				stackPos++
				continue
			}

			//fmt.Printf("Pushing Values: %s; Pos %d\n", string(value), stackPos)
			stacks[stackPos].Push(string(value))
			stackPos++
		}
	}

	for i := 0; i <= maxStackPos; i++ {
		stacks[i].Reverse()
	}

	//fmt.Printf("Stacks: %v\n", stacks)
	for inputReader.Scan() {
		line := inputReader.Text()
		parts := strings.Split(line, " ")
		moveX, _ := strconv.ParseInt(parts[1], 10, 32)
		fromS, _ := strconv.ParseInt(parts[3], 10, 32)
		toS, _ := strconv.ParseInt(parts[5], 10, 32)

		value := stacks[fromS-1].PopN(int(moveX))
		stacks[toS-1].Push(value...)
		//fmt.Printf("Stacks: %v\n", stacks)
	}

	if inputReader.Err() != nil {
		return "", inputReader.Err()
	}

	result := ""
	for i := 0; i <= maxStackPos; i++ {
		result += stacks[i].Top()
	}

	return result, nil
}
