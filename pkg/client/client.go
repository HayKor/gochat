package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Start() error {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	defer conn.Close()
	c.conn = conn

	go c.acceptMesssagesLoop()
	c.sendMessagesLoop()

	return nil
}

func (c *Client) acceptMesssagesLoop() {
	for {
		message := make([]byte, 1024)
		n, err := c.conn.Read(message)
		if err != nil {
			fmt.Println("Disconnected from server")
			return
		}
		fmt.Print(string(message[:n]))
	}
}

func (c *Client) sendMessagesLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			msg := scanner.Text()

			_, err := c.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err)
				break
			}
		}
	}
}
