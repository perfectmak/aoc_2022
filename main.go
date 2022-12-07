package main

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day7"
)

func main() {
	result, err := day7.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", result)
}
