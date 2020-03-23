package main

import (
	"./heap";
	"fmt"
)

func main() {
	var input = []int{10,50,8, 9,44,11,23,48, 78,78, 77, 79}
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
	sortedArray := heap.Heapsort(input)

	fmt.Println("Input Array", input)
	fmt.Println("Sorted input Array", sortedArray)
}