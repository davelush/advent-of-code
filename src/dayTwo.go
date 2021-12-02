package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input-daytwo")
	check(err)
	scanner := bufio.NewScanner(f)

	distance := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		str := scanner.Text()
		instr := strings.Fields(str)
		direction := instr[0]
		amount, _ := strconv.Atoi(instr[1])
		if direction == "forward" {
			distance += amount
			depth += aim * amount
		} else if direction == "down" {
			aim += amount
		} else if direction == "up" {
			aim -= amount
		}
	}
	fmt.Println(depth * distance)

}
