package main

import "fmt"

type MatrixArray struct {
	array [][]int
	cap   int
}

func (m *MatrixArray) New(cap int) {
	m.cap = cap
	m.array = make([][]int, 0, cap)
}

func (m *MatrixArray) Size() int {
	ln := 0
	for _, item := range m.array {
		ln += len(item)
	}
	return ln
}
func (m *MatrixArray) Add(item int) {
	m.resize()
	m.array[len(m.array)-1][len(m.array[len(m.array)-1])-1] = item
}

func (m *MatrixArray) AddToFixedPosition(item int, index int) {
	arrNum := 0
	mod := 0
	if index != 0 {
		arrNum = index / 100
		mod = index % 100
		if mod == 0 {
			arrNum--
			mod = 9
		}
	}
	newArr := m.makeEmptyArrayWithNewSize()
	isAdded := false
	if len(m.array) == 0 {
		inArr := make([]int, 1, m.cap)
		inArr[0] = item
		newArr[0] = inArr
		m.array = newArr
		return
	}
	for i := range m.array {
	loopIn:
		for ind := range m.array[i] {
			if i == arrNum {
				if ind != mod {
					continue loopIn
				}

				newArr[i][ind] = item
				isAdded = true
				if mod == 9 {
					newArr[i+1][0] = m.array[i][ind]
				}
			}
			if !isAdded {
				newArr[i][ind] = m.array[i][ind]
			} else {
				if mod == 9 && ind == 9 {
					newArr[i+1][0] = m.array[i][ind]
				} else {
					fmt.Println(i, ind)
					newArr[i][ind+1] = m.array[i][ind]
				}
			}
		}
	}
	m.array = newArr
}
func (m *MatrixArray) Get(index int) int {
	if index == 0 {
		return m.array[0][0]
	}
	arrNum := index / 100
	mod := index % 100
	if mod == 0 {
		arrNum--
		mod = 9
	}
	return m.array[arrNum][mod]
}

//func (m *MatrixArray) Remove(index int) int {
//	newArr := m.resize(-1)
//	deletedElement := 0
//	isDeleted := false
//	for i := range m.array {
//		if i == index {
//			deletedElement = m.array[i]
//			isDeleted = true
//			continue
//		}
//		if !isDeleted {
//			newArr[i] = m.array[i]
//		} else {
//			newArr[i-1] = m.array[i]
//		}
//	}
//	m.array = newArr
//	return deletedElement
//}

func (m *MatrixArray) resize() {
	inArr := make([]int, len(m.array[len(m.array)-1]), m.cap)
	m.array = append(m.array, inArr)
}

func (m *MatrixArray) makeEmptyArrayWithNewSize() [][]int {
	inArr := make([]int, 0, 100)
	newArr := [][]int{inArr}
	return newArr
}
