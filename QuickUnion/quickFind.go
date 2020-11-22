package main

import "fmt"

func main() {
	slice := createSlice(5)
	fmt.Println(slice)
}

func createSlice(size int) []int {
	var sl []int
	for i:= 0; i < size; i++ {
		sl = append(sl, i)
	}
	return sl
}

func connected(p, q int) {
	return 
}
