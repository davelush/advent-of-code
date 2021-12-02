package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array [3]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	f, err := os.Open("src/input")
	check(err)
	scanner := bufio.NewScanner(f)

	var windowA [3]int
	var windowB [3]int
	var increaseCount int = -1

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		// update windows for comparison
		if sum(windowA) > sum(windowB) {
			increaseCount++
		}
		//
	}
	fmt.Println(increaseCount)
}
