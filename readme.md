# My Baby Steps with Go

This repo is the result of my first time learning Go basics sintax and operations.


## First things to Know

- Statically and StronglyTyped
- Is Compiled
- Built-in Concurrency
- Simplicity by design
- Garbage Collector

## Modules and Packages

**Package** is just a folder that contains a collection of go files and a Collection of packages is known as a **Module**

!["Illustration of  Go Module and packages"](./assets/go-module-package.png)



A new project in Go (which is just a new module) can be done by

```bash
go mod init baby_steps
```

This creates a new file called go.mod that contais the module name and the go version that you are using

!["Go Module Example"](./assets/go-mod.png)


To our first step lets create a directory `main/step1` the create a file called `main.go`, so the file structure look like this:
```text
.
├── go.mod
├── main
│   └── step1
│       └── main.go
```

Inside our `main.go` let's start coding.
Remember that a go file needs to be inside a package.

so the first line that we're gonna write is

```go
package main
```

and all the files inside this folder should have the same package name. Go docs give some package name guidelines [here](https://go.dev/blog/package-names).

This package name `main` is important because tells the go compiler the entrypoint of our applications is in this file.

With that, the go compiler will look for the `main` function. which is required for packages with the `main` name.

If we had another module called `potato` it would not need a `potato` function. This is only for `main` package.

Ok, lets create out main function.

## Functions

Functions in go are defined with the `func` keyword followed by the `name of the function` with its parameters beteween parenthesis. The block code of the function are inside curly-braces. 
This will be something like this:

```go
func main() {
    //function code goes here 
}
```
 ## The famous *Hello, World*
Lets import the `fmt` module to help us print something in the terminal and then use it in the `main` function to call `Println` command. You have to use all the modules that you import, because the go compiler will slap you face if you have unused imports

 ```go
import "fmt"

func main() {
    fmt.Println("IDon'tHaveAGirlfriend")
}
```

Ok then, lets run this function with:
```bash
# this will create a new file called main.go
go build main/step1/main.go
#now we just call the created file
./main
```

Or you can just:

```bash
#This will build and execute your code in one command
go run main/step1/main.go
```

And that is it, we've crated out First Go code. Congratulations, I think:

!["Go Execution"](./assets/build-run-go-file.png)



## Variable declarations
They are declared starting with `var` keyword, followed by the variable name and the type.
And just like the imports, the Go compiler scream if you have unused variables.

If you do set the value of the var when declarating one, Go will set default values.

If you are lazy, Go can infer the type and even let you remove the keyword `var` if you use `:=`

```go
func main() {
	var potatoQuantity int = 666
	fmt.Println(potatoQuantity)
    //666
    var anotherPotato int
	fmt.Println(potatoQuantity)
    // 0

    var grape = 0 // now this a int

    anotherGrape := 0 //  shortest delcaration possible

    ant, anotherAnt, uncleAnt := 0,2,4


}
```
## Types

### Some numbers

We have multiple `int` types.
Some like:
 `int8`, `int16`, `int32`, `int64`: this defines the size of the number that will be stored in the memory.
  Lets say that you want to use int8, which can store at max: number 127. If you add + 1 (now 128) the compiler will throw an overflow error. By default, if you put just `int` it will assume at least  32bits value, based on you system.

So, if you have something like this:
```go
package main

import "fmt"

func main() {
	var potatoQuantity int8 = 128 //exceeds the max value of the type
	fmt.Println(potatoQuantity)
}
```
You will receive the error
```text
 cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)
 ```

 But this occurs only in compilation time. At execution time, if you do it at execution time, you will see something like this
```go
func main() {
	var potatoQuantity int8 = 127
    // adding the value at runtime
	fmt.Println(potatoQuantity + 1)
    // >>> -128
}
```

So, be aware of this behaviour to avoid nights of debbuging. You`re welcome.

Another typw of int is a unsigned int: `uint` that will allow you to store all positive values possible for the int type. You canhave a comination like: uint, uint32, unint8 etc 

For floating numbers we have: `float32` and `float64`. as the in types, be aware of the data types because fi the value you stored in the type exceeds you will get the errors and weird behaviour that we saw previously.

As for arithmetic operations, you only can do it with var of the same type. You cannot do `float32` + `int32`. But you can do a type casting to solve like:
```go
func main() {
	var potatoQuantity int32 = 127
	var potatoPrice float32 = 12.2
	fmt.Println(potatoQuantity * int32(potatoPrice))
}
```

### strings
strings are stored in the `string` type 😛. And can declared with double quotes or back quotes.

Double quotes are for single lines, if you want to break a line you can add `\n`. like
```go
fmt.Println("Break the \nline")
```

With backquotes you can just hit enter and you just fine:
```go
fmt.Println(`Break the 
    line`
)
```

For concatenations just:
```go
fmt.Println("Concate"+"nated")
```

Be aware o how Go count the size of string. you can use the built-in `len`
```go
	fmt.Println(len("four"))
    // 4
	fmt.Println(len("fours"))
    // 5
	fmt.Println(len("©"))
    // 2
	fmt.Println(len("©©"))
    // 2
```

But this fucntion calculate it based on its byte size. Since Go uses the UTF-8 enconding. So any char outside of the vanilla ASCII table are stored with more than a single byte

To avoid this pain you can use the built-in: `unicode/utf8` and call `utf8.RuneCountInString`:

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("four"))
    //4
	fmt.Println(len("fours"))
    //5
	fmt.Println(utf8.RuneCountInString("©"))
    //1
	fmt.Println(utf8.RuneCountInString("©©"))
    //2
}
```

### booleans

just do 
```go
var boolean bool = true
```
### Constants
Almost everything we said before apply to constants but, you cannot change their value because they are constants in case you did`nt noticed.

```go
const dog = "Husky"
```
Its value must be defined at declaration time

## Functions and Control Structures

The first curly brace of a function must be in the same line of its initial definition.

```go
func main() {
	insult("Lucas")
}

func insult(your_name string) 
{
	fmt.Print("My name is " + your_name + ", and I use Windows")
}
//main/step1/main.go:12:1: syntax error: unexpected semicolon or newline before {
```

As we already saw, wu must use the `func` keyword followed by the name of the function and its parameters.

One Cool thing about the Go functions is that you can return multiple values delcaring like this:

```go
func main() {
	var result, remainder = divideAndReturnRemainder(8, 2)
	fmt.Printf("result: %v. remainder: %v", result, remainder)
}

func divideAndReturnRemainder(number1 int, number2 int) (int, int) {
	var result = number1 / number2
	var remainder = number1 % number2
	return result, remainder
}
```
*divideAndReturnRemainder* are the function that is returning multiple values.

### Error Handling

But every body is gangsta until we use this function to divide by zero
```go
divideAndReturnRemainder(8,0)
//panic: runtime error: integer divide by zeros
```

Of course we will receive an error, because, apparently math does not accept division by zero.

The Go way to handle this scenario is to have a return type error alongside with the no-error values.

`error` is a built-in type in Go and its default value is `nil` (that is the `null` of Go).

Since out function returns an error. The caller have to check if this error is `nil` or if it contains an error. Something like this:
```go
func main() {
	var result, remainder, err = divideAndReturnRemainder(8, 0)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("result: %v. remainder: %v", result, remainder)
}

func divideAndReturnRemainder(number1 int, number2 int) (int, int, error) {
	var err error
	if number2 == 0 {
		err = errors.New("math is not Mathing right now")
		return 0, 0, err
	}
	var result = number1 / number2
	var remainder = number1 % number2
	return result, remainder, err
}
```

This pattern of error handling is common in Go, so if you import another package, they probably will be doing like us. So make sure that you check if *err != nill*

### Switch ... case ...

Nothing much to add. just take a look at the structure:

```go
func main() {
    var day = 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day")
    }
}
```

### Arrays
As we already know, arrays are:
- Fixed Length
- Same Type
- Indexable (Start at index 0)
- Contiguous memory

The array below show how to declare one. And it obeys what we know about arrays.

```go
func main() {
	var intArr [3]int32

	fmt.Println(intArr)
	// [0 0 0]

	intArr[2] = 26
	fmt.Println(intArr)
	// [0 0 26]

	//Addr positions
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	//0xc0000120a0
	//0xc0000120a4
	//0xc0000120a8

}
```

### Slices
Under the hood slices are just arrays with some [salse](https://go.dev/doc/effective_go#slices).

Slices are delcared almost like arrays, except you can omit the array size. And you can add 'infinity' things of the same type, using `append(slice, elems)`

```go
func main() {
    var sliceNumbers []int
    sliceNumbers = append(sliceNumbers, 1)
    sliceNumbers = append(sliceNumbers, 2, 3, 4)
}
```
### Maps
