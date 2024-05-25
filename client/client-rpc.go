/**
* TCPArithClient
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"sync"
	"time"
)

//Define shared types for data communication

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]
	log.Printf("start: %d\n", time.Now().Unix())
	//creates a connection request
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Args is a common type used for sending arguments to the server
	args := Args{17, 8}
	var reply int
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//calls the Multiply function of the Arith type on the server and gets the data back in reply
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	}(&wg)

	//Quotient is a shared type that is used to get data back from the server
	var quot Quotient
	//calls the Divide function of the Arith type on the server and gets the data back in quot
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
	log.Printf("done first: %d\n", time.Now().Unix())
	wg.Wait()
	log.Printf("done second: %d\n", time.Now().Unix())
}
