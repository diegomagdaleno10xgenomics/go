package main

import (
	"fmt"
	"log"
	"net"
)

const (
	NET = "tcp"
	HOST = "djxmmx.net"
	PORT = "17"
)

func main() {
	conn, err := net.Dial(NET, HOST + ":" + PORT)

	checkErr(err)

	defer conn.Close()

	msg := ""

	_, err = conn.Write([]byte(msg))

	checkErr(err)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	checkErr(err)

	fmt.Println(string(reply))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
