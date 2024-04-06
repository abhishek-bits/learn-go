package main

import (
	"fmt"
	"log"
)

// "net/http"

func main() {

	/************* DEFER ***************/

	/*
		fmt.Println("start")
		defer fmt.Println("middle")
		fmt.Println("end")
	*/

	/*
	 * Example scenario taken from Documentation
	 */
	/*
		res, err := http.Get("http://www.google.com/robots.txt")
		if err != nil {
			log.Fatal(err)
		}
		// We can delay the resource cleanup process.
		// This way we can associate the opening and
		// closing of resource right next to each other.
		defer res.Body.Close()
		robots, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
	*/

	/*
		a := "start"
		defer fmt.Println(a)
		a = "end"
	*/

	/***************** PANIC ****************/

	/*
		b, c := 1, 0
		ans := b / c
		fmt.Println(ans)
	*/

	// Applying panic for same situtation.
	/*
		fmt.Println("start")
		panic("something bad happended")
		fmt.Println("end")
	*/

	/*
	 * Example scenario from documentation
	 */
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Go!"))
		})
		// If the port 8080 is already running,
		// We can configure our application to panic
		// because the method ListenAndServe
		// would only return an error in this case.
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err.Error())
		}
	*/

	/*********** DEFER and PANIC ***************/

	// Defer is given a higher priority than panic().
	/*
		fmt.Println("start")
		defer fmt.Println("this was deferred")
		panic("something bad happended")
		fmt.Println("end")
	*/

	/***************** RECOVER *****************/

	/*
		fmt.Println("start")
		// anonymous function passed to defered.
		defer func() {
			// if the application didn't panic,
			// then err will be nil.
			if err := recover(); err != nil {
				// will print the err into the console.
				log.Println("Error: ", err)
			}
		}()
		panic("something bad happened")
		fmt.Println("end")
	*/

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

/**
 * Once, the function panics, it means we cannot move ahead.
 * the control passes to recover and then passed back to the
 * method that called panicker(), here main() after which the
 * following statements are executed.
 */
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// gracefully handle the panic situation.
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking!")
}
