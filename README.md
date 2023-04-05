# Learn Go With Tests

[🔖](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration)

## Go Fundamentals
### [Hello, World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

*Testing syntax tips...*

- test files are suffixed with _SOMENAME_`_test.go`
- the test func starts with `Test`_SomeName_
- `t` is the hook to the testing framework with methods like `t.Fail()` and `t.Errorf(...)`
  - the testing framework gives you a `t` you can call methods on
- `*testing.TB` is an interface combo of `T` test and `B` benchmark
- `t.Helper` required for safe debugging with helper methods

*General testing tips...*

- use subtests to triangulate on a general behaviour
- you can include examples in test files
  - ```go
    func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
    }
  ```

*General Go tips...*

- command for running `go test` is `go mod init SOMENAME`
- declare vars with `:=`
- build a string with placeholder values `%q` (q: double quotes)
- define constants with `const`
- Use a _named return value_ and avoid explicitly declaring it in the method body
  - it will be assigned the empty value of the type, eg `""` or `0`
  - `return` will return the _named return value_, no need to explicitly return the var
  - it will be automatically added to the Go Doc
- public funcs are announced in `PascalCase` and private funcs are whispered in `camelCase`

**TDD Cycle**

1. Write a test -> Make compiler pass 
2. Run the test -> Check the error message
3. Write enough to make it pass
4. Refactor

### Benchmarking

`b testing.B` -> `b.N` executes the benchmark code n times & measures the time is takes.
run from the containing directory:

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("❤️")
	}
}
```
```bash
% go test -bench=.

goos: darwin
goarch: arm64
pkg: hello/iteration
BenchmarkRepeat-8       14493148                82.71 ns/op
PASS
ok      hello/iteration 2.430s
```

### Coverage

Go has a built-in coverage tool, you can run in the terminal - really helpful to check you're not losing coverage when deleting tests

```bash
% go test -cover    
PASS
        hello/arraysandslices   coverage: 100.0% of statements
ok      hello/arraysandslices   0.129s
```

### Arrays

Arrays have a fixed capacity which is defined when declared:
```go
// that could look like this
myArray1 := [5]int{1, 2, 3, 4, 5}
// or like this
myArray2 := [...]int{1, 2, 3, 4, 5}

// use slices for collections of variable length
mySlice := []int{1, 2, 3, 4, 5}
```
**Arrays...**

- the `%v` placeholder works well for arrays
- use `array[index]` to get the value at that index
- `range` lets you use syntax like for number in numbers
  - `for _, number := range numbers { //do something }}`
- arrays of different sizes are **different types**