package main

import (
	"log"

	"github.com/HayKor/gochat/pkg/server"
)

func main() {
	srv := server.NewServer(":3000")
	log.Fatal(srv.Start())
}
