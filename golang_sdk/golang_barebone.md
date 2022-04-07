## FAQ:

0. How to setup golang and start execution
1. What are tokens
2. What is Keyword
3. What is an Identifier
4. What is declaration and Initialization
5. What is literal
6. What is Statements and Expression in Go
7. What is package keyword in go
8. What the compiler invokes automatically, like a semicolon after each newline
9. FUNCTION
10. CHANNEL
11. GOROUTINES
12. INTERFACE AND METHODS
13. POINTER
14. STANDARD LIBRARY

## A0. GoLang setup

1. Arch linux setup

   ```
   sudo pacman -S go
   #it will install it in /usr/lib/go
   ```

   > in ~/.profile add below and re login user

   ```
   #golang
   export GOROOT=/usr/lib/go
   export GOPATH=$HOME/go
   export GOBIN=$GOPATH/bin
   export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

   #protobuf
   export PATH=$PATH:/home/common/Programs/Arch_linux/protoc-3.17.1-linux-aarch_64/bin
   ```

   > a helloworld.go file

   ```
   package main

   import "fmt"

   func main() {
      fmt.Println("Hello, Arch!")
   }
   ```

   > run using

   ```
   go run helloworld.go
   ```

## A1. TOKENS

1. identifiers
2. keywords
3. operators and punctuation
4. literals
5. white spaces from below are ignored but not ignored if below are used to separate tokens, otherwise tokens will get combined and will form a single token
   - spaces
   - horizontal tabs
   - carriage returns
   - newlines

## A2. KEYWORDS

    RESERVED KEYWORDS:
    	#these following keywords can't be use as identifier
    	break        default      func         interface    select
    	case         defer        go           map          struct
    	chan         else         goto         package      switch
    	const        fallthrough  if           range        type
    	continue     for          import       return       var

## A3. IDENTIFIER

    PREDECLARED IDENTIFIERS:
    	Types:
    		bool byte complex64 complex128 error float32 float64
    		int int8 int16 int32 int64 rune string
    		uint uint8 uint16 uint32 uint64 uintptr
    	Constants:
    		true false iota
    	Zero value:
    		nil
    	Functions:
    		append cap close complex copy delete imag len
    		make new panic print println real recover
    	#predeclared nil is the only untyped value who has not a default type in Go
    	#nil is not a keyword so nil:=123 can be use
    BLANK AND NON-BLANK IDENTIFIER

## A4. DECLARATION and Initialization

    #binds an NON-BLANK identifier to a TYPE,fun,....etc
    TYPE DECLARATION:
    	#predeclared types
    		bool byte complex64 complex128 error float32 float64
    		int int8 int16 int32 int64 rune string
    		uint uint8 uint16 uint32 uint64 uintptr
    	#using TypeName

    	#using TypeLiteral
    		#construction of this following composite types can be done using type literal
    			array, struct, pointer, function, interface, slice, map, and channel
    	Alias DECLARATION:
    		identifier =  Type
    		#both identifier and the Type are identical
    		eg: type x = int
    		eg: type y = x
    	Defined DECLARATION:
    		#defined type always distinct from any other typesincluding the type it is created from. ,means once you define a type then it's a new type,
    		identifier Type
    	underlying type:
    		#each type T has an underlying types.
    		# If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself
    		# else T's underlying type is the underlying type of the type to which T refers in its type declaration.
    		eg
    			type (
    				A1 = string  //underlying type(UT) of string is string, A1 is string
    				A2 = A1		//UT of A1 is string , A2 is string
    			)

    			type (
    				B1 string   //UT of string is string, B1 is string
    				B2 B1		//UT of B1 is string, B2 is string
    				B3 []B1     //UT of []B1 is []B1, B3 is []B1  //here []B1 is type literal
    				B4 B3		//UT of B3 is []B1, B4 is []B1
    			)
    variable declarion:
    	var v int  //initialize a zero same as var v int = 0
    	var v int = 2
    	var v = 2 //implicit type
    	v := 2  //:= for "var" and this is implicit type
    variable initialization after declaration:
    	v = 2
    	v := 2 (error since v is already declared,  := for new variable)
    constant defining
    	#you can not declare a constant
    	#you can not change a constant by reinitializing since it's not a declaration
    	const v int = 2
    	const v = 2 //implicit type
    #declare then use ;declare = declare[initialize] //optional
    #declare can be dynamic type with initialize or static type with or without initialize
    #CO: something has a type but not declared
    #ER: if something is declared but not used

## A5. LITERALS

