package main

import (
	"flag"
	"github.com/krispo95/otus-tester/tester"
	"strconv"
)

type Test struct{}

func (t Test) Run(str []string) []string {
	res := StringLength(str)
	return []string{res}
}

func main() {
	path := flag.String("path",
		"C:/Users/pogos/workspace/go/otus/tickets/0.String",
		"path to test files")
	test := Test{}
	tester.RunTest(*path, test)
}
func StringLength(str []string) string {
	if len(str) == 0 {
		return strconv.Itoa(0)
	}
	return strconv.Itoa(len(str[0]))
}
