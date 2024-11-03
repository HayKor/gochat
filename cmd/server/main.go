package main

import (
	"log"

	"github.com/HayKor/gochat/pkg/server"
)

func main() {
	server := server.NewServer(":3000")
	log.Fatal(server.Start())
}
