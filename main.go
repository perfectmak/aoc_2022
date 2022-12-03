package main

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day3"
)

func main() {
	result, err := day3.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result)
}
