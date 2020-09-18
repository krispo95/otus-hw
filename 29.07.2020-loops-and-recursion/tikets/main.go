package main

import (
	"flag"
	"github.com/krispo95/otus-tester/tester"
	"strconv"
)

type Test struct{}

func (t Test) Run(str []string) []string {
	res := LuckyTickets(str)
	return []string{res}
}

func main() {
	path := flag.String("path",
		"C:/Users/pogos/workspace/go/otus/tickets/1.Tickets",
		"path to test files")
	test := Test{}
	tester.RunTest(*path, test)
}

func LuckyTickets(str []string) string {
	if len(str) == 0 {
		return strconv.Itoa(0)
	}
	n, _ := strconv.Atoi(str[0])
	sums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		sums = CreateNextColumn(sums)
	}
	count := 0
	for _, cnt := range sums {
		count += cnt * cnt
	}
	return strconv.Itoa(count)
}

func CreateNextColumn(prevColumn []int) (nextColumn []int) {
	newLen := len(prevColumn) + 9
	for i := 0; i < newLen; i++ {
		newItem := 0
		for j := 0; j <= 9; j++ {
			if i-j >= 0 && i-j <= len(prevColumn)-1 {
				newItem += prevColumn[i-j]
			}
		}
		nextColumn = append(nextColumn, newItem)
	}
	return nextColumn
}
