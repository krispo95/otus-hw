package main

import (
	"flag"
	"fmt"
	"github.com/krispo95/otus-tester/tester"
	"otus-hw/07.02.2020-bit-arithmetic/bitboardKing"
	"strconv"
)

type Test struct{}

func (t Test) Run(str []string) []string {
	position, err := strconv.ParseUint(str[0], 10, 64)
	if err != nil {
		panic(err)
	}
	res := bitboardKing.BitBoardKing(position)
	stringResults := make([]string, len(res))
	for i := range res {
		stringResults[i] = fmt.Sprint(res[i])
	}
	return stringResults
}

func main() {
	path := flag.String("path",
		"/home/krispo/workspace/go_projects/otus-hw-test-files/0.BITS/1.BitboardKing",
		"path to test files")
	test := Test{}
	tester.RunTest(*path, test)
}
