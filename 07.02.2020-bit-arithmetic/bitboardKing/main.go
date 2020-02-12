package main

import (
	"flag"
	"fmt"
	"github.com/krispo95/otus-tester/tester"
	"strconv"
)

type Test struct{}

func (t Test) Run(str []string) []string {
	position, err := strconv.ParseUint(str[0], 10, 64)
	if err != nil {
		panic(err)
	}
	res := BitBoardKing(position)
	stringResults := make([]string, len(res))
	for i := range res {
		stringResults[i] = fmt.Sprint(res[i])
	}
	return stringResults
}

func main() {
	path := flag.String("path",
		"/home/krispo/workspace/0.BITS/1.Bitboard - Король",
		"path to test files")
	test := Test{}
	tester.RunTest(*path, test)
}
