## Encoding

By default, gRPC uses Protocol Buffers, Google’s mature open source mechanism for serializing structured data (although it can be used with other data formats such as JSON).

Benefits of protobuf
- Common message exchange description, that means that you can describe the messages that you are passing
- fast encoding and decoding
- Space efficient encoding
- Simple to write/read
- Many *generated* languages supported

## Protobuf
Protoco buffers are google's language-neutral, platform-neutral extensible mechanism for serializing structured data

## Installation

Install protoc --> https://grpc.io/docs/protoc-installation/
```sh
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```



```sh
protoc --go-out=. demo.proto
```

## Important

well known types --> https://protobuf.dev/reference/protobuf/google.protobuf/

```sh
# to generate a slice
repeated string params = position; 
# to generate maps
map <string, string> params = position; 
# to generate enum

# nesting
```
