package main

type SingleArray struct {
	array []int
}

func (s *SingleArray) Size() int {
	return len(s.array)
}
func (s *SingleArray) Add(item int, index int) {
	newArr := s.Resize(s.array, 1)
	isAdded := false
	if len(s.array) == 0 {
		newArr[0] = item
		s.array = newArr
		return
	}
	for i := range s.array {
		if i == index {
			newArr[i] = item
			newArr[i+1] = s.array[i]
			isAdded = true
			continue
		}
		if !isAdded {
			newArr[i] = s.array[i]
		} else {
			newArr[i+1] = s.array[i]
		}
	}
	s.array = newArr
}
func (s *SingleArray) Get(index int) int {
	return s.array[index]
}
func (s *SingleArray) Remove(index int) int {
	newArr := s.Resize(s.array, -1)
	deletedElement := 0
	isDeleted := false
	for i := range s.array {
		if i == index {
			deletedElement = s.array[i]
			isDeleted = true
			continue
		}
		if !isDeleted {
			newArr[i] = s.array[i]
		} else {
			newArr[i-1] = s.array[i]
		}
	}
	s.array = newArr
	return deletedElement
}
func (s *SingleArray) Resize(array []int, delta int) []int {
	newArr := make([]int, len(array)+delta)
	return newArr
}
