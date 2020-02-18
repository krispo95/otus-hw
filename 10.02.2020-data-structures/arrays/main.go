package main

import "fmt"

type Array interface {
	Size() int
	Add(item int64, index int)
	Get(index int) int64
	Remove(index int) int64
}

func main() {
	s := SingleArray{}
	s.New()
	s.Add(1, 0)
	fmt.Println(s.Get(0))
	s.Add(2, 0)
	fmt.Println(s.Get(0))
	s.Add(3, 0)
	fmt.Println(s.Get(0))
	fmt.Println(s.Remove(1), s.Size())
	v := VectorArray{}
	v.New(100)
	v.Add(1, 0)
	fmt.Println(v.Get(0))
	v.Add(2, 0)
	fmt.Println(v.Get(0))
	v.Add(3, 0)
	fmt.Println(v.Get(0))
	fmt.Println(v.Remove(0), v.Size())

}
