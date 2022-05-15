## CONTEXT

```
import "context"

func main(){
    root_ctx := context.Background() //root context
    ctx, cancel := context.WithTimeout(root_ctx, time.Duration(134*time.Millisecond())) //ctx here is a new context with timeout made from root_context
    defer cancel() //release resources at the end of the main function
    ctx = context.withValue(ctx,"key1","value1") //new key value context made from timeout context
    fmt.Println(ctx.Value("key1"))
}
//use when to propagae information in different go routine,
```

## SELECT STATEMENT

```
//golang select statement
select{
    case message:= <-chan1:
        fmt.Println("chan1")
    case message:= <-chan2:
        fmt.Println("chan2")
}

//select executes/chose the case which channel got a data first. so, if chan2 gets the data before chan1 then chan2 will get executed. it's different than switch because we can select among multiple channel or say like multiple variable. in switch case, it was reading a variable data and then chose the case.

```

## CHANNEL

```
//channel error

var x chan string //x is a variable of type string
x<-"data"
fmt.Println(<-x) //will produce error, deadlock but why?
```

```
//channel error
var x chan int

go func() {
    x <- 3
}()
fmt.Println(<-x)
time.Sleep(time.Second * 60)
```

```
//correct channel
x := make(chan int)

go func() {
    x <- 3
}()
fmt.Println(<-x)
time.Sleep(time.Second * 60)
```

Goroutine 1 Goroutine 2
msg:=<-c(stops and wait) c<-"hi"(sends)
synchronization.
go can detect deadlock at runtime - when in one goroutine is waiting to receive a message on a channel and no other go routine is sending to that channel. and also a channel can be closed on a sender or receiver side, because go channel are passed by references.

chan<-"message" when hit this instruction then the routine is blocked because it expects someone to receive. to solve this problem we can use buffered channel.

## EMPTY SLICE vs 0 initialized SLICE

how a slice can be empty?
a pointer points to empty struct that is struct{}
and empty == nil
why do we need empty?
empty struct,
empty slice,
empty channel

integer declated without a value will default to `0`
string declared without a value will defailt to `""`
a pointer declared without a value will default to `nil`
a channel declared without a value will default to `nil`
nil cause panic when dereferencing, so if check for nil may cause error.
