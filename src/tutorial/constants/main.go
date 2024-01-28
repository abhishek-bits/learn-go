package main

import (
	"fmt"
)

// Just like primitives,
// constants can also be shadowed
const b int16 = 27

// iota is a special symbol (untyped int)
// It is a counter which is used when
// we are creating enumerated constants.
// Once the first value in this block is assigned to iota
// the other constants will be inferred to have the
// respective value of the iota (the counter) value.
const (
	d = iota
	e // e = iota
	f // f = iota
)

// The value of iota is only scoped to a
// specific constant block
// So, in this constant block the counter
// will assign the value starting from 0.
const (
	d2 = iota
)

// Practical usage of Enumerated constants
const (
	_ = iota
	catSpecialist
	dogSpecialist = iota + 10 // jumping by some offset.
	snakeSpecialist
)

// Common operations on constants are
// still allowed in go during initialization.
// - arithmetic
// - logical
// - bitwise
// - bitshift
const (
	_ = iota + 5
	catSpecialist2
)

// Bitshifting would help us avoid calling "math" package
// when all we want to do is to raise things to the power of two
const (
	_  = iota             // ignore first value by assigning to blank identifier.
	KB = 1 << (10 * iota) // 1 << (10 * 1) = 1 << 10 = 2^10
	MB                    // 1 << (10 * 2) = 1 << 20 = 2^20
	GB                    // 1 << (10 * 3) = 1 << 30 = 2^30
	TB
	PB
	EB
	ZB
	YB
)

// Bitshifting can also be used to set boolean flags
// inside of a single byte.
const (
	isAdmin            = 1 << iota // 0000 0001 (first bit is set)
	isHeadquaters                  // 0000 0010 (second bit is set)
	canSeeFinancials               // 0000 0100 (third bit is set)
	canSeeAfrica                   // 0000 1000
	canSeeAsia                     // 0001 0000
	canSeeEurope                   // 0010 0000
	canSeeNorthAmerica             // 0100 0000
	canSeeSouthAmerica             // 1000 0000
)

func main() {
	/*
	 * Naming convention
	 */

	// Declaring a constant that will be exported.
	// Here, we will keep the first character capitalized.
	const MyConst = 1000

	// Declaring a constant that is private to this module.
	const myConst = 10

	const myFloatConst = 10.0

	// Above are all untyped constants
	// Let us now declare a typed constant.
	const a int = 32
	fmt.Printf("%v, %T\n", a, a)

	// We cannot assign any such expression to
	// a constant variable whose result is not
	// known at compile time
	// const b float64 = math.Sin(1.57) // compiler error.

	// Constants can be made up of any primitive value.

	// Shadowing constant 'b'
	fmt.Printf("%v, %T\n", b, b)
	// Not only can we change the value of the constant
	// But we can also change its type.
	const b int32 = 42
	fmt.Printf("%v, %T\n", b, b)

	// Constants can be used in arithmetic operations
	// along with any primitive variables.
	// But, the type of the constant and these variables
	// should match otherwise compiler error occurs.
	var c int8 = 3
	// fmt.Printf("%v, %T\n", b + c, b + c) // error
	// Either we can use conversion method over variable
	fmt.Printf("%v, %T\n", b+int32(c), b+int32(c))
	// Or we can use conversion method over constant.
	fmt.Printf("%v, %T\n", int8(b)+c, int8(b)+c)

	// In case of an untyped constant, the compiler would
	// do the same job for us thus avoiding compiler error.
	// In this case, the compiler replaces the constant
	// with its value (here 10).
	fmt.Printf("%v, %T\n", myConst+c, myConst+c)

	/*
	 * Enumerated constants
	 */

	// iota will change it's value as and when
	// the constants are assigned.
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)

	fmt.Printf("%v, %T\n", d2, d2)

	var specialistType int = catSpecialist
	fmt.Println(specialistType == catSpecialist)

	// But what if the specialistType was not assigned
	// any constant in the initialization, in that case,
	// the default value is assumed to be zero which in
	// the enumerated constants still mean catSpecialist
	// Solution: Add a new enumerated constant called
	// errorSpecialist to hold the value 0.

	// But this variable 'errorSpecialist' is unnecessary
	// using the memory.
	// Can we do better?
	// Yes, any value that we don't want to use from iota
	// we can assign it to an underscore '_'.
	// It means that we are throwing away this value
	// of iota.

	fmt.Println(catSpecialist2)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// Utilizing enumerated constants initialzed via bit-shifting.
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles) // 100101

	// We can also apply bit-masking.
	fmt.Printf("Has Admin Rights? %v\n", isAdmin&roles == isAdmin)
}
