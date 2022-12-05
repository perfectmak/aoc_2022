package main

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day5"
)

func main() {
	result, err := day5.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", result)
}
