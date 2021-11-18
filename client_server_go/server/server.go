package main


import (
	"fmt"
	"os"
	"net"
	"log"
	"bufio"
)


const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)


func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

	// Listening and the main loop.
	l, err := net.Listen(connType, connHost + ":" + connPort)
	if err != nil{
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		// Accept client connections.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle clients concurrently.
		go handleConnection(c)
	}
}


func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))
	
	conn.Write(buffer)

	handleConnection(conn)
}
