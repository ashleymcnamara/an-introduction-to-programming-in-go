# [An introduction to programming in Go](http://www.golang-book.com)

## Chapter 1 - Getting Started
* go help
```
Go is a tool for managing Go source code.

Usage:

        go command [arguments]

The commands are:

    build       compile packages and dependencies
    clean       remove object files
    env         print Go environment information
    fix         run go tool fix on packages
    fmt         run gofmt on package sources
    generate    generate Go files by processing source
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

    c           calling between Go and C
    filetype    file types
    gopath      GOPATH environment variable
    importpath  import path syntax
    packages    description of package lists
    testflag    description of testing flags
    testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```

* godoc fmt Println
```
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
```
* go run program.go

## Chapter 2 - Your first program

## Chapter 3 - Types

### Integers

**Unsigned integer** types:  
- uint8  
- uint16   
- uint32  
- uint64

**Signed integer** types:
- int8
- int16
- int32
- int64

Alias types: 
- **byte** is the same as uint8
- **rune** is the same as int32

Machine dependent integer types:
- uint, int, uintptr
They are machine dependent because their size depends on the type of architecture you are using

1 byte = 8 bits
1024 bytes = 1 kilobyte
1024 kilobytes = 1 megabyte

### Floating point numbers
- They are inexact
- Like integers they have a certain size (32 or 64 bits)
- Using a larger sized floating point number increases its precision (how many digits it can represent)
- In addition to number there are several other values which can be represented: 
    - "not a number" NaN for things 0/0
    - positive and negative infinity

Floating point types:
- **float32** (single precision)
- **float64** (double precision)

### Complex numbers
- **complex64**
- **complex128**

### Operators
+ addition
- subtraction
* multiplication
/ division
% remainder

### Strings
- A string is a sequence of characters with a definite length used to represent text.
- Go strings are made of individual bytes, usually one for each character.
(Characters from other languages like Chinese are represented by more than one byte)
- String literals can be created using double quotes `"Hello World"` or back ticks ``Hello World``. The difference between thse is that double quoted strings cannot contain newlines and they allow special escape sequences, like \n, which gets replaced with a newline, and \t gets replaced with a tab character.

### Booleans
- A boolean value is a special 1 bit integer type used to represent true and false (or on and off)
- Three logical operators are used with boolean values
&& and
|| or
! not

### Chapter 4 - Variables
* Defining multiple variables: Use the keyword var (or const) followed by parentheses with each variable on its own line.
```
var (
    a = 5
    b = 10
    c = 15
  )
```

### Chapter 5 Control Structures

#### Loops
Other programming languages have a lot of different types of loops (while, do, until, foreach, ...) but Go has only has one (i.e. for loop) that can be used in a variety of different ways.

### Chapter 6 Arrays, Slices and Maps
#### Arrays
* An array is a numbered sequence of elements of a single type with a fixed length
* Arrays are indexed starting from 0
* Major issue with arrays: their length is fixed and part of the array's type name.  
In order to remove the last item, we actually had to change the type as well.   
Go's solution to this problem is to use a different type: slices.
#### Slices
* A slice is a segment of an array. 
* Like arrays slices are indexable and have a length. 
* Unlike arrays this length is allowed to change.
* If you want to create a slice you should use the built-in make function
* Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. 
* The make function also allows a 3rd parameter which represents the capacity of the underlying array which the slice points to
* Another way to create slices is to use the [low : high] expression:  
low is the index of where to start the slice and high is the index where to end it (but not including the index itself).
* For convenience we are also allowed to omit low, high or even both low and high. ```
arr[0:] is the same as arr[0:len(arr)],
arr[:5] is the same as arr[0:5] and  
arr[:] is the same as arr[0:len(arr)]
```
* Slice functions: append and copy
#### Maps
* A map is an unordered collection of key-value pairs, also known as an associative array, a hash table or a dictionary
* Maps are used to look up a value byy its associate key
* The lenght of a map can change as we add items to it
* Maps are not sequential. We have x[1] and with an array that would imply there must be a x[0], but maps don't have this requirement
* If we lookup an element that doesn't exist in a map, it returns the zero value for the value type

