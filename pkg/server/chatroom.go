package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type ChatRoom struct {
	Name string

	mu      sync.RWMutex
	Clients map[net.Conn]bool
}

func (cr *ChatRoom) Broadcast(msg Message) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	for client := range cr.Clients {
		msgToSend := fmt.Sprintf("[%s][%s]: %s", msg.RoomName, msg.From, msg.Content)
		_, err := client.Write([]byte(msgToSend))
		if err != nil {
			log.Println("Problem reading message from user ", err.Error())

			// TODO: Change that to be more graceful
			cr.mu.Lock()
			client.Close()
			delete(cr.Clients, client)
			cr.mu.Unlock()
		}
	}
}

func (cr *ChatRoom) BroadcastSystemMessage(s string) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	for client := range cr.Clients {
		_, err := client.Write([]byte(s))
		if err != nil {
			log.Println("Problem while sending system message: ", err.Error())

			client.Close()
			delete(cr.Clients, client)
		}
	}
}
