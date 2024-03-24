package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Unlike under languages, Go does not
	// have any comma operator that basically
	// separates out two expressions, so the below
	// code is not allowed
	/*
		for i := 0, j := 0; i < 5; i++, j++ {
			fmt.Println(i)
		}
	*/

	// We can however do the same thing as shown:
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// In case, when 'i' is scoped outside
	// of the for loop:
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// Also, in case when the counter is
	// updated somewhere within the body of
	// the for loop
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Here is an infinite loop which breaks only
	// once it hits some complex logic.
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// Using continue to skip an iteration
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}

	// Break out of loops in nested-loop case
	// using labels.
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	/* Working with Collections */

	s := []int{1, 2, 3}

	// For-Range loop is a special form of for loop
	for key, value := range s {
		fmt.Println(key, value)
	}

	// Same thing for a string
	str := "Hello Go!"
	for k, v := range str {
		fmt.Println(k, string(v))
	}

	// Same thing for Maps
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	for key, value := range statePopulations {
		fmt.Println(key, value)
	}

	// What if we only want keys ?
	for key := range statePopulations {
		fmt.Println(key)
	}

	// What if we only want values ?
	for _, value := range statePopulations {
		fmt.Println(value)
	}

}
