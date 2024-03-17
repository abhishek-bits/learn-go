package main

import (
	"fmt"
	"reflect"
)

// Declaring a STRUCT

// Struct name starts with the capital letter
// and all its fields are small letters
// (if only we don't want to export them)
// Capital letter would mean that this struct
// will be exported to the outside.
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max:"100"` // Tag applied to the field.
	Origin string
}

/*
 * Embedding in Structs
 *
 * Since, Go does not support Inheritance or
 * in general, is-a relationship.
 * It provides composition based approach.
 *
 * NOTE that Bird struct still has no relationship
 * with the Animal struct.
 */
type Bird struct {
	Animal   // Composition
	SpeedKPH float32
	CanFly   bool
}

func main() {

	// Initialize a struct variable.
	// It is HIGHLY RECOMMENDED to initialize
	// the struct this way as the named params
	// would make sure that the values are assigned
	// to the right field even if the ordering is
	// changed in the original struct.
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	// Print the struct
	fmt.Println(aDoctor)

	// Print a field of the struct
	// using dot operator
	fmt.Println(aDoctor.actorName)

	// Declaring a struct in a very precise form
	// But this way we can't use the struct anywhere else.
	bDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(bDoctor)

	// Unlike Slices and Maps,
	// In structs, copies refer to different memory locations.
	aDoctor2 := aDoctor
	aDoctor2.actorName = "Tom Baker"
	fmt.Println(aDoctor.actorName)  // John Pertwee
	fmt.Println(aDoctor2.actorName) // Tom Baker

	// However, just like Arrays, Structs can also have
	// pointer based memory address mapping.
	aDoctor3 := &aDoctor
	aDoctor3.actorName = "Tom Baker"
	fmt.Println(aDoctor.actorName)  // Tom Baker
	fmt.Println(aDoctor3.actorName) // Tom Baker

	// Application of Embedding to initialize
	// Bird object
	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Println(bird)

	// We can also use the more declarative syntax:
	bird2 := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		SpeedKPH: 30,
		CanFly:   false,
	}
	fmt.Println(bird2)

	// TAGS:-
	// We'll use reflect library to find out
	// the Tag for a particular field within the struct.
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
