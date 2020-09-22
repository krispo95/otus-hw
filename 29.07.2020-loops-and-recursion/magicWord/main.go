package main

import (
	"fmt"
	"math"
)

func main() {
	PrintWord(1)
	PrintWord(2)
	PrintWord(3)
	PrintWord(4)
	PrintWord(5)
	PrintWord(6)
	PrintWord(7)
	PrintWord(8)
	PrintWord(9)
	PrintWord(10)
	PrintWord(11)
	PrintWord(12)
	PrintWord(13)
	PrintWord(14)
	PrintWord(17)
	PrintWord(18)
	PrintWord(19)
	PrintWord(20)
	PrintWord(22)
	PrintWord(23)
	PrintWord(24)
	PrintWord(25)

}
func PrintWord(num int) {
	fmt.Println()
	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			switch num {
			case 1:
				fmt.Print(Word1(x, y))
			case 2:
				fmt.Print(Word2(x, y))
			case 3:
				fmt.Print(Word3(x, y))
			case 4:
				fmt.Print(Word4(x, y))
			case 5:
				fmt.Print(Word5(x, y))
			case 6:
				fmt.Print(Word6(x, y))
			case 7:
				fmt.Print(Word7(x, y))
			case 8:
				fmt.Print(Word8(x, y))
			case 9:
				fmt.Print(Word9(x, y))
			case 10:
				fmt.Print(Word10(x, y))
			case 11:
				fmt.Print(Word11(x, y))
			case 12:
				fmt.Print(Word12(x, y))
			case 13:
				fmt.Print(Word13(x, y))
			case 14:
				fmt.Print(Word14(x, y))
			case 17:
				fmt.Print(Word17(x, y))
			case 18:
				fmt.Print(Word18(x, y))
			case 19:
				fmt.Print(Word19(x, y))
			case 20:
				fmt.Print(Word20(x, y))
			case 22:
				fmt.Print(Word22(x, y))
			case 23:
				fmt.Print(Word23(x, y))
			case 24:
				fmt.Print(Word24(x, y))
			case 25:
				fmt.Print(Word25(x, y))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func Word1(x, y int) string {
	if x < y {
		return "  #"
	}
	return "  ."
}
func Word2(x, y int) string {
	if x == y {
		return "  #"
	}
	return "  ."
}
func Word3(x, y int) string {
	if x+y == 24 {
		return "  #"
	}
	return "  ."
}
func Word4(x, y int) string {
	if x+y < 30 {
		return "  #"
	}
	return "  ."
}
func Word5(x, y int) string {
	if y-2*x < 2 && y-2*x >= 0 {
		return "  #"
	}
	return "  ."
}
func Word6(x, y int) string {
	if x < 10 || y < 10 {
		return "  #"
	}
	return "  ."
}
func Word7(x, y int) string {
	if x > 15 && y > 15 {
		return "  #"
	}
	return "  ."
}
func Word8(x, y int) string {
	if x*y == 0 {
		return "  #"
	}
	return "  ."
}
func Word9(x, y int) string {
	if x-y > 11 || y-x > 11 {
		return "  #"
	}
	return "  ."
}
func Word10(x, y int) string {
	if y-x < x && x < y {
		return "  #"
	}
	return "  ."
}
func Word11(x, y int) string {
	if x == 1 || x == 23 || y == 1 || y == 23 {
		return "  #"
	}
	return "  ."
}
func Word12(x, y int) string {
	if y*y <= 400-x*x {
		return "  #"
	}
	return "  ."
}
func Word13(x, y int) string {
	if y > 19-x && y < 29-x {
		return "  #"
	}
	return "  ."
}
func Word14(x, y int) string {
	if x*y < 100 {
		return "  #"
	}
	return "  ."
}
func Word17(x, y int) string {
	if float64(x) > 15+math.Sin(float64(y)*math.Pi/10)*8 {
		return "  #"
	}
	return "  ."
}
func Word18(x, y int) string {
	if (x == 0 || x == 1 || y == 0 || y == 1) && !(x == 0 && y == 0) {
		return "  #"
	}
	return "  ."
}
func Word19(x, y int) string {
	if x == 0 || x == 24 || y == 0 || y == 24 {
		return "  #"
	}
	return "  ."
}
func Word20(x, y int) string {
	if (x+y)%2 == 0 {
		return "  #"
	}
	return "  ."
}

func Word22(x, y int) string {
	if (x+y)%3 == 0 {
		return "  #"
	}
	return "  ."
}
func Word23(x, y int) string {
	if x%3 == 0 && y%2 == 0 {
		return "  #"
	}
	return "  ."
}
func Word24(x, y int) string {
	if y == x || y == 24-x {
		return "  #"
	}
	return "  ."
}
func Word25(x, y int) string {
	if x%6 == 0 || y%6 == 0 {
		return "  #"
	}
	return "  ."
}
