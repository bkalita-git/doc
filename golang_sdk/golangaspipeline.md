# Golang as pipeline of types and actions

```
type MyInterface interface{
    a_method(parameter_1 int) (return_val_1 int, err error)
}

func (object *MyStruct)a_method(parameter_1 int) (return_val1 int, err error){
    return (object.field1,nil)
}

type MyStruct struct{
    field1 int
    field2 int
}
func NewMyStruct() *MyStruct{
    return &MyStruct{field1:1,field2:2}
}
```

> In our above example MyStruct is a type and it implements a method named **_"a_method()"_**. A function which needs any two variables of any Type as param such that each Type implements a particular method. example-

```
io.Copy(var1 type1,var2 type2)
//here according to Copy documentation, type1 must implement a write method and type2 must implement a Read method. That is the way we can pipeline tasks.
```

This is like overloading because Copy function can take variables of different types.

## Task1: applying-

```
action : io.Copy
type1  : os.File
typ12  : os.FIle
```

Code

```
file1, file1_err := os.Create("file1.txt")
file2, file2_err := os.Open("testfile.txt")
if file2_err != nil {
    log.Panic(file2_err)
}
if file1_err != nil {
    log.Panic(file1_err)
}

io.Copy(file1, file2)

file1.Close()
file2.Close()
```

## Task2: applying-

```
action : io.Copy
type1  : net/http.Server
type2  :
```

## Discover Types which implements interface and if does not exist then make a type and implement the interface.

1. Is is Type? if yes, Does it implement any interface? Does any Function created it?
2. Is is an Interface? if yes, Does it have any Type which implemented it? Does any function returns it? if No, Create Type and implement the Interface. Example- http.Handler

generic

1. Addr Interface
2. Conn Interface
   - Pipe [func]
   - Dial [func]
   - DialTimeout [func]
3. Listener Interface
   - UnixListener [struct]
   - TCPListener [struct]
   - Listen [func]
   -
4. /net/http

```
Server Struct<-Handler Interface[ServeHTTP]
Handler Interface:
    instance of HandlerFunc [func]
    instance of ServeMux [struct]
    instance of Custom_Type [struct] which implements ServeHTTP

we can make special type of Handler for example a ServeMuxHandler


we are basically Modifying the struct.
framework implements custom servemux
servemux ServeHTTP will execute the appropriate one
Handle and HandleFunc

route([string]{"GET","POST","DELETE"},"/bipul/<id:int>/",function_name)

```