THEY CONSTRUCT VALUES TO IDENTIFIER WITH THEIR TYPE
integer
floating-point
imaginary
rune
string
composite: CompositeLit = LiteralType LiteralValue
struct
array
slice
map
function literal
const supports :
true : const x = true
false : const x = false
iota : const x = iota
numeric literal : const x = 2
string literal : const x = "ok"
rune literal : const x = '\a'
imaginary literal : const x = 0i
variable supports :
const supports
function literals : var x = func() {}
custom array type literal : type cst [3]int; var x = cst{1,2,3}
array literal : var x = [3]int{123}
custom struct type literal : type cst struct{ x int }; var x = cst{x: 2}
struct literal : var x = struct{ x int }{2}
we can declare an identifier as a type
we can declare an identifier as a type and we can attach one or more method to it
we can declare an identifier as a constant we have use constant literal to declare both value and type
we can declare an identifier as a variable and [we can declare type of that variable OR we can use variable literal to declare both value and type]
we can declare an identifier as a function and [we can declare type of that function OR we can use function literal to declare both value and type]
we can declare an identifier as a method and [we can declare type of that method OR we can use method literal to declare both value and type]
construction of a new type or alias type using type literals:
new type can be constructed from type  
 eg: type x int, here x and int both are type but x is constructed from int
eg: type y x
new type can be constructed from type literal  
 eg: type art [2]int, here x and []int both are type but x is constructed from type literal, [2]int is an ArrayType literal
eg: type slt []int, []int is an SliceType literal
eg: type stt struct{field_name int} //struct{field_name int is an StructType literal, and field_name is a filed with type int
eg: type pot *int //*int is a PointerType literal
eg: type fut func() // func() is a FunctionType literal
eg: type ift interface{} // interface{} is an InterfaceType literal
eg: type mpt map[int]int // map[int]int is a MapType literal
eg: type cht int or type cht chan int or type cht chan <- int //int, chan int etc are channel literal, So? int is a literal too?
alias type can be constructed from type

## A6. Statements and Expression

    In go
    	++ and -- are statements not expressions we can't do a:=b++
    	The operator affects the operand directly. It will not produce any value. so only b++ works
    	The assignment operator is not an expression, we cant do x=2 instancr do x:=2
    In C
    	++ and -- are expressions, so we can do, a = b++
    	The assignment operator is an expression so we can do x=2

## A7. PACKAGES

1. No need to import a file.go in file1.go iff both have same package name
2. if one function lives in another file.go and both are in same package then use
   ```
   go run .
   #this will do the linking
   ```

## A8. What does compiler invoke

    a newline or end of file may trigger the insertion of semicolon.

## A9. FUNCTION

1. function can have name
2. function can be anonymous in this case arguments can be passed after immediate the end of the function
   and value can be assiged if it returns any as var:=func..{}(arg..)
3. function can return more than one value
4. function parameters can have type or not
5. function must include return types if they return any and return types can have name.
6. function can have variables number of parameters func sum(nums ...int) int {}
7. function can be recursive
8. defer function call, eg: defer myfunc(). this will be called after end of sorrounding functions
9. pass parameters by value, even if a pointer is sent, the function will make a copy of it, in c it's different

```
func alter(x *int) {
	*x = 5 //the caller x will be changed now
	y := 7
	x = &y //but the caller will not point to y because x is passed by value not reference, it's juste changed in this scope
}
```

10. function as a parameters or return a function, such functions are called higher order function. eg:
    ```
    func fun1(x int, f2 fun2(int) int){
    	return f(x)
    }
    //and call it by
    fun1(3,inc)  //and
    fun inc(x int) int{
    	x++;
    	return x
    }
    ```
11. custom function type
    ```
    type output func(string) string
    func hello(name string) string{return name}
    func main(){var f output;f=helllo;f("peter")}
    ```
12. filter function
13. closure
    ```
    func enclosingfun() func() int{ 			//here "func() int" is a return type which is also anonymous function
    	i := 0
    	return func() int{
    		i++
    		return i
    	}
    }
    nextInt := enclosingfun()
    nextInt() //output 1
    nextInt() //output 2
    ```

## A10. CHANNEL

    don't think channel as a data structure
    A channel allows one goroutine to signal another goroutine about a particular event.
    p := <-ch // Receive
    ch <- "paper" // Send

## A11. GOROUTINES

## A12. INTERFACE AND METHODS

```
type Interface_name interface{
	method1() return_type
}

type User_data_type_1 struct{x int}
type User_data_type_2 struct{x int, y int}


func (receiver_data User_data_type_1) method1() return_type{
	return receiver_data.x
}

func (receiver_data User_data_type_2) method1() return_type{
	return receiver_data.x+receiver_data.y
}

func helperfun(dynamic_data Interface_name){
	//
	//call dynamic_data.method()
}

var data_of_type_1 = User_data_type_1{x:2}  //value of type User_data_type_1 is a Interface_name
var data_of_type_2 = User_data_type_1{x:2, y:3}
helperfun(data_of_type_1)  //helper fun
helperfun(data_of_type_2)
```

interface can act as type1, type2
iff type1 and type2 implements a method which is declared in the interface.
it works in function parameter also in return type
var i error = Mystruct{}, here error is an interface having a Error signature and Mystruct implements it and i can accept the structure

### But why interface over Struct Method?

Look at the following program, We can store different struct as a slice using interface.

