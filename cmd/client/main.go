package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			message := make([]byte, 1024)
			n, err := conn.Read(message)
			if err != nil {
				fmt.Println("Disconnected from server")
				return
			}
			fmt.Print(string(message[:n]))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			msg := scanner.Text()

			_, err := conn.Write([]byte(msg))
			if err != nil {
				fmt.Println("Error sending message:", err)
				break
			}
		}
	}
}
