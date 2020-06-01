package main

import "testing"

var slice100000 = generateSlice(100000)
var sliceAlmost100000 = generateAlmostSortedSlice(100000)

var slice100000000 = generateSlice(100000000)
var sliceAlmost100000000 = generateAlmostSortedSlice(100000000)

func Benchmark_shell100000000(b *testing.B) {
	shell(slice100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		shell(slice100000000)
	}
}

func Benchmark_shellAlmost100000000(b *testing.B) {
	shell(sliceAlmost100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		shell(sliceAlmost100000000)
	}
}
func Benchmark_heapSort100000000(b *testing.B) {
	heapSort(slice100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapSort(slice100000000)
	}
}

func Benchmark_heapSortAlmost100000000(b *testing.B) {
	heapSort(sliceAlmost100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapSort(sliceAlmost100000000)
	}
}

func Benchmark_shell100000(b *testing.B) {
	shell(slice100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		shell(slice100000)
	}
}

func Benchmark_shellAlmost100000(b *testing.B) {
	shell(sliceAlmost100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		shell(sliceAlmost100000)
	}
}
func Benchmark_heapSort100000(b *testing.B) {
	heapSort(slice100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapSort(slice100000)
	}
}

func Benchmark_heapSortAlmost100000(b *testing.B) {
	heapSort(sliceAlmost100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapSort(sliceAlmost100000)
	}
}
