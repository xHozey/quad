package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	for i := 0; i < x; i++ {
		if i == 0 {
			z01.PrintRune('A')
		} else if i == x-1 {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune('\n')

	for j := 0; j < y-2; j++ {
		z01.PrintRune('B')
		for i := 0; i < x-2; i++ {
			z01.PrintRune(' ')
		}
		if x != 1 && y != 5 {
			z01.PrintRune('B')
		}
		z01.PrintRune('\n')

	}

	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 {
				z01.PrintRune('C')
			} else if i == x-1 {
				z01.PrintRune('A')
			} else {
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(5, 3)
}
