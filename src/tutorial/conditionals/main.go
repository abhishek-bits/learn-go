package main

import (
	"fmt"
	"math"
)

func main() {

	a := true

	// In Go, it is mandatory to use curly braces
	// even in a single line if block
	if a {
		fmt.Println("The test is true")
	}

	// Example below shows the usage of if block
	// in the initializer syntax.
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	// Both the 'pop' and 'ok' variables are defined
	// within the scope of the if statement.
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	// Variable 'pop' cannot be used outside the if block
	// fmt.Println(pop) // Error!

	// Comparison operators
	number := 50
	guess := 70
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	// Logical Operators
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		// code here...
	}
	fmt.Println(!true) // false

	// Short-Circuiting exists as in other languages.

	// if, else if, else
	if guess < 1 {
		// failure code here...
	} else if guess > 100 {
		// failure code here...
	} else {
		// actual code here...
	}

	// Below example works for 0.1 but does not
	// work for other such as 0.123.
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("Same!")
	} else {
		fmt.Println("Different")
	}
	// Reason is that floating point calculations give us
	// different results than numeric types.
	// Solution ? Go for absolute value
	myNum = 0.123567899
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("Same!")
	} else {
		fmt.Println("Different")
	}

	/************************************************/

	// NOTE:
	// 1. Unlike other programming languages, here, we do not need to put break statement. (It is implied in switch statement)
	// 2. We do not any curly braces for case block, we still can add any number of statements.

	/* SWITCHES with TAGS */

	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not one or two")
	}

	// Multiple cases is also possible:
	switch num {
	case 1, 5, 10:
		fmt.Println("one, filve, ten")
	case 2, 4, 6:
		fmt.Println("two, four, six")
	default:
		fmt.Println("another number")
	}

	// switch case can also have mathematical expressions
	switch num := 2 + 3; num {
	case 1, 5, 10:
		fmt.Println("one, filve, ten")
	case 2, 4, 6:
		fmt.Println("two, four, six")
	default:
		fmt.Println("another number")
	}

	/* SWITCHES with NO TAGS */

	// switch case also supports tag-less syntax
	// here, the evaluation is in the order of
	// TOP -> BOTTOM
	i := 10
	switch {
	case i <= 10:
		fmt.Println("i <= 10")
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// In case if we want the next case to be executed
	// if the previous one was caught, then we can use
	// fallthrough. However, in this case, it is the programmer's
	// responsibility to take care of flow of execution.
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Switch case can also be used to
	// check for the type of the variable.

	// In Go, interface type can take any value
	// of all types available.
	var j interface{}
	// j = 1 		// j will now be an int type.
	// j = "1" 		// j will now be a string type.
	// j = [3]int{} // j will now be an array of INT TYPE and of SIZE THREE.
	j = [2]string{} // j will now be an array of STRING TYPE and of SIZE TWO.

	// NOTE that, in go an array is not equilavent to
	// every other array, that other array should be
	// of same type and length.

	// Below statement tells go compiler to pull
	// the actual underlying type of that interface
	// and use that for whatever we are doing next.
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float")
	case string:
		fmt.Println("j is string")
	case [2]int:
		fmt.Println("j is [2]int")
	case [2]string:
		fmt.Println("j is [2]string")
	default:
		fmt.Println("j is another type")
	}

	// In some cases, we might want to leave the case early.
	// In that case, we can use the break statement.
	j = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
		break
		fmt.Println("j was an int")
	default:
		fmt.Println("none")
	}
}
