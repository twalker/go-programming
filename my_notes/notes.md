* Go is all about type


* A slice allows you to hold many values of the same type; structs allow you to hold many values of different types.


## Zero values
int 0
string/slice ""
pointers, function, interfaces, rlices, channles, maps nil


fmt 
  Println - print line
  Printf - Format print
SprintX - to string (a variable)
FprintX - to a file or a writer interface

casting in go is called "Conversion"

underlying type

() = parens
[] = brackets
{} = curly braces

* Pass by value (not by reference or any other term)

# Numeric
Binary:
2 powerN
3 porch lights = 8 messages (2 pow 8)

coding schemes
utf-8 is the worlds most popular coding scheme
bit == 0 or 1
byte == 8 bits
binary digits  == bits
 
byte is an alias for uint8
rune (aka a utf8 character) is an alias for int32

in general use type int of float64 (if decimals are needed)

# String
* Strings are immutable
* Strings are a slice of bytes
%#U = format string for utf8 hexidecimal (character)


N^0 = 1

Decimal - Base 10
42, 420
|---|---|---|---|---|
| ten thousands | thousands | hundreds | tens | ones |
| 10^4 | 10^3 | 10^2 | 10^1 | 10^0 |
| 4 | 2 | 4 | 2 | 0 |

Binary - Base 2
| thirty twos |sixteens | eights | fours | twos | ones |
| 2^5 |2^4 | 2^3 | 2^2 | 2^1 | 2^0 |
| 1| 0| 1 | 0 | 1 | 0 |
42
|| | 1 | 1 | 0 |
1 4 32

Hexidecimal - Base 16 - 0-9 digits, 10-15 = A-F
| 65,536's | 4096's | 256s | 16's | ones |
| 16^4 | 16^3 | 16^2 | 16^1 | 16^0 |
911
|||3|8 |F|
38F (or 0x38F) in hexidecimal is 911 
0x prefix denotes that it is hexidecimal

Formatters
%T = type
%d = decimal
%b = binary
%#x = hexidecimal
%#U = unicode

 
### Pointers
jimPointer := &jim
&variable = & is an operator to give me the memory address of the value this variable is pointing at
*pointer  = * is an operator to give me the value this memory address is pointing at
Turn address into value with *address
Turn value into address with &value


Control flows
- Sequence
- Loop
- Conditionals


Loops
for init; condition; post {

}


if (statement) {

}
if init statement; statement {

}
e.g
if x:= 42; x==42 {
  
}

switch statement

switch {
  case false:
    //statement
  case true:
    //statement
  case true:
    //no fall through
  case true:
    fallthrough // keyword
  default:
    // statement
}

switch value {
  case val1:
  case val2:
  default:
}


switch value {
  case val1, val2, val3:
  case val4:
  default:
}

## Aggregate types


// composite literal
x := type{values}
x := []int{1,2,3}

// looping over slice with range
x := []int{1, 2, 3}
for i, v := range x {
  fmt.Println(i, v)
}

// slice of a slice
x[start (including):end (up to but not including)]
x[1:3] 

append
y := []int{4,5,6}
x = append(x, 2, 3)
x = append(x, y...)

delete
x = append(x[:1], x[2:]...)

make (allocates an array but does dynamic allocation)
make([]T, length, capacity)
make([]int{}, 10, 100)

multi-dimensional slice
jb := []string{"James", "Bond", "Chocolate"}
mp := []string{"Miss", "Moneypenny", "Strawberry"}
xp := [][]string{jb, mp}

maps
* maps don't throw errors on non-existing keys
map[keyType]valueType{/*composite literal*/}

m := map[string]int{
  "James": 32,
  "Moneypenny": 27,
}

m["James"]
// has: "comma ok" idiom

v, ok := m["noexist"]
ok == false if key not in the map.
do something if key exists
if v, ok := m["no exist"]; ok {
  doSomethingWithExistingItem()
}

add to a map
m["todd"] = 34

// range over map
for k, v := range m {
  fmt.Println(k, v)
}
delete from map
delete(mapname, key)

if v, ok := m["foo"]; ok{
  delete(m, "todd")
}


## structs
* in go, we don't say "object" or "class", we say "value of type" 
* inner type gets promoted to outer type (not "child" and "parent")

x, y = 44, 45
x, y = y, x

anonymous structs
p1 := struct {
  first string
  last string
}{
  // composite literal
  first: "James",
  last: "Bond",
}


Go is Object Oriented
1. Encapsulation
  a. state ("fields")
  b. behavior ("methods")
  c. exported & unexported; viewable & not viewable
2. Reusability
  a. inheritance ("embedded types")
3. Polymorphism
  a. interfaces
4. Overriding
  a. "promotion"

In Go:
* we create VALUES of a certain TYPE that are stored in VARIABLES
  and those VARIABLES have IDENTIFYIERS
* you don'tcreate classes, you create a TYPE
* you don't instantiate. you create a VALUE of a TYPE

* It is a bad practice to alias types


## functions
func (r receiver) identifier(parameters) (return(s)) {...}

* we define our func with paramaters (if any)
* we call our func with arguments (if any)
* everything in Go is PASS BY VALUE

variadic functions
xi := []int{3, 5, 84}
// unfurling values (spreading)
total := sum(xi...)
// variadic parameters -- converts 0 or more parameters of a single type to slice (rest)
func sum(x ...int) int {
  sum := 0
  for _, v := range x {
    sum += v
  }
  return sum
}

