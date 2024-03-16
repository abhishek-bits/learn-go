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
