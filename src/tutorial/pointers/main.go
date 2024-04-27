package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {

	var a int = 42 // a := 42
	fmt.Println(a)

	var b *int = &a // b holds the address of variable a.
	fmt.Println(&a, b)
	fmt.Println(a, *b) // de-referencing operator.

	a = 27
	fmt.Println(a, *b)

	*b = 14 // using de-reference operator to change the value
	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)

	// Pointer arithmetic is not allowed in Go
	// err := &c[2] - 2

	// In case, if we really need to do pointer arithmetic
	// We can refer to "unsafe" package.

	/*
	 * Assigning a pointer to an object
	 */

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms) // &{42}

	var ms2 *myStruct
	fmt.Println(ms2) // <nil>
	ms2 = new(myStruct)
	fmt.Println(ms2) // &{0}

	// How do we check whether an object is null
	// before we even try to read its fields ?
	// We can use dereference operator.
	ms2.foo = 14     // (*ms2).foo
	fmt.Println(ms2) // &{14}

	// But Wait! ms2 is just a pointer to an object.
	// How does it even know that there is some property
	// named "foo" in it.
	// Well this is just Go compiler helping us out here.

	/*
	 * Arrays, Slices, Maps and Pointers
	 *
	 * A slice is basically a pointer to the underlying array.
	 * When we copy one slice to another, it actually copies
	 * the memory location of the first item of that underlying
	 * array. Therefore sharing slices accross the application
	 * actually means pointing to the same underlying data.
	 *
	 * Maps have a pointer to the underlying data and they
	 * do not contain the actual data itself.
	 */

	p := []int{1, 2, 3}
	q := p // copies the pointer to the address of the first item in the array
	fmt.Println(p, q)
	p[1] = 42
	fmt.Println(p, q)

	r := map[string]string{"foo": "bar", "biz": "buz"}
	s := r
	fmt.Println(r, s)
	r["foo"] = "qux"
	fmt.Println(r, s)

}
