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
func (v *VectorArray) Add(item int, index int) {
	newArr := v.array
	if v.Size() == cap(v.array) {
		newArr = v.Resize(v.array, v.vector)
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
			newArr[i+1] = v.array[i]
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
	newArr := v.array
	if cap(v.array)-len(v.array) == v.vector {
		newArr = v.Resize(v.array, -v.vector)
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
func (v *VectorArray) Resize(array []int, delta int) []int {
	newArr := make([]int, len(array)+1, len(array)+delta)
	return newArr
}
