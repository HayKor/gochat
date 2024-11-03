package main

import (
	"log"

	"github.com/HayKor/gochat/pkg/client"
)

func main() {
	client := client.NewClient()
	log.Fatal(client.Start())
}
