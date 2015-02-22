# [An introduction to programming in Go](http://www.golang-book.com)

# Integers

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

# Floating point numbers
- They are inexact
- Like integers they have a certain size (32 or 64 bits)
- Using a larger sized floating point number increases its precision (how many digits it can represent)
- In addition to number there are several other values which can be represented: 
    - "not a number" NaN for things 0/0
    - positive and negative infinity

Floating point types:
- **float32** (single precision)
- **float64** (double precision)

# Complex numbers
- **complex64**
- **complex128**

# Operators
+ addition
- subtraction
* multiplication
/ division
% remainder

# Strings
- A string is a sequence of characters with a definite length used to represent text.
- Go strings are made of individual bytes, usually one for each character.
(Characters from other languages like Chinese are represented by more than one byte)
- String literals can be created using double quotes `"Hello World"` or back ticks ``Hello World``. The difference between thse is that double quoted strings cannot contain newlines and they allow special escape sequences, like \n, which gets replaced with a newline, and \t gets replaced with a tab character.

# Booleans
- A boolean value is a special 1 bit integer type used to represent true and false (or on and off)
- Three logical operators are used with boolean values
&& and
|| or
! not

# Loops
- Go only has one loop type that can be used in a variaty of different way

# Documentation
* godoc fmt Println
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