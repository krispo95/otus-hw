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
	s.AddToFixedPosition(1, 0)
	s.Add(8)
	fmt.Println(s.Get(0))
	s.AddToFixedPosition(2, 0)
	fmt.Println(s.Get(0))
	s.AddToFixedPosition(3, 0)
	fmt.Println(s.Get(0))
	fmt.Println(s.Remove(1), s.Size())
	v := VectorArray{}
	v.New(100)
	v.AddToFixedPosition(1, 0)
	fmt.Println(v.Get(0))
	v.AddToFixedPosition(2, 0)
	v.Add(8)
	fmt.Println(v.Get(0))
	v.AddToFixedPosition(3, 0)
	fmt.Println(v.Get(0))
	fmt.Println(v.Remove(0), v.Size())
	f := FactorArray{}
	f.New(2)
	f.AddToFixedPosition(1, 0)
	fmt.Println(f.Get(0))
	f.AddToFixedPosition(2, 0)
	f.Add(8)
	fmt.Println(f.Get(0))
	f.AddToFixedPosition(3, 0)
	fmt.Println(f.Get(0))
	fmt.Println(f.Remove(0), f.Size())

}
