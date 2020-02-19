package main

type MatrixArray struct {
	array [][100]int
}

func (m *MatrixArray) New() {
	m.array = make([][100]int, 0)
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
	newArr := m.makeEmptyArrayWithNewSize()
	isAdded := false
	if len(m.array) == 0 {
		newArr[0] = item
		m.array = newArr
		return
	}
	for i := range m.array {
		if i == index {
			newArr[i] = item
			newArr[i+1] = m.array[i]
			isAdded = true
			continue
		}
		if !isAdded {
			newArr[i] = m.array[i]
		} else {
			newArr[i+1] = m.array[i]
		}
	}
	m.array = newArr
}
func (m *MatrixArray) Get(index int) int {
	return m.array[index]
}
func (m *MatrixArray) Remove(index int) int {
	newArr := m.resize(-1)
	deletedElement := 0
	isDeleted := false
	for i := range m.array {
		if i == index {
			deletedElement = m.array[i]
			isDeleted = true
			continue
		}
		if !isDeleted {
			newArr[i] = m.array[i]
		} else {
			newArr[i-1] = m.array[i]
		}
	}
	m.array = newArr
	return deletedElement
}

func (m *MatrixArray) resize() {
	m.array = append(m.array, [100]int{0})
}

func (m *MatrixArray) makeEmptyArrayWithNewSize() [][]int {
	inArr := make([]int, 0, len(m.array[len(m.array)-1])-1)
	newArr := [][]int{inArr}
	return newArr
}
