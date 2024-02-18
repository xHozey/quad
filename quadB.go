package main

import (
	"fmt"
)

func QuadA(x, y int) {
	for i := 0; i < x; i++ {
		if i == 0 {
			fmt.Printf("/")
		} else if i == x-1 {
			fmt.Printf("\\")
		} else {
			fmt.Printf("*")
		}
	}
	fmt.Printf("\n")

	for j := 0; j < y-2; j++ {
		fmt.Printf("*")
		for i := 0; i < x-2; i++ {
			fmt.Printf(" ")
		}
		if x != 1 && y != 5 {
			fmt.Printf("*")
		}
		fmt.Printf("\n")

	}

	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 {
				fmt.Printf("\\")
			} else if i == x-1 {
				fmt.Printf("/")
			} else {
				fmt.Printf("*")
			}
		}
	}
	fmt.Printf("\n")
}

func main() {
	QuadA(5, 3)
}