```
package main

import (
	"fmt"
	"math"
)

type Rect struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (rect Rect) Area() float64 {
	return rect.length * rect.width
}
func (circle Circle) Area() float64 {
	return 2 * math.Pi * circle.radius
}

type Figure interface {
	Area() float64
}

func main() {
	c1 := Circle{radius: 5}
	r1 := Rect{length: 5, width: 5}
	figures := []Figure{c1, r1}
	for _, f := range figures {
		fmt.Printf("\nArea %v", f.Area())
	}
}
```

## A13. POINTER

    *type is a pointer type eg: type x *string  or var x *string
    *x is explicit dereferencing/indirect,ie value of that address in this variable

    var x = struct{x int}{x:0}		  V/S		var x = &struct{x int}{x:0}
    	fmt.Println(x)		   {0}					fmt.Println(x)			&{0}
    	fmt.Println(x.x)	   0					fmt.Println(x.x)		0   	.means the dereferencing
    	fmt.Println((&x).x)	   0					//fmt.Println((&x).x)
    	//fmt.Println((*x).x)						fmt.Println((*x).x)		0
    	fmt.Println(&x)		   &{0}	 				fmt.Println(&x)			0xc00000e028
    	//fmt.Println(*x)							fmt.Println(*x)			{0}
    	fmt.Println(&(x.x))	   0xc000100010			fmt.Println(&(x.x))		0xc000018050
    	//fmt.Println(*(x.x))						//fmt.Println(*(x.x))

    var x = []int{0,1}		   				    var x = &[]int{0,1}
    	fmt.Println(x)		   [0,1]				fmt.Println(x)			&[0 1]
    	fmt.Println(x[0])	   0				    //fmt.Println(x[0])				[]means the dereferencing
    	//fmt.Println((&x)[0])  					//fmt.Println((&x)[0])
    	//fmt.Println((*x)[0]) 						fmt.Println((*x)[0])	0
    	fmt.Println(&x)		   &[0 1]				fmt.Println(&x)			0xc00000e028
    	//fmt.Println(*x)							fmt.Println(*x)			[0 1]
    	fmt.Println(&(x[0]))   0xc0000b6020			//fmt.Println(&(x[0]))	0xc000018050
    	//fmt.Println(*(x[0]))						//fmt.Println(*(x[0]))

    var x = [2]int{0,1}		   				    var x = &[2]int{0,1}
    	fmt.Println(x)		   [0,1]				fmt.Println(x)			&[0 1]
    	fmt.Println(x[0])	   0				    fmt.Println(x[0])		0
    	fmt.Println((&x)[0])   0					//fmt.Println((&x)[0])
    	//fmt.Println((*x)[0]) 						fmt.Println((*x)[0])	0
    	fmt.Println(&x)		   &[0 1]				fmt.Println(&x)			0xc00000e028
    	//fmt.Println(*x)							fmt.Println(*x)			[0 1]
    	fmt.Println(&(x[0]))   0xc0000b6020			fmt.Println(&(x[0]))	0xc000018050
    	//fmt.Println(*(x[0]))						//fmt.Println(*(x[0]))

    var x = map[string]int{"a":0,"b":1}	   	    var x = &map[string]int{"a":0,"b":1}
    	fmt.Println(x)		     map[a:0 b:1]		fmt.Println(x)			 &map[a:0 b:1]
    	fmt.Println(x["a"])	     0				    //fmt.Println(x["a"])
    	//fmt.Println((&x)["a"]) 					//fmt.Println((&x)["a"])
    	//fmt.Println((*x)["a"]) 					fmt.Println((*x)["a"])	 0
    	fmt.Println(&x)		     &map[a:0 b:1]		fmt.Println(&x)			 0xc000102018
    	//fmt.Println(*x)							fmt.Println(*x)			 map[a:0 b:1]
    	//fmt.Println(&(x["a"])) 					//fmt.Println(&(x["a"]))
    	//fmt.Println(*(x["a"]))					//fmt.Println(*(x["a"]))


    var x = &2 will not work but var x = &[2]int{1,2} will  work, but why?
    what x holds. if  x holds address use * else don't use *
    The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array.

# A14. STANDARD LIBRARY

Archive,Zip,bytes,crypto,encoding,ast,syscall,regexp,net,io,os,database connection,

## UNCLEARED:

1. statement and expression:  
   It is said that, “A statement performs some action such as printing values, looping, or if statement.
   On the other hand, expressions also produce values and we can assign these values to new variables.”

## UPDATES

1. golang 1.18 support generic functions and the doc is this - https://go.dev/doc/tutorial/generics

## REFERENCES:

1.  [GOLANG LANGUAGE SPECIFICATION] https://golang.org/ref/spec
2.  [GOLANG BEHAVIOUR OF CHANNEL] https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
3.  [10 THINGS YOU DON'T KNOW ABOUT GOLANG] https://talks.golang.org/2012/10things.slide
4.  [GOLANG PACKAGES] https://golang.org/pkg
5.  [GOLANG BUILTIN PACKAGES] https://golang.org/pkg/builtin
6.  [RUNTIME GO FILE SOURCES] https://github.com/golang/go/blob/master/src/runtime/slice.go
7.  [FUNCTIONS IN GOLANG] https://zetcode.com/golang/function
