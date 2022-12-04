package main

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day4"
)

func main() {
	result, err := day4.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result)
}
