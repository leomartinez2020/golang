package main

import "fmt"

func main() {
	var seq Lista
	PopulateArray(&seq, 10)
	fmt.Println(seq)
	fmt.Println(seq.Root(3))
	testQuickUnion()
}

func testQuickUnion() {
	var seq Lista
	PopulateArray(&seq, 10)
	fmt.Println(seq)
	seq.Union(4, 3)
	seq.Union(3, 8)
	seq.Union(6, 5)
	seq.Union(9, 4)
	seq.Union(2, 1)
	fmt.Println(seq.Connected(8, 9))
	fmt.Println(seq.Connected(5, 4))
	seq.Union(5, 0)
	seq.Union(7, 2)
	seq.Union(6, 1)
	seq.Union(7, 3)
	fmt.Println(seq)
}

type Lista struct {
	arr []int
}

func PopulateArray(l *Lista, size int) {
	for i := 0; i < size; i++ {
		l.arr = append(l.arr, i)
	}
}

func (l Lista) Root(i int) int {
	for i != l.arr[i] {
		i = l.arr[i]
	}
	return i
}

func (l Lista) Connected(p, q int) bool {
	return l.Root(p) == l.Root(q)
}

func (l *Lista) Union(p, q int) {
	i := l.Root(p)
	j := l.Root(q)
	l.arr[i] = j
}
