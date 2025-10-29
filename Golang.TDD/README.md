
This repository documents my journey learning Go programming language using Test-Driven Development (TDD) methodology.

## What is TDD?

Test-Driven Development is a software development approach where you:
1. Write a failing test
2. Write the minimum code to make the test pass
3. Refactor the code while keeping the tests green


## Running Tests

```bash
go test -v
```

## Running benchmarks

For linux

```bash
go test -bench=.
```

For window

```bash
go test -bench="."
```

## Goals

- Learn Golang basics
- Practice TDD principles
- Build small projects incrementally
- Understand Go testing framework

## Discipline

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Go by examples](https://gobyexample.com/)