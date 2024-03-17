package main

import "fmt"

func main() {

	// Declaring a Map using Literal Syntax
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	// NOTE:
	// A Slice is not a valid key type.
	// But Array can be used as a key.

	// a := map[[]int]string{} // Error, Invalid Key Type.

	a := map[[3]int]string{} // Okay

	fmt.Println(a) // map[]

	// Declaring a Map using make() function
	statePopulations2 := make(map[string]int)

	fmt.Println(statePopulations2)

	// Accessing the value of keys in Map
	fmt.Println(statePopulations["Ohio"])

	// Applying some invalid key
	fmt.Println(statePopulations["Georgia"]) // 0

	// Adding a new key-value pair to Map
	statePopulations["Georgia"] = 10310371

	fmt.Println(statePopulations["Georgia"])

	// NOTE that the ordering of elements in a map
	// is not guaranteed.

	// Delete a key-value pair from the Map using
	// the built-in delete() function.
	delete(statePopulations, "Georgia")

	fmt.Println(statePopulations["Georgia"]) // 0

	// Getting 0 could mean multiple things here.
	// In order to handle such ambiguous result,
	// we can do as follows:
	pop, ok := statePopulations["Georgia"]

	// Now, if variable 'ok; is false that means,
	// that the key does not exist in Map
	fmt.Println(pop, ok)

	// Also, if we are not interested in the value at all,
	// but we need to know whether a key exists in the Map,
	// we can do as follows:

	_, ok2 := statePopulations["Georgia"]

	fmt.Println(ok2)

	// We can find out the number of key-value pairs
	// in a Map by using len() method
	fmt.Println(len(statePopulations))

	// Just like Slices, Maps are also reference based
	// Thus, copying Maps would mean pointing to the
	// same memory.
	sp := statePopulations

	delete(sp, "Ohio")

	fmt.Println(statePopulations)
	fmt.Println(sp)
}
