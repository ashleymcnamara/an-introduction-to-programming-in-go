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
