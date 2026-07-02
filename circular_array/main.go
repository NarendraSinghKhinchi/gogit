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
	newOffset := (ca.offset + offset) % size
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

	fmt.Println("Initial:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Get(%d): %d\n", i, ca.Get(i))
	}

	ca.Rotate(2)
	fmt.Println("\nAfter Rotate(2):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Get(%d): %d\n", i, ca.Get(i))
	}

	ca.Rotate(-1)
	fmt.Println("\nAfter Rotate(-1):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Get(%d): %d\n", i, ca.Get(i))
	}
}
