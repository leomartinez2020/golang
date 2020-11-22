package main

import "fmt"

func main() {
	p1 := Point{3, 5}
	p2 := Point{6, 9}
	fmt.Println(p1.SlopeTo(p2))
	p3 := Point{4, 6}
	p4 := Point{8, 1}
	fmt.Println(p3.CompareTo(p4))
	p5 := Point{2, 4}
	p6 := Point{5, 7}
	fmt.Println(p5.CompareTo(p6))
}

type Point struct {
	x, y float64
}

func (p1 Point) SlopeTo(p2 Point) float64 {
	return (p2.y - p1.y) / (p2.x - p1.x)
}

func (p1 Point) CompareTo(p2 Point) int {
	if p1.y == p2.y && p1.x < p2.x {
		return 1
	}
	if p1.y == p2.y && p1.x == p2.x {
		return 0
	}
	if p1.y > p2.y {
		return 1
	} else {
		return -1
	}
}
