package main

import (
	"github.com/mbaum0/badsec/server"
)

func main() {
	server := server.CreateServer()
	server.Run()
}
