package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve() (uint64, error) {
	return Solve2()
}

func toRangePair(rangeStr ...string) [2]uint64 {
	start, err := strconv.ParseUint(rangeStr[0], 10, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseUint(rangeStr[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return [2]uint64{start, end}
}

func isWithin(parent, child [2]uint64) bool {
	return parent[0] <= child[0] && parent[1] >= child[1]
}

func Solve1() (uint64, error) {
	//inputFile, err := os.Open("day4/sample.txt")
	inputFile, err := os.Open("day4/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return 0, err
	}

	totalScore := uint64(0)
	for inputReader.Scan() {
		line := strings.Split(inputReader.Text(), ",")
		firstRange := toRangePair(strings.Split(line[0], "-")...)
		secondRange := toRangePair(strings.Split(line[1], "-")...)

		if isWithin(firstRange, secondRange) || isWithin(secondRange, firstRange) {
			// either one within
			totalScore++
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}

func isOverlap(first, second [2]uint64) bool {
	return (first[0] <= second[0] && first[1] >= second[0]) ||
		(second[0] <= first[0] && second[1] >= first[0])
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day4/sample.txt")
	inputFile, err := os.Open("day4/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return 0, err
	}

	totalScore := uint64(0)
	for inputReader.Scan() {
		line := strings.Split(inputReader.Text(), ",")
		firstRange := toRangePair(strings.Split(line[0], "-")...)
		secondRange := toRangePair(strings.Split(line[1], "-")...)

		if isWithin(firstRange, secondRange) ||
			isWithin(secondRange, firstRange) ||
			isOverlap(firstRange, secondRange) {
			// either one overlaps
			totalScore++
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}
