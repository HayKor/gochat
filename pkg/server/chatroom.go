package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type ChatRoom struct {
	Name string

	mu      sync.Mutex
	Clients map[net.Conn]bool
}

func (cr *ChatRoom) Broadcast(msg Message) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for client := range cr.Clients {
		msgToSend := fmt.Sprintf("[%s][%s]: %s", msg.RoomName, msg.From, msg.Content)
		_, err := client.Write([]byte(msgToSend))
		if err != nil {
			log.Println("Problem reading message from user ", err.Error())

			// TODO: Change that to be more graceful
			client.Close()
			delete(cr.Clients, client)
		}
	}
}

func (cr *ChatRoom) BroadcastSystemMessage(s string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for client := range cr.Clients {
		_, err := client.Write([]byte(s))
		if err != nil {
			log.Println("Problem while sending system message: ", err.Error())

			client.Close()
			delete(cr.Clients, client)
		}
	}
}
