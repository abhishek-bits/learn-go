package main

import (
	"fmt"
	"runtime"
	"sync"
)

// WaitGroup is designed to synchronize multiple goroutines together.
// In this application, there are two goroutines that we care about:
// 1. Executed in the main() method
// 2. go routine that we spawn at line
var wg = sync.WaitGroup{}

var counter = 0

var m = sync.RWMutex{}

func main() {
	/*
	 * Named function
	 */
	// go sayHello()

	/*
	 * Anonymous function
	 */
	// var msg = "Hello"

	/*
		// Note that anonymous function will have access
		// to variables in the outer scope even in
		// a concurrent scenario.
		go func() {
			fmt.Println(msg)
		}()

		// RACE CONDITION Problem :-
		// Now, that there exists a relationship between the
		// goroutine and the variables in the outer scope there
		// exists a problem here.
		// The updated msg variable may reflect in the
		// goroutine but this is not what we need.
		msg = "Goodbye"
	*/

	/*
	 * The below code snippet is still NOT a best practise
	 * to avoid RACE CONDITION because we are using sleep()
	 * but this is not acceptable in production. What if anonymous
	 * method takes more time than what main() has to wait.
	 * Solution: WaitGroup
	 */
	/*
		go func(msg string) {
			fmt.Println(msg)
		}(msg)

		// No effect in the goroutine.
		msg = "Goodbye"

		// We shall stop the main() method to terminate
		// so that the child thread could execute
		// sayHello()
		time.Sleep(100 * time.Millisecond)
	*/

	/*
	 * Applying WaitGroup
	 */
	/*
		// How many WaitGroups do we have to wait for.
		wg.Add(1)

		go func(msg string) {
			fmt.Println(msg)
			// Tell the WaitGroup that execution is complete.
			// This basically decrements the counter for WaitGroup
			// that main() is waiting on.
			wg.Done()
		}(msg)

		msg = "Goodbye"

		// make sure main() does not terminate yet.
		wg.Wait()
	*/

	/*
	 * A more practical example of WaitGroup
	 */

	for i := 0; i < 10; i++ {
		wg.Add(2)

		// There is still no synchronization in these goroutines
		// that are running in parallel.
		// Solution 1: WaitGroups
		// Solution 2: Mutexes

		// We should always apply the Mutex Lock
		// before we run the goroutine

		m.RLock() // Apply Read Lock
		go sayHello()
		m.Lock() // Apply Write Lock
		go increment()
	}

	wg.Wait()

	// By default, go will give us as many # threads
	// as is the number of cores in the machine.
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// However, we can change this

	runtime.GOMAXPROCS(1) // A Single Threaded application
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	runtime.GOMAXPROCS(4)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

}

func sayHello() {
	// Applying a Read Lock within the goroutine
	// is still not a good idea
	// m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// Applying a Write Lock within the goroutine
	// is still not a good idea
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
