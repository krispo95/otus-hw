package main

func MergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	middle := len(numbers) / 2
	left := numbers[:middle]
	right := numbers[middle:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) (res []int) {
	res = make([]int, len(left)+len(right))

	i := 0
	lPtr := 0
	rPtr := 0
	for lPtr <= len(left)-1 && rPtr <= len(right)-1 {
		if left[lPtr] < right[rPtr] {
			res[i] = left[lPtr]
			lPtr++
		} else {
			res[i] = right[rPtr]
			rPtr++
		}
		i++
	}

	for _ = 1; lPtr <= len(left)-1; lPtr++ {
		res[i] = left[lPtr]
		i++
	}
	for _ = 1; rPtr <= len(right)-1; rPtr++ {
		res[i] = right[rPtr]
		i++
	}
	return
}
