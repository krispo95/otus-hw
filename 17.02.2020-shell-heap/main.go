package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice := generateSlice(100)
	slice2 := generateAlmostSortedSlice(100)
	fmt.Println("\nslice before shell sort: \n\n", slice)
	shell(slice)
	fmt.Println("\nslice after shell sort: \n\n", slice)
	fmt.Println("\nslice before shell sort: \n\n", slice2)
	shell(slice2)
	fmt.Println("\nslice after shell sort: \n\n", slice2)

	fmt.Println("\nslice before shell sort: \n\n", slice)
	heapSort(slice)
	fmt.Println("\nslice after shell sort: \n\n", slice)
	fmt.Println("\nslice before shell sort: \n\n", slice2)
	heapSort(slice2)
	fmt.Println("\nslice after shell sort: \n\n", slice2)
}

func generateSlice(length int) []int {
	slice := make([]int, length, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func generateAlmostSortedSlice(length int) []int {
	slice := generateSlice(length)
	shell(slice)
	mixCount := length * 5 / 200
	for i := 0; i < mixCount; i++ {
		n, m := rand.Intn(length-1), rand.Intn(length-1)
		slice[n], slice[m] = slice[m], slice[n]
	}
	return slice
}
