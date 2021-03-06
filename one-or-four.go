package main

import (
	"fmt"
)

func oof(num int) int {
	fmt.Printf("%d", num)
	if num == 0 || num == 1 || num == 4 {
		return num
	}
	fmt.Printf(" -> ")
	sum := 0
	if num < 0 {
		num = -num
	}
	for {
		sum += (num % 10) * (num % 10)
		num = num / 10
		if num <= 9 {
			sum += num * num
			break
		}
	}
	return oof(sum)
}

func main() {
	var num int
	fmt.Printf("Input Number: ")
	fmt.Scanf("%d", &num)
	oof(num)
	fmt.Printf("\n")
}
