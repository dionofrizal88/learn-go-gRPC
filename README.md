# svc-golang-gRPC

## Install the protocol compiler plugins for Go:

**Using this command':**
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
**Update PATH so that the protoc compiler can find the plugins::**
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```
**Create proto file in proto folder**
**Generate gRPC code using MAKEFILE**
```
$ make -f MakeFile greet
```
**The Command will run protoc command like:**
```
$ protoc -Icalculator/proto --go_opt=module=github.com/dionofrizal88/go-grpc --go_out=. --go-grpc_opt=module=github.com/dionofrizal88/go-grpc --go-grpc_out=. calculator/proto/*.proto
$ go build -o bin/calculator/server ./calculator/server
$ go build -o bin/calculator/client ./calculator/client
```
**Folder proto will contain grpc.pb.go and pb.go**
**Run this command to update go module:**
```
$ go mod tidy
```
**Running the server and client**
**Run the following command::**
```
$ ./bin/greet/server greet
$ ./bin/greet/client
```