* "The only mistake in life is not asking questions"

Methods
reciever is in first order.

113. Methods 5:24 he talks about coming from other languages, the ease of programming, and the thoughtfullness of the designers
Interfaces
* Interfaces allow us to define behavior, and also allows us to do Polymorphism.
* how to define any Go variable!!: 
  ```
  keyword identifier type
  var hotdog int
  type person struct
  func foo()
  ```
type human interface{
  speak()
}
* A value can be of more than one type!!
An interface says, "If you've got this method, then you're my type"


// assertion (more powerful than conversion)
func bar(h human){
  switch h.(type) {
    case person:
      h.(person).first
    case secretAgent
      h.(secretAgent).first
  }
}

* Bill Kennedy good instructor on more advanced Go

Anonymous func -- anonymous self-executing functions
func(x int){

}(1/*execute func arguments*/)

func expression
f := func(x int) {}
f(1984)

returning a function
func foo() func() int {
  return func() int {
    return 451
  } 
}

* Functions are types

callbacks
func identifier(callback signature, parameters, return type)
See callback exercise

scope

{code block creates a scope}

// closure
a := incrementor() // uses x for this value
b := incrementor() // uses it's own x incrementor
a()
b()

func incrementor() func() int {
  var x = int
  return func() int {
    x++
    return x
  }
}


recursion

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}
// loop instead of recursion
func loopFactorial(n int) int {
  total := 1
  //for init; condition; post {} 
  for ; n > 0; n-- {
    total *= n
  }
  return total
}


* A true sign of a master, is one that makes the difficult look easy
* Focus on what's important, not what's urgent



## Pointers
`&` operator gets address in memory
Here's the value aaaaannnnndddd what is the address?

Type `*int` is a pointer to an int (or other type)
var b *int = &a

`*` operator dereference an address

a := 42
fmt.Println(&a) // & gives the address

b = &a
fmt.Println(*b) // * gives the value stored at an address when you have an address

* use pointer when you need to modify a value
* _Pass by value_ (forget pass by reference, or pass by copy)

Mutate === Change

Method set (is a type) === functions attached to a type

* a non-pointer receiver works with values that are pointers or non-pointers
* a pointer receiver only woks with values that are pointers

Receivers     Values
---------     ------
(tT)          T and *T
(t *T)        *T

# Writer and Reader interfaces

# sort
x := []int{5,7,6}
sort.Ints(x)

## sort custom
see example.

# bcrypt
hashes values

### Concurrency vs. parallelism


Concurrency: 
Design pattern - can run in paralell

Parallelism:
Multiple CPU

Running concurrent code in parallel

sync primatives: WaitGroup (Add, Done)

"It takes time to understand the wisdom of the masters."

function literal === anonymous function

* self-executing anonymous function:
  func(){}()


# Channels

Channels block!
Both send and recieve need to happen at the same time. They are synchronized.
Send will block until recieved.


// make a channel
c:= make(chan int)
cs:= make(chan <- int) // or send-only
cr:= make(<-chan int) // or recieve-only

// put a value on a channel
c <- 42
// take values off a channel
<- c

fanOut, fanIn concurrent design patterns

package context for request scoped data (e.g. userId, auth token, correlationId)

# Error handling

* handle errors where they occur
* always check for errors
* re-assign `err` 
```
f, err := os.Create("names.txt")
if err != nil {
  fmt.Println(err)
  return
}
defer f.Close()

r := strings.NewReader("Wassup")
io.Copy(f, r)
```
``` 
if err != nil {
  fmt.Println("err happened", err)
  // prefer log, because it be output to a file
  log.Println("err happened", err)
  log.Fatalln(err) // exit 1
  log.Panicln() // exit 1, but calls deferred cleanup from call stack and can be recovered

  panic(err)
}
```
// common idiom with init clause in if condition
```
if r:= recover(); r != nil {
  log.Println("Recovered in f", r )
}
```

Errors with Info
fmt.Errorf("some err", arg, arg)


# Documentation
go doc fmt.Println

Go Goals
* efficient execution
* efficient compilation
* ease of programming

# Testing

```go

import "testing"
// Test${NameOfFunction} // <- Convention
func TestAdd(t *testing.T) { // <- Needs to accept this parameter
	v := Add(a, b)
	exp := 4
	if v != exp {
    // Like log.Fatal
		t.Errorf("Expected %v, got %v", exp, v)
	}
}

// Table test
type test struct {
  data []int
  answer int
}

tests := []test{
  test{[]int{21,21}, 42},
  test{[]int{3,4,5}, 12},
  test{[]int{-1,0,1}, 0},
}

func TestMySum(t *testing.T) {
  for _, v := range tests {
    x:= mySum(v.data...)
    if x != v.answer {
      // "Expected" "Got" idiom
      t.Error("Expected", v.answer, "Got", x)
    }
  }
}
```
* // Output in Example test can act as test and documentation

## linting


TestX
ExampleX
BenchmarkX



Remember to BET

BET = Benchmark + Example + Test
```go
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.B)
```

## Commands
```sh
go test
go test -bench .
go tes -cover
go test -coverprofile c.out
go tool cover -html=c.out

golangci-lint run
```

