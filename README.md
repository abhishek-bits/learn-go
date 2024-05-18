# Go (golang)

- Ref: [YouTube Tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU)

## Introduction

- Strong and statically typed.
- Excellent Community
- Key features
  - **Simplicity**
  - **Fast compile times** Go focuses on extremely fast compile times.
  - **Garbage collected** Go runtime manages garabage collection. Accross all previous versions, golang has evolved in regards to the amount of time the application has to pause due to garbage collection.
  - **Built-in concurrency** Go supports concurrent application development without use of any library.
  - **Compile to standalone binaries** When a Go application is compiled everything (including external libraries) is bundled into that single binary.

## Useful Resources

- [Official Website](https://go.dev)
- [Go Forum](https://forum.golangbridge.org/)
- [Go Playground](https://go.dev/play/)

## Fundamentals

- `main` package is the entry point of the go application.
- `import` is used to import required libraries. (In the example below, `fmt` (pronounced as "fimt") is a library to format strings).
- `func` is used to declare a function.
- `main()` method is going to contain the code that runs first.

```go
package main

import("fmt")

func main() {
  fmt.Println("Hello, World!")
}
```

## Installation

- [Windows Installation](https://www.youtube.com/watch?v=DFiXJKIF2ss)
- Linux Installation can be found wihin the same tutorial.

## Go Project Initialization and Configuration

To initialize a folder as a go module, run the following command:

```shell
go mod init [FOLDER_NAME]
```

This command will create a `go.mod` file within the folder which will help the compiler determine that this is a go module.

To run a file within this folder:

```shell
go run src/[FOLDER_NAME]/[file_name].go
```

The below command builds the entire go binary:

```shell
go build [FOLDER_NAME]
```

The below command puts the build binary into the bin folder:

```shell
go install [FOLDER_NAME]
```

These binary can be directly run using the path reference:

```shell
bin/[APP_NAME]
```

## Variables

### Variable Declaration

There are 3 ways to declare variables in go:

1. Declare first, initialize later: `var foo int`
2. Full Declaration Syntax: `var foo int = 42`
3. Shorthand Syntax (Automatic Inference): `foo := 42`

**NOTE**: All declared variables / methods must be used otherwise the code will not compile.

### Redeclaration and Shadowing

We can't redeclare variables in the same scope but we can re-assign the same variable to a different value (of the same type). We can however shadow (override) a variable that declared in the outerscope which may be of different type and/or value. In this case, the variable with the innermost scope takes the highest precedence.

### Visibility

- lower case first letter will have package level scope.
- upper case first letter will be used to export the functionality globally.
- there is no 'private' scope available in go.

### Naming Conventions

- Pascal or camelCase
  - Capitalize acronyms (HTTP, URL, etc)
- Variable names should be as short as reasonable
  - use longer names for variables with longer lives (like variable that will be exported).

### Type Conversions

- Go does not have any implicit type conversions, it provides _conversion methods_ to convert from one type to other.
- We must use `strconv` package for strings.

## Primitives

### Boolean type

- Values are `true` or `false`.
- Not an alias for other types (e.g. int)
- Zero value is false.

### Numeric types

#### Integers

##### Signed Integers

| Type  | Range                                       |
| ----- | ------------------------------------------- |
| int8  | -128 to 127                                 |
| int16 | -32768 to 32767                             |
| int32 | -2147483648 to 2147483647                   |
| int64 | -9223372036854775808 to 9223372036854775807 |

##### Unsigned Integers

| Type   | Range          |
| ------ | -------------- |
| uint8  | 0 - 255        |
| uint16 | 0 - 65535      |
| uint32 | 0 - 4294967295 |

We don't have any unsigned int of 64 bits but we do have a type called `byte`.

A `byte` is an alias for an 8 bit unsigned integer. The reason is that an unsigned 8 bit integer is very common, becuase that's what a lot of data streams are used to encode their data.

#### Floating point

Floating point numbers in go follow IEEE 754 standard.

| Type    | Range                   |
| ------- | ----------------------- |
| float32 | +-1.18E-38 - +-3.4E38   |
| float64 | +-2.23E-308 - +-1.8E308 |

Floating point numbers do not have modulus operator or any bitwise/bitshift operators.

#### Complex numbers

| Type       | Range               |
| ---------- | ------------------- |
| complex64  | float32 + float32 i |
| complex128 | float64 + float64 i |

Useful methods:

| Method      | Usage                                                        |
| ----------- | ------------------------------------------------------------ |
| `real()`    | To extract the real part of the complex number.              |
| `imag()`    | To extract the imaginary part of the complex number.         |
| `complex()` | To create the complex number given real and imaginary parts. |

### Text types

#### String

Strings in go:

- are collection of indexed **UTF-8** (or **ASCII**) characters.
- are basically an alias for **bytes**.
- are immutable meaning their values cannot be modified.

Given a string `s`, we can convert into a slice of bytes using `[]byte(s)` which will give us `[]uint8`. Similary we can use `string()` method to convert a `[]byte` (or `[]unit8`) into string.

**NOTE**: `uint8` is basically an alias for `byte`.

Converting strings to stream of bytes is useful because it makes working with strings way more efficient. This ability is really useful when we sending our string results out to an external service.

#### Runes

Runes in go:

- represents any **UTF-32** character.
- **true type alias** for an `int32`.

**NOTE**: Where strings can be converted back and forth between collection of bytes, runes are a true type alias meaning declaring an `int32` or a `rune` would mean the same thing.

Ref: [strings#Reader](https://pkg.go.dev/strings#Reader)

## Constants

- Immutable, but can be shadowed.
- Replaced by the compiler at compile time i.e. value must be calculable at compile time.

### Naming Convention

Named like variables.

- `PascalCase` for exported constants.
- `camelCase` for internal constants.

### Typed Constants

- Typed constants work like immutable variables
- They can interoperate only with same time (We have to apply conversion methods explicitly).

### Untyped Constants

- Untyped constants work like literals.
- They can interoperate with similar types becuase they belong to general types (`int` / `float`) hence provide more flexibility.

### Enumerated Constants

- Special symbol `iota` allows related constants to be created easily.
- `iota` starts with value `0` in each `const` block and increments by one. We can however skip some values for `iota` by using any offset and subsequent constants will then be assigned the next set of incremented values.
- We must watch out for constant values which match zero values for variables.

### Enumeration Expression

- Having constants defined using enumeration expression (such as bitshift) for a byte makes the operation really efficient.
- Makes the code clear and concise.

## Arrays and Slices

These are the first two collection types available in Go.

### Arrays

- Collection of items with the same type.
- Fixed size.
- Access via zero-based index.

#### Declaration

3 different ways:

1. Literal style `a := [3]int{1,2,3}`
2. Robust style `a := [...]int{1,2,3}`
3. Generic `var a [3]int`

#### Built-in functions

- `len()` returns the length of the array

#### Working with arrays

- Copies refer to different underlying data.

### Slices

- Backed by arrays.

#### Declaration

- Slice existing array or slice.
- Literal style: `a := []int{1,2,3}`
- Via make function
  - Create a slice with capacity and length = 10 `a := make([]int, 10)`.
  - Create a slice with length = 10 and capacity = 100 `a := make([]int, 10, 100)`.

#### Built-in functions

- `len()` returns the length of the slice.
- `cap()` returns the length of the underlying array.
- `append()` adds elements to the slice.
  - May cause expensive copy operation if underlying array is too small.

#### Working with slices

- Copies refer to same underlying array.

## Maps

- Collection of value types that are accessed via keys
- Created via literals or via `make` function.
- Members are accessed via `[key]` syntax
  - `myMap["key"] = "value"`
- Check for presence with `value, ok` form of result.
- Maps are reference types. Multiple assignments refer to the same underlying data.

## Structs

- Collection of disparate data types that describe a single concept.
- Keyed by named fields.
  - If Capital first letter then exported outside of the package.
  - If Lower first letter then scoped within the same package.
- Normally created as types, but we can create anonymous structs if we want to.
  - Common use case of short-lived anonymous struct would be to generate a JSON response to a web service call.
- Structs are value types. Multiple assignments refer to their own copies of the data.

### Embedding

There is no concept of inheritance in Go, but can use composition via concept called embedding. So, when we have to embed one struct into other, we just give the type of the Struct and we don't give it a field name. Go is then going to automatically interpret that for us and delegate any calls for fields or methods in the containing struct down to the embedded struct if the Top level struct does not contain a member with that name.

### Tags

Tags can be added to struct fields to describe field. This may be some set of rules that must be obeyed for that particular field.

## Conditionals

### If statments

Both the initializer as well as conditional expression together in a typical if statement in go:

```go
if pop, ok := statePopulations["Florida"]; ok {
  fmt.Println(pop)
}
```

We have to be very careful when working with mathematical expression using float types; Here is a more advanced method, that basically uses _errror function_ to compare float types:

```go
myNum = 0.123567899
if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
  fmt.Println("Same!")
} else {
  fmt.Println("Different")
}
```

### Switch statements

- Break statement in case block of a switch statement is already implied; However, we can apply `break` in case we want to branch out early.

- We can use `fallthrough` to bypass the default implied `break` behavior regardless of whether the expression given in the following case block executes to false.

- We do not need to have curly braces to write multiple statments in a specific case of a switch statement.

- Tag initialization with mathematical expression and cases with multiple tags:

```go
switch num := 2 + 3; num {
case 1, 5, 10:
  fmt.Println("one, filve, ten")
case 2, 4, 6:
  fmt.Println("two, four, six")
default:
  fmt.Println("another number")
}
```

- We can also have range checks for each particular case of a switch statement. Note that, in case of overlapping cases, the first case (from the TOP) is given the first preference.

```go
i := 1
switch {
case i <= 10:
  fmt.Println("less than or equal to ten")
  fallthrough
case i <= 20:
  fmt.Println("less than or equal to twenty")
// ...
}
```

- In addition to values, or mathemtical expressions, we can also create a switch statement to compare the type of a variable:

```go
var j interface{} = 1 // j holds an int type.
switch j.(type) {
case int:
  fmt.Println("j is an int")
case float64:
  fmt.Println("j is a float")
//...
}
```

## Looping

### Simple loops

- `for initializer; test; incrementer {}`
- `for test {}`
- `for {}`

### Exiting early

- `break`
- `continue`
- Labels (for exiting out of entire nested loops)

### Looping over collections

- arrays, slices, maps, strings, channels (will be discussed later)
- `for k, v := range collection {}`

## Defer, Panic and Recover

### `defer`

- Any function followed by `defer` keyword will execute only after all all the statements in the parent function are executed but before it actually returns.
- `defer` functions are executed in **LIFO** order.
- `defer` functions take the values of function parameteres exactly as they were declared before the defer call. Any updates to these variables later on are not considered by `defer`.
- Should not be used when using loops to read resources otherwise we can have memory issues.
- `defer` gets higher priority than `panic()` i.e. panic happends after all deferred statements are executed.

### `panic()`

- In Go, we generally don't consider things that could go wrong in an application to be _exceptional_ events or to be events that should cause the application to shutdown.
- panics referred to a state when the application gets into a state that it cannot recover from. Ex: division by zero.
- Functions will stop executing immediately at the point of panic.
  - Deferred functions will still fire.
- If nothing handles panic, program will exit.
- If there is a panicking situation we feel that we can recover from then we can use `recover()`.

### `recover()`

- `recover()` method simply states that the application is in a state that it can continue to execute.
- Used to recover from panics.
- Only useful in deferred functions. Why? Because deferred functions will execute even after a panic siutation.
- However, before we can proceed to the rest of the "good" part we need to go through the `recover()` logic so that the error can be acknowledged.
- Additionally, if we want our application to terminate immediately whenever any panic happens, once `recover()` logic is read, we can re-call the `panic()` method.

## Pointers

### Creating Pointers

- Pointer types use an asterisk (`*`) as a prefix to type pointed to
  - `*int*` - a pointer to an integer.
- Use the addressof operator (`&`) to get address of variable.

### De-Referencing Pointers

- Dereference a pointer by preceding with an asterisk (`*`).
- Complex types (e.g. structs) are automatically dereferenced.

### Create Pointers to Objects

- Can use the addressof operator (`&`) if value type already exists

```go
ms := myStruct{foo: 42}
p := &ms
```

- Use addressof operator before initializer

```go
p := &myStruct{foo: 42}
```

- Use the `new` keyword.
  - Can't initialize fields at the same time.

### Types with Internal Pointers

- All assignment operations in Go are copy operation except **Slices** and **Maps**.
- **Slices** and **Maps** contain internal pointers, so copies point to same underlying data.

## Functions

- The entry-point of the Go application is always within the `main` _package_ and withing that `main` _package_ we have to have a function called `main()` which takes no parameters and returns no values.

- The **lowercase first letter** of the method name means that the scope of the method is **private** whereas it will be **public** for **uppercase first letter**.

- Passing in pointers as arguments turn out to be very useful as they help us avoid passing very large struct's as a copy.

```go
func foo() {
  // ...
}
```

### Parameters

- Parameters as comma delimited list of variables and types

```go
func foo(bar string, baz int)
```

- Parameters of same type; list type once

```go
func foo(bar, baz int)
```

- When pointers are passed in, the function can change the value in the caller

  - This is **always true** for data of **slices** and **maps**.

- Use _variatic parameters_ to send the list of same types in
  - Must be the last parameter
  - Recevied as a slice
  - In the below example, baz is a variatic parameter which internally is a slice of integers

```go
func foo(bar string, baz ...int)
```

### Return values

- Single return values just list its type

```go
func foo() int {
  result := 0
  // ...
  return result
}
```

- Specify **multiple return values** surrounded by parantheses. In the example below, we return the result and an error as `(result type, error)`, which is very common idiom.

```go
func foo() (int, error)
```

- Can use **named return values**
  - Initializes returned variable.
  - Return using `return` keyword on its own.

```go
func foo() (int, error) (result int) {
  return
}
```

- Can return addresses of local variables
  - In this case, such varibles are automatically promoted from local memory (stack) to shared memory (heap).

### Anonymous functions

These are the functions that do not have any names. Local variables created inside these functions will not have be visible outside. But inner anonymous functions can make use of the variables available outside the scope.

- Immediately invoked

```go
func() {
  // ...
}()
```

- Assigned to a variable or passed as an argument to a function

```go
a := func() {
  // ...
}
a()
```

### Type Signature

This is basically a function signature, with no parameter names.

It is often most convenient when assigning anonymous functions to a variable using the `:=` syntax and the type is going to be automatically inferred.

However, if we are:

- using a **function as a parameter** to another function
- or **as a return value** from a function

then it is compulsory to provide the type signature.

```go
var f func(string, string, int) (int, error)
```

### Methods

Function that executes in context of a type.

General Syntax:

```go
func (g greeter) greet() {
  // ...
}
```

Receiver (here the type `greater`) can either be a value or a pointer.

- Value receiver gets copy of type.
  - Inefficient memory utilization: As for every single method invocation, entire object is copied into memory.
  - Changes made to the properties of the object will not reflect in the original object.
- Pointer receiver gets pointer to type.
  - Efficient memory utilization as only the memory location is passed on.

## Interfaces

- Declaring the interface:

```go
type Writer interface {
  Write([]byte)(int, error)
}
```

- Declaring the struct to embed this functionality:

```go
type ConsoleWriter struct {}
```

- Implementing the interface:

```go
func (cw ConsoleWriter) Write(dat []byte)(int, error) {
  n, err = fmt.Println(string(daata))
  return n, err
}
```

- We can embed multiple interfaces together into one interface:

```go
type Closer interface {
  Close() error
}

type WriterCloser interface {
  Writer
  Closer
}
```

- Type Conversion:

```go
var wc WriterCloser = NewBufferedWriterCloser()
bwc := wc.(*BufferedWriterCloser)
```

- Everything can be cast(ed) to an empty `interface`; even primitives.

```go
var i interface{} = 0 // int
```

### Implementing Interfaces

Here, we use the concept of **Method Sets**.

When we are implementing an interface:

- If we use a value type, then all the methods that implement the interface have to have a value receiver only. In other words, _The method set for a value type is the set of all methods that have value receivers_.

- If we are implenting the interface with a pointer then we need not worry about the about the type of receiver used. In other words, _The method set for a pointer is the sum of all the methods with value receivers and all of the methods with pointer receivers_.

### Best Practices

- **Use many, small interfaces**. The smaller interfaces are, the more useful and powerful they are gonna be.

  - Single method interfaces are some of the most powerful and flexible.
    - `io.Writer`
    - `io.Reader`
    - `interface{}` // empty interface

- **Don't export interfaces for types that will be consumed**.

  - If we are creating a library and someone else will be consuming the type then in this case it is better to go ahead and publish that concrete type itself.
  - Example: `db` package.

- **Do export interfaces for types that will be used by the package**.

  - If we are defining a type that we will be consuming in the package then in this case, it is better to export interfaces.
  - This way, whoever is using the package can create their own concrete types and implement the interfaces that we need.
  - Hence, we don't have to worry about the implementation; instead we are only concerned about the behaviors that they are exposing to us.

- **Design functions and methods to receive interfaces whenever possible**.
  - If we have the option to receive interfaces, in case we don't need the underlying data (variables) then we should go ahead and define the interfaces that we are expecting.
  - That way, we are making our methods and functions more flexible as we can have variours implementations to support different business objectives.

## GoRoutines

- Tool to implement concurrent programming.

### Creating GoRoutines

- Use `go` keyword in front of function call.
- When using anonymous functions, pass data as local variables.

### Synchronization

Often times concurrent threads would request and override same resource.

#### WaitGroups

Use `sync.WatiGroup` to wait for groups of goroutines to complete. Three methods available in `sync.WaitGroup` are:

- `add()` to inform the wait-groups that we have more goroutines to wait for.
- `wait()` to wait until all goroutines in wait-group are completed.
- `done()` to notify the wait-group that one of the goroutine is completed with its work.

#### Mutexes

Use `sync.Mutex` and `sync.RWMutex` to protect shared data access.

### Parallelism

- By default, Go will use # CPU threads equal to available cores.
- We can change the degree of parallelism with `runtime.GOMAXPROCS`.
- More threads can increase performance, but too many can slow it down.

### Best Practices

- Don't create goroutines in libraries
  - Let consumer control concurrency.
- When creating a goroutine, know how it will end.
  - Avoid subtle memory leaks.
- Before we deploy our application into production, find out with which degree of parallelism, our application runs most efficient.
- Check for **race conditions** at compile time.
  - Go compiler itself comes with this solution:

```shell
go run -race main.go
```

## Channels

Channels are a way to pass data among **goroutines** in a safe manner and prevent issues such as Race Condition and memory sharing problems that can cause issues in your applications that are very difficult to debug.

In short, Channels support Data Synchronization for **goroutines**.

- Channels are created using `chan` keyword.

```go
ch := make(chan int)
```

- Send message into the channel:

```go
ch <- val
```

- Receive message from channel:

```go
val := <-ch
```

- We can have multiple senders and receivers

- By default Channels are Bi-Directional, but we can restrict the flow of information:

  - Send-only: `chan <- int`
  - Receive-only: `<-chan int`

-

### Buffered Channels

- Block sender side till the receiver is available.
- Block receiver side till message is available.

```go
ch := make(chan int, 50)
```

When both the sender or receiver operate at a different frequency than the other side.

**Practical Application**: When the sender sends bulk requests in intervals of every 1 hr / day / months. Then we have to have a buffer at the recevier's end to store that much amount of data to be processed.

This way our channel that is receiving the data does not deadlock as it always has a buffer to store the data.

### For-Range loops with Channels

- Use to monitor channel and process messages as they arrive.
- Loop exits when channel is closed.

### Select statements with Channels

Allows goroutines to monitor several channels at once.

```go
for {
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
```

- **Blocks if all channels block**

  - If there is no message available in any channel than select statement is blocked by default and then when a message comes in it will be processed

- **If multiple channels receive value simultaneously, behavior is undefined**
  - Because of the highly parallel nature of many Go applications, we can get into situations where messages arrive on two channels at the same time virtually at the same time.
  - So one of those cases will get executed from the `select` block but we can't be sure which one's going to get executed.
  - So the ordering of statements in `select` really doesn't matter from the standpoint of how those conflicts are going to get resolved.
  - However, if we do want a non blocking `select` sttement, we can add the `default` case here. So, if there are no messages in any of the monitored channels, then the default case will go ahead and fire.
