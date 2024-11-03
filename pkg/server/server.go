package server

import (
	"fmt"
	"log"
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

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.listener = listener

	go s.acceptLoop()
	<-s.quitCh

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("Couldn't accept connection:", err.Error())
			continue
		}
		log.Printf("New connection from: %s\n", conn.RemoteAddr().String())
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	var roomName string
	var userName string

	fmt.Fprint(conn, "Enter chat room name: ")
	fmt.Fscanln(conn, &roomName)

	fmt.Fprint(conn, "Enter your name: ")
	fmt.Fscanln(conn, &userName)

	room := s.GetOrCreateChatRoom(roomName)
	room.Clients[conn] = true

	defer func() {
		room.mu.Lock()
		fmt.Fprint(conn, "You've been disconnected from the room.\n")
		delete(room.Clients, conn)
		room.mu.Unlock()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading message.", err)
			break
		}
		msg := string(buf[:n])
		room.Broadcast(Message{
			RoomName: roomName,
			From:     userName,
			Content:  string(msg),
		})
	}
}
