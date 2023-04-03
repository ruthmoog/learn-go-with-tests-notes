# Learn Go With Tests

## Go Fundamentals
### [Hello, World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

Modules command for running `go test` is `go mod init SOMENAME`.

- test files are suffixed with SOMENAME`_test.go`
- the test func starts with `Test`SomeName
- `t` is the hook to the testing framework with methods(?) like `t.Fail()` and `t.Errorf(...)`

- declare vars with `:=`
- build a string with placeholder values `%q` (q: double quotes)
- define constants with `const`
- 