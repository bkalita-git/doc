> The Go runtime in the flexible environment is the software stack responsible for building and running your code. SO stack frame builds by runtime. For example, during its init function, the net/http/pprof package registers HTTP handlers that provide debugging information.

- A go program with `import _ "apackage"` is for without using any of `apackage` but to execute `init()` inside `apackage` before `main()`
- We can't overload same name in go.
- Most of things in golang is happened at compiled time.

```
import _ "net/http/pprof"
```

- Profiling
- Tracing
- Debuging

Delve and GDB

- Stop compiler optimization for debugging

```
go build -gcflags=all="-N -l"
```

- run the compiled code and use go tool pprof

```
./a.out
go tool pprof -http localhost:8888 http://localhost:7777/debug/pprof/heap
```

- build flag for heap escapes analysis

```
go build -gcflags="-m -l"
```

## call stack

```
main.main
    main.other_function //stack frame
    main.main           //stack frame
    rumtime.main        //stack frame
    runtime.gopark      //stack frame
runtime.gopark
```

## Reference

1. https://pkg.go.dev/net/http/pprof
