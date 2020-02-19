package main

type FactorArray struct {
	array  []int
	factor int
}

func (f *FactorArray) New(factor int) {
	f.array = make([]int, 0, 0)
	f.factor = factor
}

func (f *FactorArray) Size() int {
	return len(f.array)
}
func (f *FactorArray) Add(item int) {
	newArr := f.resize(f.factor)
	newArr = append(newArr, item)
	f.array = newArr
}

func (f *FactorArray) AddToFixedPosition(item int, index int) {
	newArr := []int{}
	if f.Size() == cap(f.array) {
		newArr = f.makeEmptyArrayWithNewSize(len(f.array)+1, (len(f.array)+1)*f.factor)
	} else {
		newArr = f.makeEmptyArrayWithNewSize(len(f.array)+1, cap(f.array))
	}
	if len(f.array) == 0 {
		newArr[0] = item
		f.array = newArr
		return
	}
	isAdded := false
	for i := range f.array {
		if i == index {
			newArr[i] = item
			newArr[i+1] = f.array[i]
			isAdded = true
			continue
		}
		if !isAdded {
			newArr[i] = f.array[i]
		} else {
			newArr[i+1] = f.array[i]
		}
	}
	f.array = newArr
}
func (f *FactorArray) Get(index int) int {
	return f.array[index]
}
func (f *FactorArray) Remove(index int) int {
	newArr := []int{}
	if cap(f.array)/len(f.array) == f.factor {
		newArr = f.makeEmptyArrayWithNewSize(len(f.array)-1, len(f.array)/f.factor)
	} else {
		newArr = f.makeEmptyArrayWithNewSize(len(f.array)-1, cap(f.array))
	}
	deletedElement := 0
	isDeleted := false
	for i := range f.array {
		if i == index {
			deletedElement = f.array[i]
			isDeleted = true
			continue
		}
		if !isDeleted {
			newArr[i] = f.array[i]
		} else {
			newArr[i-1] = f.array[i]
		}
	}
	f.array = newArr
	return deletedElement
}

func (f *FactorArray) resize(factor int) []int {
	newArr := make([]int, len(f.array), len(f.array)*factor)
	for i := range f.array {
		newArr[i] = f.array[i]
	}
	return newArr
}

func (f *FactorArray) makeEmptyArrayWithNewSize(leng int, cap int) []int {
	newArr := make([]int, leng, cap)
	return newArr
}
