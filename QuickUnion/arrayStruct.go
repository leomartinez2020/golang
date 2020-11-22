package main

import "fmt"

func main() {
	var seq Lista
	PopulateArray(&seq, 10)
	fmt.Println(seq)
	seq.Union(4, 3)
	seq.Union(3, 8)
	seq.Union(6, 5)
	seq.Union(9, 4)
	seq.Union(2, 1)
	fmt.Println(seq.Connected(8, 9))
	fmt.Println(seq.Connected(5, 0))
	seq.Union(5, 0)
	seq.Union(7, 2)
	seq.Union(6, 1)
	fmt.Println(seq)
	//seq.Union(1,2)
	//fmt.Println(seq.Connected(1,2))
}

type Lista struct {
	arr []int
}

func PopulateArray(l *Lista, size int) {
	for i := 0; i < size; i++ {
		l.arr = append(l.arr, i)
	}
}

func (l Lista) Connected(p, q int) bool {
	return l.arr[p] == l.arr[q]
}

func (l *Lista) Union(p, q int) {
	pid := l.arr[p]
	qid := l.arr[q]
	for i := 0; i < len(l.arr); i++ {
		if l.arr[i] == pid {
			l.arr[i] = qid
		}
	}
}
