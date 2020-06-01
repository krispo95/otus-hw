package main

type HeapWalker struct {
	slice    []int
	heapSize int
}

func MakeHeapWalker(slice []int) HeapWalker {
	h := HeapWalker{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MakeHeap(i)
	}
	return h
}

func (h HeapWalker) MakeHeap(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.heapSize && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.heapSize && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MakeHeap(max)
	}
}

func heapSort(slice []int) []int {
	h := MakeHeapWalker(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MakeHeap(0)
	}
	return h.slice
}
