/**
* TCPArithServer
 */

package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

//Define shared types for data communication

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//Define functions called by the client on the server

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	//time.Sleep(50 * time.Second)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)
	//registers the type Arith to be used with rpc
	err := rpc.Register(arith)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	//server starts listening on port 1234
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//accepts a connection request and serves it on a different goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
