package day1;

import (
	"os"
	"bufio"
	"strconv"
	"sort"
)

func Solve() (uint64, error) {
	return Solve2();
}



func Solve1() (uint64, error) {
	//inputFile, err := os.Open("day1/sample.txt");
	inputFile, err := os.Open("day1/input.txt");
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile);
	if err != nil {
		return 0, err
	}
	
	currentMax := uint64(0);
	runningSum := uint64(0);
	for inputReader.Scan() {
		line := inputReader.Text();
		if line == "" {
			if runningSum > currentMax {
				currentMax = runningSum
			}
			runningSum = 0;
			continue
		}

		lineNum, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return 0, err
		}

		runningSum += lineNum

	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}


	return currentMax, nil
}

func insertMax(values []uint64, newValue uint64) {
	for idx, value := range values {
		if value == 0 {
			values[idx] = newValue
			break
		}
		if newValue > value {
			values[idx] = newValue
			break
		}
	}

	//fmt.Printf("%v\n", values)
	sort.SliceStable(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	//fmt.Printf("%v\n", values)
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day1/sample.txt");
	inputFile, err := os.Open("day1/input.txt");
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile);
	if err != nil {
		return 0, err
	}
	
	maxValues := []uint64{0, 0, 0}
	runningSum := uint64(0);
	for inputReader.Scan() {
		line := inputReader.Text();
		if line == "" {
			insertMax(maxValues, runningSum)
			runningSum = 0;
			continue
		}

		lineNum, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return 0, err
		}

		runningSum += lineNum

	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	insertMax(maxValues, runningSum)

	currentMax := uint64(0);
	for _, value := range maxValues {
		currentMax += value
	}
	return currentMax, nil
}
