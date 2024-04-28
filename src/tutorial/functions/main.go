package main

import "fmt"

func main() {
	// sayMessage("Hello Go!")

	// for i := 0; i < 5; i++ {
	// 	sayMessage2("Hello Go!", i)
	// }

	// sayGreeting("Hello", "Stacey")

	greeting := "Hello"
	name := "Stacey"
	sayGreeting2(&greeting, &name)
	fmt.Println(name)

	sum(1, 2, 3, 4, 5)

	sPtr := sum3(1, 2, 3, 4, 5)
	fmt.Println("Result: ", *sPtr)

	sum := sum4(1, 2, 3, 4, 5)
	fmt.Println("Result: ", sum)

	// Receiving multiple values from a function call.
	d, err := divide2(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Division result: ", d)

	/*
	 * Anonymous function
	 */

	func() {
		fmt.Println("Hello Anonymous!")
	}()

	for i := 0; i < 5; i++ {
		// Inner functions can take advantage of
		// variables that are available in the outer scope.
		func() {
			fmt.Println(i)
		}()
	}

	// The above inner method works only in a synchronous scenario
	// But if this inner method is running asynchronously we can get
	// an abnormal behavior. For example: the couter variable may
	// increment further while the inner method has not yet executed.
	// Solution:
	for i := 0; i < 5; i++ {
		func(i int) { // Clearly define the variable usage
			fmt.Println(i)
		}(i) // Pass that variable here
	}

	// Variable representing an anonymous function
	// Note that in this case the function cannot be re-used
	// in other parts of the application.
	f := func() { // var f func() = func() {
		fmt.Println("Hello variable func!")
	}

	f() // Executing the variable.

	// More complicated scenario for the above situation.
	var divide func(float64, float64) (float64, error) = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	res, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	/*
	 * Working with methods
	 */
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

	fmt.Println("The name is: ", g.name)

	g.greet2()

	fmt.Println("The name is: ", g.name)

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}

/*
 * For methods with the same type of parameters
 * we can use the comma separator to list out the params
 * and mention the type at the end.
 */
// func sayGreeting(greeting string, name string) {
// 	fmt.Println(greeting, name)
// }
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

/*
 * Function with pointer arguments
 */
func sayGreeting2(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted" // manipulate the value
	fmt.Println(*name)
}

/*
 * Function with a variatic parameter.
 *
 * - A function can have only one variatic parameter.
 * - A variatic parameter should be the last parameter.
 */
// Here, go compiler would take all the last (comma separated)
// integer values from the caller and convert them into a
// slice of integers.
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

/*
 * Adding return type to a function
 */
func sum2(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

/*
 * Ability to return a local variable as a pointer.
 *
 * In such a case, the value of the local variable is
 * copied from the local execution stack to the heap.
 */
func sum3(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

/*
 * Named return values.
 *
 * Here, result variable will be available in the
 * scope of our sum4 function and its value will be
 * implicitly returned.
 */
func sum4(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	// go compiler automatically interprets that
	// result variable is to be returned.
	return
}

/*
 * Function returning multiple values
 */
// This divide method does not handle the case
// when denominator is zero. We can however use
// if statement and panic in case if b is zero,
// but in general we want to avoid panicking.
func divide(a, b float64) float64 {
	return a / b
}
func divide2(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

/*
 * Method:
 *
 * Declaration looks similar to a function but here,
 * we are embedding this functionality into the
 * struct called 'greeter'
 */

// NOTE: Here, we are getting a copy of the struct object
// so any changes done to the properties of this struct
// within this method will not have any effect on the
// object visible in the caller.
// Tradeoff: If the object passed is very large it could
// consume a lot of heap space every single time this
// method is invoked.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// Now, here we have passed a pointer to the method
// This means that we now have efficient memory utilization
// Also, any changes done to the properties of this object
// will also reflect in the caller.
// Due to the implicit de-referencing of pointers in
// the go language.
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
