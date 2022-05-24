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
