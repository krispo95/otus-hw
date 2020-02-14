package main

import (
	"flag"
	"github.com/krispo95/otus-tester/tester"
	"strings"
)

type Test struct{}

func (t Test) Run(str []string) []string {

	res := BitBoardTrucker(str[0])
	sliceRes := strings.Split(res, "\n")
	return sliceRes
}

func main() {
	path := flag.String("path",
		"/home/krispo/workspace/0.BITS/4.Bitboard - Дальнобойщики",
		"path to test files")
	test := Test{}
	tester.RunTest(*path, test)
}
