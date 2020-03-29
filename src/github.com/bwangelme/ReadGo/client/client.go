package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/bwangelme/ReadGo/server"
)

func Main(num1, num2 int) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	args := &server.Args{num1, num2}
	var reply int
	fmt.Println(num1, num2)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
