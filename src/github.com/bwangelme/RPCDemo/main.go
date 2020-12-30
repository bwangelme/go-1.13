package main

import (
	"flag"

	"github.com/bwangelme/RPCDemo/client"

	"github.com/bwangelme/RPCDemo/server"
)

var isServer = flag.Bool("server", false, "Is Running Server")
var n1 = flag.Int("n1", 0, "Number 1")
var n2 = flag.Int("n2", 0, "Number 2")

func main() {
	flag.Parse()

	if *isServer {
		server.Main()
	} else {
		client.Main(*n1, *n2)
	}
}
