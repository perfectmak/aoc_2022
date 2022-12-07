package main

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day6"
)

func main() {
	result, err := day6.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", result)
}
