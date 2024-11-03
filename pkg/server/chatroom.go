package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type ChatRoom struct {
	Name    string
	Clients map[net.Conn]bool
	mu      sync.Mutex
}

func (cr *ChatRoom) Broadcast(msg Message) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for client := range cr.Clients {
		_, err := client.Write([]byte(fmt.Sprintf("[%s][%s]: %s\n", msg.RoomName, msg.From, msg.Content)))
		if err != nil {
			log.Println("Problem reading message from user ", err.Error())

			// TODO: Change that to be more graceful
			client.Close()
			delete(cr.Clients, client)
		}
	}
}
