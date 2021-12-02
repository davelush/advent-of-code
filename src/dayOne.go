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

func shiftArray(arr [3]int, i int) [3]int {
	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = i
	return arr
}

func dayOne() {
	f, err := os.Open("src/input")
	check(err)
	scanner := bufio.NewScanner(f)

	var windowA [3]int
	var windowB [3]int
	var increaseCount int = -1
	var rowNum int = 0

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		windowA = shiftArray(windowA, depth)
		if rowNum >= 3 {
			if sum(windowA) > sum(windowB) {
				increaseCount++
			}
		}
		windowB = shiftArray(windowB, depth)
		rowNum++
	}
	fmt.Println(increaseCount)
}
