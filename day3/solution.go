package day3

import (
	"bufio"
	"os"
)

func Solve() (uint64, error) {
	return Solve2()
}

type Set map[rune]bool

func stringToSet(value string) Set {
	set := Set{}
	for _, char := range value {
		set[char] = true
	}

	return set
}

func setIntersection(sets ...Set) []rune {
	if len(sets) == 0 {
		return []rune{}
	}

	intersection := []rune{}
	firstSet := sets[0]
	remainingSets := sets[1:]
	for item, _ := range firstSet {
		existsInOthers := true
		for _, set := range remainingSets {
			if !set[item] {
				existsInOthers = false
				break
			}
		}

		if existsInOthers {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func Solve1() (uint64, error) {
	//inputFile, err := os.Open("day3/sample.txt")
	inputFile, err := os.Open("day3/input.txt")
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
		ruckSackItems := inputReader.Text()
		sackLen := len(ruckSackItems)
		firstHalf := ruckSackItems[:(sackLen / 2)]
		secondHalf := ruckSackItems[(sackLen / 2):]

		// TODO: create a generic function to return set map from any input
		firstSet := stringToSet(firstHalf)
		secondSet := stringToSet(secondHalf)

		intersection := setIntersection(firstSet, secondSet)

		for _, item := range intersection {
			itemInt := int(item)
			if itemInt >= 97 && itemInt <= 122 {
				totalScore += uint64(itemInt-int('a')) + 1
			}

			if itemInt >= 65 && itemInt <= 90 {
				totalScore += uint64(itemInt-int('A')) + 27
			}
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day3/sample.txt")
	inputFile, err := os.Open("day3/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)
	if err != nil {
		return 0, err
	}

	totalScore := uint64(0)

	sets := []Set{}
	for inputReader.Scan() {
		sets = append(sets, stringToSet(inputReader.Text()))
		if len(sets) != 3 {
			continue
		}

		intersection := setIntersection(sets...)

		for _, item := range intersection {
			itemInt := int(item)
			if itemInt >= 97 && itemInt <= 122 {
				totalScore += uint64(itemInt-int('a')) + 1
			}

			if itemInt >= 65 && itemInt <= 90 {
				totalScore += uint64(itemInt-int('A')) + 27
			}
		}

		sets = []Set{}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	return totalScore, nil
}
