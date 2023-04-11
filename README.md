# Learn Go With Tests

[🔖](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#:~:text=change%20SumAll%20to-,SumAllTails,-%2C%20where%20it%20will)

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
    } ```
    
*General Go tips...*

- command for running `go test` is `go mod init SOMENAME`
- declare vars with `:=`
- build a string with placeholder values `%q` (q: double quotes) [see fmt package](https://pkg.go.dev/fmt)
- define constants with `const`
- Use a _named return value_ and avoid explicitly declaring it in the method body
  - it will be assigned the empty value of the type, eg `""` or `0`
  - `return` will return the _named return value_, no need to explicitly return the var
  - it will be automatically added to the Go Doc
- public funcs are announced in `PascalCase` and private funcs are whispered in `camelCase`
- Go does not allow method overloading;
  - you can have methods with the same name in their own packages
  - you can define methods on the type instead, eg shapes' _Rectangle struct_

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
mySlice1 := []int{1, 2, 3, 4, 5}
// or like this
mySlice2 := make([]int, 5)
```
**Arrays...**

- the `%v` placeholder works well for arrays
- use `array[index]` to get the value at that index
- `range` lets you use syntax like for number in numbers
  - `for _, number := range numbers { //do something }}`
- arrays of different sizes are **different types**

**Slices**

- you can't use equality operators (`==` etc), you can use `reflect.DeepEqual` instead
  - DeepEquals isn't type safe so will compile even when types are mismatched
  - (but your test should fail...!)
- use `slice[index]` to get the value at that index 
  - or, set a value at that index `slice[index] = 19`
  - setting a value to an index that doesn't exist yet will throw a `runtime error`
  - so, use `append` instead, `slice = append(slice, 19)
- `slice[low:high]` slices the contents of the slice into a spanning chunk
  - leave blank on one of the sides of the `:` to capture everything to that side of it

## Structs, methods & interfaces

**Structs**

- A struct is a named collection of fields where you can store data.
- Can use a struct to create a simple type, eg _rectangle_
- Declare a struct like this,
  - ```go
    type Rectangle struct {
      Width  float64
      Height float64
    }
  ``` 
- Or an anonymous struct, for example,
  - ```go
    // We are declaring a slice of structs by using []struct with two fields
    areaTests := []struct{
      shape Shape
      want float64
    }{
      {Rectangle{12, 6}, 72.0},
      {Circle{10}, 314.1592653589793},
    }
  ```
    
**Interfaces**

- interface resolution is implicit
- If the type you pass in matches what the interface is asking for, it will compile
- Declare an interface like,
  - ```Go
    type Shape interface {
    Area() float64
    }```
- using interfaces to declare only **what you need** is important design