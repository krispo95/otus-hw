package main

import (
	"log"
	"os"
	"otus-hw/28.08.2020-linear-sort/externalMergeSort"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename must be set")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = externalMergeSort.MergeSort(file)
	if err != nil {
		panic(err)
	}
}
