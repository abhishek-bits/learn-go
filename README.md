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
