package main

import (
	"fmt"
	"strconv"
)

/*
 * Package level variable declarating (Full Declaration Syntax)
 */
var i float32 = 42

/*
 * Declraring a block of variables which may/may not be related
 */
var (
	actorName   string = "Elizabeth Sladen"
	companion   string = "Sarah Jane Smith"
	doctorNumer int    = 3
	season      int    = 11
)

func main() {

	/*
	 * Declaring first and then initializing
	 */

	// Declaring a variable in go
	var a int

	// Initializing
	a = 42

	fmt.Println(a)

	// Updating
	a = 27

	fmt.Println(a)

	/*
	 * Declaration and initialization at the same time.4
	 * (Full Declaration Syntax)
	 */

	// This type of declaration is important when go
	// cannot infer the type on its own or
	// what the variable can be assigned to at run-time.
	var b int = 42

	fmt.Println(b)

	/*
	 * Declaration and Initialization shorthand statement
	 */

	// When we do both declaration and initialization
	// at the same time, then we don't have to declare
	// the type, go compiler will inference that on its own.
	c := 42

	fmt.Println(c)

	/*
	 * Uses of printf()
	 */
	var d float32 = 27

	// %v is used to fetch the value of the variable
	// %T is used to fetch the type of the variable
	fmt.Printf("%v, %T\n", d, d)

	fmt.Printf("%v, %T\n", c, c)

	// The problem with automatic type inference is that
	// the value is always inferred to be of highest precession
	// This is the limitiation of this shorthand initialization.
	// So, when we want to have more control over type,
	// we should go with the full declaration syntax
	e := 42.

	fmt.Printf("%v, %T\n", e, e) // float64

	fmt.Printf("%v, %T\n", i, i)

	/*
	 * Variable Shadowing
	 */

	// Declarating a variable already avaialable in the
	// global scope will override the existing value and its type.
	// i.e. variable with the innermost scope takes precedence.
	// here, variable 'i; becomes shadow variable.
	var i int = 12

	fmt.Printf("%v, %T\n", i, i)

	/*
	 * Naming Conventions
	 */

	// lower case variable names are scoped internal to the package / method.

	// Upper case variable names will trigger the go compiler to
	// expose the variable to the outside world.

	// Acronyms such as URL, ID, etc. should always be declared in upper case
	var theURL string = "https://google.com"
	fmt.Println(theURL)

	/*
	 * Type Casting (Conversions)
	 */
	var f int = 42
	fmt.Printf("%v, %T\n", f, f)

	// Explicit conversion using "conversion functions"
	var g float32
	g = float32(f) // integer i converted to float32 j.

	fmt.Printf("%v, %T\n", g, g)

	// In type conversions, when converting from a
	// higher precision type (float32) to a lower precision type (int)
	// there may be loss of information.

	// More useful application of type conversion is
	// when we are converting numbers to strings

	var h string

	// NOTE:
	// The string() conversion method will convert the
	// given number into its unicode equivalent i.e.
	// 42 will be considered as '*', this is because in go
	// strings are simply an alias to a stream of bytes.
	i = 42
	h = string(i)

	fmt.Printf("%v, %T\n", h, h) // '*'

	// To solve this problem, we need to import
	// the string conversion package (strconv)

	h = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", h, h) // '42'

}
