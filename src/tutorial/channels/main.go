package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

// Creating a Log Channel
var logCh = make(chan logEntry, 50)

// Signal Only Channel
// Empty struct with zero memory allocation
var doneCh = make(chan struct{})

/*
 * We'll use WaitGroups to make sure our main()
 * routine waits until all our goroutines are
 * done.
 */
var wg = sync.WaitGroup{}

/*
 * We'll use Channels to synchronize the data
 * between the WaitGroups.
 */

func main() {

	/*
	 * Creating a channel.
	 *
	 * - using make()
	 * - use `chan` keyword to create a channel
	 * - specify the type of data to be passed amongst channels.
	 *
	 * Keep in mind that this makes the channel strongly typed.
	 * Hence, we can only send integers through the channel that
	 * we are creating.
	 */
	// ch := make(chan int)

	// handleWithChannel(ch)

	// An Asynchronous Scenario over the same channel.
	// This means only one message can be sent and consumed
	// at a single time.
	// NOTE:-
	// DEADLOCK situation will occur if # Write Go-Routines
	// are greater than the # Read goroutines.
	// Because the statement: `ch <- 42` will
	// pause the execution of the go-routine
	// Only when there is someone to cosume this
	// value only then the execution continues.
	/*
		for j := 0; j < 5; j++ {
			handleWithChannel(ch)
		}
	*/

	// handleWithChannel2(ch)

	// handleWithChannel3(ch)

	/*
	 * Handling Deadlocks with a Buffered Channel
	 */
	// The second parameter in the make() method here
	// is going to tell Go that this channel has got
	// an internal data store that can store upto 50 integers.
	buffCh := make(chan int, 50)

	handleWithChannel4(buffCh)

	wg.Wait()

	// Create a go routine that listens to the Log Channel
	go logger()

	// Close the channel after main() is done execution.
	/*defer func() {
		close(logCh)
	}()*/

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{}
}

func handleWithChannel(ch chan int) {

	wg.Add(2)

	/*
	 * Receiving Go-Routine
	 * This will only receive the data through the channel.
	 */
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	/*
	 * Sending Go-Routine
	 * This will only send the data through the channel.
	 */
	go func() {
		// ch <- 42

		// Lets day we pass a variable to the channel
		// instead of a hard coded value.
		// Even though the variable is updated any where else
		// The value passed to the channel does not change.
		i := 42
		ch <- i
		i = 27

		wg.Done()
	}()
}

// Both the Go Routines are acting as
// Readers and Writers
func handleWithChannel2(ch chan int) {

	wg.Add(2)

	go func() {
		i := <-ch // Reader
		fmt.Println(i)
		ch <- 27 // Writer
		wg.Done()
	}()

	go func() {
		ch <- 42          // Writer
		fmt.Println(<-ch) // Reader
		wg.Done()
	}()
}

// Sometimes we want to really restrict the
// Go Routines to be either Read only or
// a Write only.
// We'll do that by specifying the arrow
// direction for the channel in the parameter
// of the Goroutine function
func handleWithChannel3(ch chan int) {

	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

}

// Deadlock problem solved using Buffered Channel
func handleWithChannel4(buffCh chan int) {

	wg.Add(2)

	go func(ch <-chan int) {
		// When we work with Buffered Channels,
		// We know that at a time we may have to process
		// multiple data items together.
		// Thus the below approach has to be improved:
		/*
			i := <-ch
			fmt.Println(i)
		*/

		// For-Range syntax to range over a Channel
		// This way Go automatically detects when the Channel
		// is closed so that it can execute.
		for i := range ch {
			fmt.Println(i)
		}

		// Manually handling if Channel is closed
		/*for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}*/

		wg.Done()
	}(buffCh)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27 // May cause data Loss if not proerly handled at the reciever

		// Handling Deadlock at the Receiver's end
		// Let the channel know that the Sender is done
		// sending data into the channel
		close(ch)

		// NOTE: When we close the channel, we are
		// not allowed to send any data into the channel.

		wg.Done()
	}(buffCh)

}

func logger() {

	// With thos for-range logger code,
	// When do we know the logger() routine has to terminate ?
	// Acutally when the main() method would terminate,
	// logger() routine would be forcibly shutdown.
	// While this is acceptable is some situations,
	// in other, we may need graceful termination.
	/*for entry := range logCh {
		fmt.Printf(
			"%v - [%v]%v\n",
			entry.time.Format("2006-01-02T15:04:05"),
			entry.severity,
			entry.message)
	}*/

	// Gracefully terminating the logger code.
	// Application of Signal-Only Channel and Select statement
	// Using select statement to improve the Logger code
	// to solve this problem:
	for {
		// NOTE: Do not have a default case
		select {
		case entry := <-logCh:
			fmt.Printf(
				"%v - [%v]%v\n",
				entry.time.Format("2006-01-02T15:04:05"),
				entry.severity,
				entry.message)
		case <-doneCh:
			break
		}
	}
}
