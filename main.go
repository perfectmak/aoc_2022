package main;

import (
	"fmt"

	"github.com/perfectmak/aoc_2022/day1"
)


func main() {
	result, err := day1.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result);
}
