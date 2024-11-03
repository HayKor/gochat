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
		msgToSend := []byte(fmt.Sprintf("[%s][%s]: %s", msg.RoomName, msg.From, msg.Content))
		_, err := client.Write(msgToSend)
		if err != nil {
			log.Println("Problem reading message from user ", err.Error())

			// TODO: Change that to be more graceful
			client.Close()
			delete(cr.Clients, client)
		}
	}
}
