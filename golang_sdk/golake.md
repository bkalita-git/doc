sql/driver package has interfaces if someone's willing to build custom sql driver. Other wise to use already made driver, One should import the driver to the program and use it by sql package. The link of available 3rd party drivers are https://golang.org/s/sqldrivers

A program when you write it's not the end. It gets run by another programs called runtime with garbage collector.
Rust is without runtime and no GC.

type assertion only on interface with return 2 things.
type conversion can be in possible datatype

```
func add(object EXPECTEDINTERFACETYPE){
    object.IKnowMethod()
}
now the object must implement an interface
```

the concept of interface: An interface is combinations of common properties of different things.
hmm, can a `cup` and `smartphone` have common properties? -

```
1. Use property :  Both cup and smartphone can be used
2. Price Property : Both cup and smartphone has a price
3. Materials : Both cup and smartphone is built by some materials
4. Weight : Both have different mass.
5. Fragility : Both have different effect on drop.
...
...
type Thing interface{
    Price() float64
    Weight() float64
    Fragility() float64
    Materials() []string
    Use() string

}
```

back quote
FAT Table

interface are ends with "er". That is Writer, Reader etc.

direct pointer(array pointer) and indirect pointer(slice pointer to a array pointer)

See, #golang design pattern
# Golang design pattern
- Creational
	- Abstract Factory
	- Builder
	- Factory
	- Object Pool
	- Prototype
	- Singleton
- Behavioural
	- Chain of Responsibility
	- Command
	- Iterator
	- Mediator
	- Memento
	- Null Object
	- Observer
	- State
	- Strategy
	- Template Method
	- Visitor
- Structural 
	- Adapter
	- Bridge
	- Composite
	- Decorator
	- Facade
	- Flyweight
	- Proxy
	
	
## CONTEXT
each request makes a goroutine inside the httpserver. Now, developer can make additional go routine for that request handler function. Now we want to control all those goroutines and stuff for that particular request. and for that we use context package.
1 context has 4 methods
```
error()
done()
value()
deadline()
```
so it's useful when client request the server and before the server sending response the client disconnect. so the server processing is useless.


In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Ie. you can't just cancel the ctx by calling ctx.Cancel()

comma for same action on multiple
delete from
updade   set
'a%' same as a*
_a% this _ is first position
[acs] means any of acs these
[!letter]
like/not like


## package
how to run a package? well it should be from the main package.

## struct tags in go 
type User struct{
	Name string `example:"name"`
}
reflect package can use this

## project layout
>serviceatlas:serviceatlas module
	>app:app package
		>logger: logger package
		>container: container package
			>
		>config: config package
	>applicationservice: applicationservice package	
		>
	>cmd: main package
		>grpcserver: main package
		>grpcclient: main package
	>domain
		>model: model package
		>usecase: usecase package
	>script: sql file = manual run
	>tool: tool package
	
