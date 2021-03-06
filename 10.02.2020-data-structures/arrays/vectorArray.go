package main

type VectorArray struct {
	array  []int
	vector int
}

func (v *VectorArray) New(vector int) {
	v.array = make([]int, 0, 0)
	v.vector = vector
}

func (v *VectorArray) Size() int {
	return len(v.array)
}
func (v *VectorArray) Add(item int) {
	newArr := v.resize(v.vector)
	newArr = append(newArr, item)
	v.array = newArr
}

func (v *VectorArray) AddToFixedPosition(item int, index int) {
	newArr := []int{}
	if v.Size() == cap(v.array) {
		newArr = v.makeEmptyArrayWithNewSize(len(v.array)+1, cap(v.array)+v.vector)
	} else {
		newArr = v.makeEmptyArrayWithNewSize(len(v.array)+1, cap(v.array))
	}
	isAdded := false
	if len(v.array) == 0 {
		newArr[0] = item
		v.array = newArr
		return
	}
	for i := range v.array {
		if i == index {
			newArr[i] = item
			newArr[i] = v.array[i]
			isAdded = true
			continue
		}
		if !isAdded {
			newArr[i] = v.array[i]
		} else {
			newArr[i+1] = v.array[i]
		}
	}
	v.array = newArr
}
func (v *VectorArray) Get(index int) int {
	return v.array[index]
}
func (v *VectorArray) Remove(index int) int {
	newArr := []int{}
	if cap(v.array)-len(v.array) == v.vector {
		newArr = v.makeEmptyArrayWithNewSize(len(v.array)-1, len(v.array)-v.vector)
	} else {
		newArr = v.makeEmptyArrayWithNewSize(len(v.array)-1, cap(v.array))
	}
	deletedElement := 0
	isDeleted := false
	for i := range v.array {
		if i == index {
			deletedElement = v.array[i]
			isDeleted = true
			continue
		}
		if !isDeleted {
			newArr[i] = v.array[i]
		} else {
			newArr[i-1] = v.array[i]
		}
	}
	v.array = newArr
	return deletedElement
}

func (v *VectorArray) resize(delta int) []int {
	newArr := make([]int, len(v.array), len(v.array)+delta)
	for i := range v.array {
		newArr[i] = v.array[i]
	}
	return newArr
}

func (v *VectorArray) makeEmptyArrayWithNewSize(leng int, cap int) []int {
	newArr := make([]int, leng, cap)
	return newArr
}
