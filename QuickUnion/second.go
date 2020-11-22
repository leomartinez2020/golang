package main

import "fmt"

func main() {
	seq := Lista{1,2,3,4,5}
	fmt.Println(seq)
	fmt.Println(seq.Connected(1,2))
	seq.Union(1,2)
	fmt.Println(seq.Connected(1,2))
}

type Lista []int

func (l Lista) Connected(p, q int) bool {
	return l[p] == l[q]
}

func (l *Lista) Union(p, q int) {
	pid := l[p]
	qid := l[q]
	for i := 0; i < len(l); i++ {
		if l[i] == pid {
			l[i] = qid
		}
	}
}
