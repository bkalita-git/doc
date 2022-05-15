## GRPC

IS A RPC FRAMEWORK

> [a client application can directly call a method on a server application on a different machine as if it were a local object]

gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types.

- On the server side, the server implements this interface and runs a gRPC server to handle client calls.
- On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.

```
			GRPC SERVER  <=====> GRPC STUB
```

protocol buffer as the message format.

1. Unary RPC
2. Server Streaming RPC
3. Client Streaming RPC
4. Bidirectional Streaming RPC

## Why GRPC?

with SOAP and REST , websocket,we need client library, for example HTTP library.

## Data serialization

is the process of converting data objects present in complex data structures into a byte stream for storage, transfer and distribution purposes on physical devices. the byte stream can be of type XML, JSON , BSON, YAML , MessagePack, and protobuf  
eg:

```
var x = struct{int x}{x:2}
```

can be converted to [JSON byte stream] or [JSON text file] using some json encoder.  
the process of serializing an object is called **marshalling**. the object can be a structure object, or a class object with functions in it

- how protobuf encoding works:
  Protocol Buffer compiler will trasform .proto file to golang or other language class. in that .proto file the data structures are there for encoding/decoding purpose and for that data structures protobuff will produce language specific code so we can import and use to encode/decode. Protocol Buffer is a language build in 2001 at google, grpc in 2016 and golang in 2009. 1. use .proto file which containes the structure you want to encode 2. generate language specific code using the .proto file so you can use the structure 3. create the object by initialization of the structure 4. export the object by desired format, eg:/to file, byte
- how protobuf decoding works:
  1.  use .proto file which containes the structure you want to decode
  2.  generate language specific code using the .proto file so you can use the structure
  3.  import file from the earlier export eg:/from file, byte string
  4.  convert the file to object
- how to setup protocol buffer in golang
  1.  download the compiler for aarch64[already compiled]
  2.  export /home/common/Programs/Arch_linux/protoc-3.17.1-linux-aarch_64/bin and $ source ~/.profiles
  3.  use protoc command

I got 3 directores of GO in arch linux:

```
/home/bipul/go
	d_bin : executables
		go-outline
		gopkgs
		gocode
		gopls
		lv
	d_pkg
		d_sumdb

		d_mod


/usr/lib/go:
	d_api
		has 20 text files
	d_bin
		go
			the go compiler executable file
		gofmt
			executable file
	d_doc link to /usr/share/doc/go
	d_lib
		d_time
			README file
			update.bash file
	d_misc

	d_pkg
		all machine code packages with extensions .a and .h and some executable file
	d_src
		all the source code, eg built in ,runtime go  files
	d_test
		so many go files
	VERSION


/usr/bin/go links to /usr/lib/go/bin/go
```

```
export GOROOT=/usr/lib/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

```
$ go get github.com/golang/protobuf/protoc-gen-go
#/home/bipul/go/bin/protoc-gen will be there
```

```
#currency.proto file
	syntax = "proto3";
	option go_package = "/currency"; //where to out
	message RateRequest{
		string Base = 1;
		string Destination = 2;
	}

$ protoc --go_out=. currency.proto
				or
$ protoc currency.proto --go_out=plugins=grpc:.
```

## gRPC

$ go install google.golang.org/grpc@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

1.  Define a communication.proto file

    ```
    syntax = "proto3";
    option go_package = "/communication"; //where to out
    service Route{
       rpc GetFeature(Point) returns (Feature) {}
    }
    message Point {
    	int32 latitude = 1; #this 1 means a field identidy
    	int32 longitude = 2;
    }

    ```

2.  transform to go code

    ```
    protoc communication.proto --go_out=. --go-grpc_out=.
    ```

    will produce

    ```
    this include the protobuff marshalling and unmarshalling code
    communication/communication.pb.go

    #this includes all server and client interface and structure which implements them
    communication/communication_grpc.pb.go
    ```

3.  make module inside communication directory
    ```
    $ go mod init communication
    $ go mod tidy
    ```
4.  Import our communication package in client.go and server.go and use it.  
     server.go

    ```
    import (
    	"communication"
    )
    ```

    /\*
    interface
    Addr <-

        conn <- Dial
        PacketConn <- ListenPacket

        Listener <- Listen,

        listen returns listener and listener.Accept returns conn

\*/

/\*
Request response model :
like in rest we used get post method and then server uses servemux to call different function based on url and parameter.
solution 1. gRPC. is actually protobuf over RPC, so, we will still use HTTP but not the rest style, so no url concept to execute different function.
solution 2. we can use RPC,RSocket,json rpc which don't need an application level layer and json,glob as our payload. this is also not rest.
solution 3. Websocket, there is a difference between websocket and http keep-alive header. is that webscoker is full duplex so server can also send to client.
But websocket is also at Application layer. it works over HTTP but distinct from it. it is like rest and uses url but more than that.

    for fast developing, HTTP will be best over low level OSI layer ofcourse, and for the technique the gRPC is best for it.

    Generalized view:
    	A Universal Defined Protocol and What Architecture for muxing and How data are passes.

     Kademlia DHT,

Cgo lets Go packages call C code

from .proto -> encoding_decoding.go -> grpc_client_server.go{interface for client and server but you have to implement them properly}

\*/

```
proto   -> go
message -> struct
service -> interface{Method Signature} implemented for serviceclient struct

.proto is not the database, it's the structure, server will have a database and client can use it.
```

## BIDIRECTIONAL STREAMING

rpc SayHello(stream HelloRequest) returns (stream HelloReply);
HelloRequest, HelloReply struct : auto
server gets a stream type struct
client gets a stream type struct
server implements SayHello: Manual
client implements SayHello: Auto
HelloRequest sent to server
HelloReply sent to client

REFERENCES:
