package main

import (
	"os"

	"github.com/nenodias/grpc-cep/client"
	"github.com/nenodias/grpc-cep/server"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "client":
			client.Run()
		case "server":
			server.Run()
		}
	} else {
		panic("You must inform client or server by argument")
	}
}
