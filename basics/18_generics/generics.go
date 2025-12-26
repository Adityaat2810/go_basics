package main

import "fmt"

/*
WHY GENERICS?
-------------
Without generics, we write the same logic multiple times
for different data types → code duplication (violates DRY).
*/

// Non-generic version (code duplication)
func printSliceInt(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSliceString(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/*
GENERIC FUNCTION
----------------
T is a type parameter.
`any` means T can be any type.
*/
func printSliceGeneric[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/*
GENERIC FUNCTION WITH CONSTRAINT
--------------------------------
Only allows int, string, or bool.
*/
func printSliceRestricted[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/*
COMPARABLE CONSTRAINT
---------------------
`comparable` means:
T supports == and != operators.

Used when equality comparison is needed.
*/
func contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

/*
GENERIC STRUCT
--------------
Stack can hold elements of ANY type.
*/
type Stack[T any] struct {
	elements []T
}

// Push method
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// Pop method
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // zero value of T
		return zero, false
	}
	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return last, true
}

func main() {
	// int slice
	itemsInt := []int{1, 2, 3, 4, 5}
	printSliceInt(itemsInt)

	// string slice
	itemsString := []string{"golang", "ts", "js", "python"}
	printSliceString(itemsString)

	// GENERIC function works for both
	printSliceGeneric(itemsInt)
	printSliceGeneric(itemsString)

	// Restricted generic
	printSliceRestricted([]bool{true, false, true})

	// comparable example
	fmt.Println(contains(itemsInt, 3))          // true
	fmt.Println(contains(itemsString, "java"))  // false
	fmt.Println(contains([]bool{true}, false))  // false

	// Generic stack with int
	stackInt := Stack[int]{}
	stackInt.Push(10)
	stackInt.Push(20)

	// Generic stack with string
	stackString := Stack[string]{}
	stackString.Push("hello")
	stackString.Push("world")

	fmt.Println(stackInt)
	fmt.Println(stackString)
}
