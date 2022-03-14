# Chapter 1 notes

## Go basics

- Go is a compiled, static, strongly typed, imperitive language
- Multi-paradigm (OOP, procedural, functional)
- In the "C family" but has garbage collection
- Large standard library

- `go run` used to compile and run go source code 
- `go build` compiles go code once so you can run the exectuable

- go code organized into packages (one or more .go files in a single directory)
- package declaration at top, followed by packages to import, then declarations of functions, vatiables, constatn and types
- imports can be individual declarations, or grouped by parenthases
- semicolons allowed, but newlines act similarly
- // comments, normally above package declaration (for main this is a few sentences to describe program as a whole)
- comments should be full sentences for the `godoc` tool

- package main is special; it defines a standalone executable program, not a library. Within this, the main function is where execution begins

- strong stance on formatting, `gofmt` tool rewrites code into standard format, with go tool's `fmt` subcommand applying `gofmt` to all files in specified package / directory
- Additionally, `goimports` is a superset of `gofmt` and it manages addition / removal of import declarations.

- Go provides the normal arithmetic and logical operators, with some semantically changing based on the types it is applied to (sting + string concatenates the strings)
- `i++` and `i--` are allowed (postfix only, no prefix)

- blank identifier `_` wherever syntax requires a varialbe but logic does not (also Go doesn't allow unsused local variables so this helps)

- strings are immutable

### Var declarations

- `var [name] [type]` format initializes the variable to it's zero value for it's type (0 for numeric types and "" for string types)
- Can also initialize the variable with a value using `=`
- multiple variable declarations / initializations allowed in one line using `,`
- `:=` is short variable declaration, declares one or more variables and assigns type based on the initializer values `[name] := [value]` (looks like a dynamic language) (cannot be used for package level decalrations)
- Even though there are lot's of ways to declare variables in Go usually short variable declaration  `i := ""` (explicit initialization) and normal `var i string` (implicit initialization to zero value) should be used

### For loop

- for loop is the ONLY loop in Go and has a number of formats (no parenthesis are used around for loops)
- `for initialization; condition; post {//statement}`
- initialization statement is executed before the loop starts, if it is present it must be a short variable declaration, an increment or assignment statement or a function call
- condition is a boolean statement which is evaulated each loop
- post is executed after the body of the loop each time
- ANY of these parts can be omitted, if no initialization or post you can not include the semi-colons

- `for condition {...}` is a while loop

- `for  {...}` is an infinite loop (can be broken by `break` or `return` statement)

- for loop can also iterate over a rance of values from a data type like a string or slice `for i, arg := range [value] {}` (i is index of arg in the range)


## FMT package

- fmt package provides I/O operations and string formatting

## OS package

- os package provides functions and other values for dealing with the operating system in a platform-independent fashion
- variable Args contains command-line arguments ( when referenced outside of the package os.Args)
- os.Args is a slice of strings with [0] being the name of the command itself

## Strings package 

- provides simple functions to manipulate strings
- Join function is a simpler and more efficient for concatenating strings than using manual string concatenation (more functional as well)