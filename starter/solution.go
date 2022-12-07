package day4

import (
	"bufio"
	"os"
)

func Solve() (uint64, error) {
	return Solve1()
}

func Solve1() (uint64, error) {
	inputFile, err := os.Open("day4/sample.txt")
	//inputFile, err := os.Open("day4/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	totalScore := uint64(0)
	for inputReader.Scan() {
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}

func Solve2() (uint64, error) {
	inputFile, err := os.Open("day4/sample.txt")
	//inputFile, err := os.Open("day4/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	totalScore := uint64(0)
	for inputReader.Scan() {
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}
