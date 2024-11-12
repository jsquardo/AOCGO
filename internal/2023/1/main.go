package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

}

func part1(input string) int {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(error.Error(err))
	}

	var total, num1, num2 int
	for _, s := range f {
		if s == 10 {
			if num1 != 0 && num2 != 0 {
				lnAsStr := strconv.Itoa(num1) + strconv.Itoa(num2)
				joinedNum, _ := strconv.Atoi(lnAsStr)
				println("joined: ", joinedNum)
				total += joinedNum
			}

			if num1 != 0 && num2 == 0 {
				lnAsStr := strconv.Itoa(num1) + strconv.Itoa(num1)
				joinedNum, _ := strconv.Atoi(lnAsStr)
				println("joined: ", joinedNum)
				total += joinedNum
			}
			num1, num2 = 0, 0
		}

		if num, err := strconv.Atoi(string(s)); err == nil {
			if num1 == 0 {
				num1 = num
			}

			if num1 != 0 {
				num2 = num
			}
		} else {
			println(string(s))
		}
	}
	return total
}
