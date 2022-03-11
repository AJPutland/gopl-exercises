# Chapter 1 notes

- Go is a compiled, static, strongly typed, imperitive language
- Multi-paradigm (OOP, procedural, functional)
- In the "C family" but has garbage collection
- Large standard library

- `go run` used to compile and run go source code 
- `go build` compiles go code once so you can run the exectuable

- go code organized into packages (one or more .go files in a single directory)
- package declaration at top, followed by packages to import;- imports m then declarations of functions, vatiables, constatn and types

- package main is special; it defines a standalone executable program, not a library. Within this, the main function is where execution begins
- semicolons allowed, but newlines act similarly

- strong stance on formatting, `gofmt` tool rewrites code into standard format, with go tool's `fmt` subcommand applying `gofmt` to all files in specified package / directory
- Additionally, `goimports` is a superset of `gofmt` and it manages addition / removal of import declarations.

- os package provides functions and other values for dealing with the operating system in a platform-independent fashion