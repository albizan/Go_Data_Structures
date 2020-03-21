package main

import (
	"./heap";
	"fmt"
)

func main() {
	var input = []int{1,2,3,4,5}
	
	h := heap.MakeHeapFromArray(input)
	h2 := heap.MakeHeap(10)

	h2.Insert(10)
	h2.Insert(20)
	h2.Insert(30)
	h2.Insert(40)
	h2.Insert(50)

	fmt.Println(h)
	fmt.Println(h2)

	h.Pop()
	h2.Pop()

	fmt.Println(h)
	fmt.Println(h2)
}