package main

import (
	"fmt"
)

func main() {
	/*
	 * Boolean type
	 */

	// Declaring a boolean type
	var a bool = false
	fmt.Printf("%v, %T\n", a, a)

	b := 1 == 1
	c := 1 == 0

	fmt.Printf("%v, %T\n", b, b) // true
	fmt.Printf("%v, %T\n", c, c) // false

	// In go, everytime we initialize a variable,
	// it has a zero (falsy) value
	var d bool

	fmt.Printf("%v, %T\n", d, d) // false

	/*
	 * Numeric types
	 */

	// a general 'int' type would be at least
	// 32 bits, but it could be 64 bits, 128 bits
	// or more depending on the underlying system.
	// So, it will always take the maximum possible.

	// Signed Integers (default)
	e := 42
	fmt.Printf("%v, %T\n", e, e)

	// Unsigned Integer
	var f uint16 = 42
	fmt.Printf("%v, %T\n", f, f)

	// In go when two integers will be divided then
	// the result cannot be a float type
	g := 10            // 1010
	h := 3             // 0011
	fmt.Println(g / h) // 3 (not 3.3333)

	// We also cannot perform arithmetic operations
	// between two different types
	// fmt.Println(e + f);
	// To do this we have to use conversion functions
	// to convert the type of one to the other.
	fmt.Println(e + int(f))

	// Bitwise operations are also possible.
	// Again, this operation can only be performed
	// over variables of same type.
	fmt.Println(g & h) // 1010 & 0011 = 0010 (2)
	fmt.Println(g | h) // 1010 | 0011 = 1011 (11)
	fmt.Println(g ^ h) // 1010 ^ 0011 = 1001 (9)
	// &^ is AND NOT operator (Not X-NOR)
	// x &^ y means x & ~y
	fmt.Println(g &^ h) // 1010 &^ 0011 = 1010 & ~(0011) = 1010 & 1100 = 1000 (8)

	// Bit Shifting is also possible.
	i := 8
	fmt.Println(i << 3) // 8 * 2^3 = 3 * 8 = 24
	fmt.Println(i >> 3) // 8 / 2^3 = 8 / 8 = 1

	/*
	 * Floating point
	 */
	j := 3.14 // By default the highest precesion is assumed.
	fmt.Printf("%v, %T\n", j, j)
	j = 12.7e72 // (for 12.7E72) 12.7 * 10^72
	fmt.Printf("%v, %T\n", j, j)
	j = 2.1e14 // (or 2.1E14)
	fmt.Printf("%v, %T\n", j, j)

	// Explicitly specifying type
	var k float32 = 3.14
	fmt.Printf("%v, %T\n", k, k)

	// Arithmetic operations
	// We don't have modulus operator for floating numbers
	// We don't have bitwise/bitshift operator for floating numbers
	l := 10.2
	m := 3.7
	fmt.Println(l + m)
	fmt.Println(l - m)
	fmt.Println(l * m)
	fmt.Println(l / m)

	/*
	 * Complex Numbers
	 */
	var n complex64 = 1 + 2i // complex64 by default.
	fmt.Printf("%v, %T\n", n, n)

	// i is a special symbol and this variable
	// will also be read as a complex number
	o := 2i
	fmt.Printf("%v, %T\n", o, o)

	// Arithmetic operations
	p := 1 + 2i
	q := 2 + 5.2i
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)

	// But given a complex number, how do we
	// extract only the real part or
	// the imaginary part ?
	// Using real() and imag() functions
	// If we are working on complex64 numbers,
	// we'll get their answer as float32 types
	// If we are working on complex128 numbers,
	// we'll get their answer as float64 types.
	fmt.Printf("%v, %T\n", real(p), real(p))
	fmt.Printf("%v, %T\n", imag(p), imag(p))

	// Looks fine!
	// But how do I build a complex numbers
	// given that I have both the real and imaginary
	// parts available.
	// Using complex() function
	r := complex(5, 12)
	fmt.Printf("%v, %T\n", r, r)

	/*
	 * Text types.
	 */

	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	// string are basically a collection of indexed characters
	// i.e. and array of character items.
	fmt.Printf("%v, %T\n", s[2], s[2]) // 105

	// But the above statement prints the UTF-8 equivalent instead.
	// Solution: Use string() conversion method.
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// String modification is not allowed.
	// s[2] = "u"

	// We can perform string concatenation
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// Converting string into slice of bytes
	t := []byte(s)
	fmt.Printf("%v, %T\n", t, t) // array of UTF-8 values

	// Declaring a Rune
	u := 'a'
	fmt.Printf("%v, %T\n", u, u) // 97, int32

	var v rune = 'a'
	fmt.Printf("%v, %T\n", v, v) // 97, int32
}
