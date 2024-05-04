package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	// Polymorphic Behavior
	// Instantiating the Interface
	// by any of its implementation that we need.
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// Type Conversion from the Interface to the Implementation
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// Run time error for an invalid conversion.
	/*
		bwc2 := wc.(io.Reader)
		fmt.Println(bwc2)
	*/

	// A better way to handle type conversions
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	/*
	 * Empty Interface concept
	 *
	 * It is useful when there are multiple things
	 * that we'll be working with but they are not type
	 * compatible with one another
	 *
	 * But since it does not have any methods of its own,
	 * we need to either:
	 * - use type conversion, or
	 * - use reflect package to figure out
	 *   the type of object we are dealing with.
	 */
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello Youtube listeners, this is a test"))
		wc.Close()
	}
	r2, ok2 := myObj.(io.Reader)
	if ok2 {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed")
	}

	/*
	 * Type Switch concept.
	 */
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("Don't know")
	}
}

/*
 * Declaring an Interface
 *
 * Here, we shall try to simulare I/O behavior.
 */
type Writer interface {
	// Method to write a slice of bytes
	// Returns the # bytes written or an error
	Write([]byte) (int, error)
}

/*
 * Declaring a structure which will implement our Interface.
 *
 * Unlike other languages, here we do not have any
 * special keyword like 'implements'. Because in Go,
 * we don't "explicitly" tell that we need to implement
 * an interface. In Go, we are going to "implicitly"
 * tell that we are going to implement the interface
 * by actually creating a method on our ConsoleWriter
 * that has a signature of our interface.
 */
type ConsoleWriter struct{}

// Encapsulate Write() method to ConsoleWriter of type struct.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
 * Not only struct, but any type that can have a method
 * associated with it can implement an Interface.
 */
type Incrementer interface {
	Increment() int
}

// Define the type alias for an integer.
type IntCounter int

// Encapsulate Increment() method into type IntCounter.
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

/*
 * Embedding interfaces within another Interface.
 *
 * So, the WriterCloser interface will be implemented if
 * the struct has all methods of both Writer and Closer
 * interfaces.
 */
type WriterCloser interface {
	Writer
	Closer
}

// Declaring the struct to which the
type BufferedWriterCloser struct {
	// some internal variables in the struct
	buffer *bytes.Buffer
}

// Implementing Writer interface here:
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 0 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Implementing Closer interface here
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// Just a Constructor function to initialize the struct
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
