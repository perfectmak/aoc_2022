package day2

import (
	"bufio"
	"os"
	"strings"
)

func Solve() (uint64, error) {
	return Solve2()
}

const (
	RockScore     uint64 = 1
	PaperScore           = 2
	ScissorsScore        = 3
	WinScore             = 6
	DrawScore            = 3
	LoseScore            = 0
)

var choiceScores = map[string]uint64{
	"AX": RockScore + DrawScore,
	"AY": PaperScore + WinScore,
	"AZ": ScissorsScore + LoseScore,
	"BX": RockScore + LoseScore,
	"BY": PaperScore + DrawScore,
	"BZ": ScissorsScore + WinScore,
	"CX": RockScore + WinScore,
	"CY": PaperScore + LoseScore,
	"CZ": ScissorsScore + DrawScore,
}

func Solve1() (uint64, error) {
	//inputFile, err := os.Open("day2/sample.txt")
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return 0, err
	}

	totalScore := uint64(0)
	for inputReader.Scan() {
		roundChoice := strings.Replace(inputReader.Text(), " ", "", -1)
		totalScore += choiceScores[roundChoice]
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}

var choiceScores2 = map[string]uint64{
	"AX": ScissorsScore + LoseScore,
	"AY": RockScore + DrawScore,
	"AZ": PaperScore + WinScore,
	"BX": RockScore + LoseScore,
	"BY": PaperScore + DrawScore,
	"BZ": ScissorsScore + WinScore,
	"CX": PaperScore + LoseScore,
	"CY": ScissorsScore + DrawScore,
	"CZ": RockScore + WinScore,
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day2/sample.txt")
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return 0, err
	}

	totalScore := uint64(0)
	for inputReader.Scan() {
		roundChoice := strings.Replace(inputReader.Text(), " ", "", -1)
		totalScore += choiceScores2[roundChoice]
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}
