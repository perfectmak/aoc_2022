package day6

import (
	"bufio"
	"os"
)

type Set map[rune]bool

func stringToSet(value string) Set {
	set := Set{}
	for _, char := range value {
		set[char] = true
	}

	return set
}

func Solve() (uint64, error) {
	return Solve2()
}

func Solve1() (uint64, error) {
	inputFile, err := os.Open("day6/sample.txt")
	//inputFile, err := os.Open("day6/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	totalScore := uint64(0)
	for inputReader.Scan() {
		cipher := inputReader.Text()
		for i := 4; i <= len(cipher); i++ {
			set := stringToSet(cipher[i-4 : i])
			if len(set) == 4 {
				totalScore = uint64(i)
				break
			}
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day6/sample.txt")
	inputFile, err := os.Open("day6/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	totalScore := uint64(0)
	for inputReader.Scan() {
		cipher := inputReader.Text()
		for i := 14; i <= len(cipher); i++ {
			set := stringToSet(cipher[i-14 : i])
			if len(set) == 14 {
				totalScore = uint64(i)
				break
			}
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}
