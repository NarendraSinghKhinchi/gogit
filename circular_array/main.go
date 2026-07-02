package main

import "fmt"

type CustomArray[T any] struct {
	arr    []T
	offset int
}

func NewArray[T any](size int) *CustomArray[T] {
	return &CustomArray[T]{
		arr:    make([]T, size),
		offset: 0,
	}
}

func (ca *CustomArray[T]) Set(item T, idx int) {
	size := len(ca.arr)
	index := (idx + ca.offset) % size
	ca.arr[index] = item
}

func (ca *CustomArray[T]) Get(idx int) T {
	size := len(ca.arr)
	index := (idx + ca.offset) % size
	return ca.arr[index]
}

func (ca *CustomArray[T]) Rotate(offset int) {
	size := len(ca.arr)
	newOffset := (ca.offset - offset) % size
	if newOffset < 0 {
		newOffset += size
	}
	ca.offset = newOffset
}

func main() {
	ca := NewArray[int](5)

	ca.Set(1, 0)
	ca.Set(2, 1)
	ca.Set(3, 2)
	ca.Set(4, 3)
	ca.Set(5, 4)

	printArray("Initial:", ca)

	ca.Rotate(2)
	printArray("After Rotate(2) [Left shift by 2]:", ca)

	ca.Rotate(-1)
	printArray("After Rotate(-1) from initial [Right shift by 1]:", ca)

	ca.Rotate(-1) // cumulative -2
	printArray("After Rotate(-2) from initial [Right shift by 2]:", ca)
}

func printArray(msg string, ca *CustomArray[int]) {
	fmt.Println(msg)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", ca.Get(i))
	}
	fmt.Println("\n")
}
