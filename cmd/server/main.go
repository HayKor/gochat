package main

import (
	"log"

	"github.com/HayKor/gochat/pkg/server"
)

func main() {
	server := server.NewServer(":3000")
	// go func() {
	// 	for msg := range server.msgCh {
	// 		log.Printf("From %s: %s", msg.From, string(msg.Payload))
	// 	}
	// }()
	log.Fatal(server.Start())
}
