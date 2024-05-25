## go rpc test TCP server and client

How to build and run server:
```
cd server
go build server-rpc.go
./server-rpc 
```

How to build and run client:
```
cd client
go build client-rpc.go
./client-rpc localhost:1234
```
