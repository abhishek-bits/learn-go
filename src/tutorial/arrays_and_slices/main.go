package main

import (
	"fmt"
)

func main() {

	/*************************************************
	 * 						ARRAYS					 *
	 *************************************************/

	// An array containing the grades
	// grades := [3]int{97, 85, 93}

	// Another method to declare the same array
	// allowing go compiler to interpret the size
	// on its own.
	grades := [...]int{97, 85, 93}

	fmt.Printf("Grades: %v\n", grades)

	// Declaring an empty array
	var students [3]string

	fmt.Printf("Students: %v\n", students)

	// Populating array using index
	students[0] = "Lisa"

	fmt.Printf("Students: %v\n", students)

	students[1] = "Ahmed"
	students[2] = "Arnold"

	fmt.Printf("Students: %v\n", students)

	// Accessing the elements using indexed notation
	fmt.Printf("Student #1: %v\n", students[1])

	/*
	 * Built-in array methods
	 */
	// 1. len() is used to find the length of the array.
	// NOTE: length here means the total length of the array
	// not essentially the current elements count.
	fmt.Printf("Number of Students: %v\n", len(students))

	// Initializing a 2D Matrix
	// var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	// OR a lot cleaner syntax
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// NOTE:-
	// Unlike other languages, in Go, when we assign
	// an array to a variable, then it basically creates
	// a deep copy of this array i.e. it does not point
	// to the same memory location.
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [1, 5, 3]

	// But if we use the ampersand operator to point
	// to the memory address instead, then it would
	// point to the same memory location.
	// Here, a is the array itself while c is pointing to a.
	c := &a
	c[1] = 5
	fmt.Println(a) // [1, 5, 3]
	fmt.Println(c) // [1, 5, 3]

	/****************************************************
	 * 						SLICES						*
	 ****************************************************/

	// Unlike arrays that have their size
	// fixed at compile time,
	// Slices on the other hand are dynamic.
	d := []int{1, 2, 3}
	fmt.Println(d)

	// Just like arrays, slices have len()
	// but there length denotes the current
	// number of elements in the array.
	fmt.Printf("Length: %v\n", len(d))

	// In addition to length, with slices
	// we can also find out the capacity.
	fmt.Printf("Capacity: %v\n", cap(a))

	// Unlike arrays, slices are naturally reference types
	e := d
	e[1] = 5
	fmt.Println(d) // [1, 5, 3]
	fmt.Println(e) // [1, 5, 3]

	// More pratical application of slices
	// NOTE:
	// the first index is inclusive
	// and end index is exclusive.
	f := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	g := f[:]   // slice of all elements
	h := f[3:]  // slice from 4th element to end.
	i := f[:6]  // slice first 6 elements
	j := f[3:6] // slice from 4th element to 6th element.

	// Now, even if we had sliced the array already
	// If we update an element of f, then also all slices
	// will updates that respective element as they are
	// all pointing to the same array.
	f[5] = 42

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	// Another way to create a slice is to use
	// the built-in make() method with arguments:
	// 1. type of slice that we want to create
	// 2. the length of this slice
	k := make([]int, 3)
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(a))

	// Extended make() method:
	// here, third arument is the
	// initial capacity of this slice.
	k2 := make([]int, 3, 100)
	fmt.Println(k2)
	fmt.Printf("Length: %v\n", len(k2))
	fmt.Printf("Capacity: %v\n", cap(k2))

	// An example to show how capacity increases
	// as and when the elements are inserted / removed.
	a2 := []int{}
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))
	// We'll use the built-in append() method to
	// insert an element to the slice.
	a2 = append(a2, 1)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))
	a2 = append(a2, 2, 3, 4, 5)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// NOTE that the second argument in append()
	// may be a slice followed by a spread operator
	a2 = append(a2, []int{2, 3, 4, 5}...) // append(a2, 2, 3, 4, 5)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	/*
	 * Using a Slice as Stack / Deque
	 */
	l := []int{1, 2, 3, 4, 5}

	m := l[1:]     // Remove first element
	fmt.Println(m) // [2,3,4,5]

	n := l[:len(l)-1] // Remove last element
	fmt.Println(n)    // [1,2,3,4]

	// But, what if we want to remove
	// an item from the middle?
	// Solution: Use append()
	o := append(l[:2], l[3:]...)
	fmt.Println(o) // [1,2,4,5]

	// We have to be extra careful when working with append()
	// as the underlying slice will also go through the
	// changes we provide.
	fmt.Println(l)

	// Is there any alternative solution to this ?
	// Indirectly YES, we simply need to create a
	// deep copy of the underlying slice first, if only,
	// we will use it in the future.
}
