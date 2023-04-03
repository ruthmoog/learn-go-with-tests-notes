# Learn Go With Tests

[🔖](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#keep-going-more-requirements)

## Go Fundamentals
### [Hello, World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

*General testing tips...*

- modules command for running `go test` is `go mod init SOMENAME`
- use subtests to triangulate on a general behaviour

*Testing syntax tips...*

- test files are suffixed with SOMENAME`_test.go`
- the test func starts with `Test`SomeName
- `t` is the hook to the testing framework with methods(?) like `t.Fail()` and `t.Errorf(...)`
- `*testing.TB` is an interface combo of `T` test and `B` benchmark
- `t.Helper` required for safe debugging with helper methods

*General Go tips...*

- declare vars with `:=`
- build a string with placeholder values `%q` (q: double quotes)
- define constants with `const`

**TDD Cycle**

1. Write a test -> Make compiler pass 
2. Run the test -> Check the error message
3. Write enough to make it pass
4. Refactor