### Chapter 7 Functions
#### Functions 
* A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. 
* Functions (also known as procedures or subroutines) are often represented as a black box
* Functions start with the keyword func, followed by the function's name. 
* The parameters (inputs) of the function are defined like this: name type, name type, ...
* After the parameters we put the return type.
* Collectively the parameters and the return type are known as the function's signature.
* Finally we have the function body which is a series of statements between curly braces
* The names of the parameters don't have to match in the calling function.
* Functions don't have access to anything in the calling function
* Functions are built up in a “stack”. Each time we call a function we push it onto the **call stack** and each time we return from a function we pop the last function off of the stack.
#### Named return type
* We can also name the return type: ```
 func f2() (r int) {
r = 1
return
}```
#### Multiple return values
* Go is also capable of **returning multiple values** from a function
#### **Variadic functions**
* By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
```func Println(a ...interface{}) (n int, err error)
The Println function takes any number of values of any type.```  
We can also pass a slice of ints by following the slice with ...: xs := []int{1,2,3}
fmt.Println(add(xs...))
#### Closures
* It is possible to create functions inside of functions
* When you create a local function like this it also has access to other local variables
* A function like this together with the non-local variables it references is known as a closure
* makeEvenGenerator returns a function which generates even numbers.  
Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.
#### Recursion
* A function is able to call itself
#### Defer, Panic, Recover
* **Defer** - Go has a special statement called defer which schedules a function call to be run after the function completes
* Defer is used when resources need to be freed in some way.   
For example, when we open a file we need to make sure we close it later.
`f, _ := os.Open(filename)
defer f.Close()
`
* This has 3 advantages:
    - it keeps our Close call near our Open call so its easier to understand
    - if our function had multiple return statements Close will happen before both of them
    - deferred functions are run even if a run-time panic occurs
* **Panic & Recover**
* We can handle a run-time error with the built-in recover function. Recover stops the panic and returns the value that was passed to the call to panic
* A panic generally indicates a programmer error (for example attempting to access an index of an array that's out of bounds, forgetting to initialize a map, etc.) or an exceptional condition that there's no easy way to recover them. (Hence the name "panic")

### Chapter 8 - Pointers
* Pointers reference a location in memory where a value is stored rather than the value itself (they point to something else)
* In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value. 
* `*` is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the pointer points to. When we write *xPtr = 0 we are saying "store the int 0 at the memory location xPtr refers to"
* & is used to find the address of a variable. &x returns a *int (pointer to an int) because x is an int
* Another way to get a pointer is to use the built-in new function.
**new** takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
* Go is a garbage collected programming language which means memory is cleaned up automatically when nothing refers to it anymore
* Pointers are rarely used with Go's built-in types, but as we will see in the next chapter, they are extremely useful when paired with structs
### Chapter 9 - Structs and Interfaces
* The type keyword introduces a new type. It's followed by the name of the type, the keyword struct, to indicate we are defining a **struct** type and a list of fields inside of curly braces. 
* Each field has a name and a type
* Like functions we can collapse fields that have the same type
* Initialization can be done in a variety of ways: 
>go
>var c1 Circle  
>c1.x = 10  
>c1.y = 5  
>fmt.Println(c1.x, c1.y, c1.r)  
>c2 := new(Circle)      
>c3 := Circle{x: 0, y: 0, r:5}  
>c4 := Circle{0, 0, 5}  
>fmt.Println(circleAreaC(c4)).
* Like with other data types, this will create a local Circle variable that is by default set to zero. For a struct zero means each of the fields is set to their corresponding zero value (0 for ints, 0.0 for floats, "" for string s, nil for pointers, ...)
* We can access fields using the . operator
* One thing to remember is that arguments are always copied in Go.   
If we attempted to modify one of the fields inside of the circleArea function, it would not modify the original variable.
* **Methods**: In between the keyword func and the name of the function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating
the function in this way it allows us to call the function using the . operator
>func (c *Circle) area() float64 {
>return math.Pi * c.r*c.r
>}
* **Embedded Tyoes**
    * A struct's fields usually represent the has-a relationship. For example a Circle has a radius
    * We would rather say an Android is a Person, rather than an Android has a Person. Go
supports relationships like this by using an embedded type. Also known as **anonymous fields**, embedded types look like this:  
> type Android struct {
>   Person
>  Model string
> }

    * **has-a relationship**
    type Android1 struct {
      Person Person
      Model string
    }
    
    * **is-a relationship**
    type Android2 struct {
      Person
      Model string
    }
    
* **Interfaces**
* Like a struct an interface is created using the type keyword, followed by a name and the keyword interface . But instead of defining fields, we define a “method set”.  
A method set is a list of methods that a type must have in order to “implement” the interface.
