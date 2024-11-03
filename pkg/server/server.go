package server

import (
	"net"
)

type Server struct {
	ListenAddr string
	ChatRooms  map[string]*ChatRoom

	listener net.Listener
	quitCh   chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		ChatRooms:  make(map[string]*ChatRoom),
		quitCh:     make(chan struct{}),
	}
}

func (s *Server) GetOrCreateChatRoom(name string) *ChatRoom {
	if room, ok := s.ChatRooms[name]; ok {
		return room
	}
	room := &ChatRoom{
		Name:    name,
		Clients: make(map[net.Conn]bool),
	}

	// Register that room
	s.ChatRooms[name] = room
	return room
}
