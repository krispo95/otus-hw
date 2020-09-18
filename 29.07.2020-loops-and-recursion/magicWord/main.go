package main

import "fmt"

type Test struct{}

func main() {
	PrintWord(1)
}
func PrintWord(num int) {
	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			switch num {
			case 1:
				fmt.Print(Word1(x, y))
			case 2:
				fmt.Print(Word2(x, y))
			case 3:
				fmt.Print(Word3(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
				//case 1: fmt.Print(Word1(x, y))
			}
		}
		fmt.Println()
	}
}
func Word1(x, y int) string {
	if x < y {
		return "# "
	}
	return ". "
}
func Word2(x, y int) string {
	if x == y {
		return "# "
	}
	return ". "
}
func Word3(x, y int) string {
	if x+y == 24 {
		return "# "
	}
	return ". "
}
