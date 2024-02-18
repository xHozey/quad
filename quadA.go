package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	// Print the first row
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune('o')
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune('\n')

	// Print the middle rows
	for j := 0; j < y-2; j++ {
		z01.PrintRune('|')
		for i := 0; i < x-2; i++ {
			z01.PrintRune(' ')
		}
		if x != 1 && y != 5 {
			z01.PrintRune('|')
		}
		z01.PrintRune('\n')
	}

	// Print the last row if y > 1
	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(1, 5)
}
