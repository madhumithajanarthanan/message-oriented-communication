package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	connHost = "localhost"
	connPort = "8084"
	connType = "tcp"
)

func main() {
	//connecting to server
	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	//reading the pickup and drop location from the client
	fmt.Print("pickup loacation: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("drop location: ")
	input2, _ := reader.ReadString('\n')

	//sending the details to server
	conn.Write([]byte(input))
	conn.Write([]byte(input2))
	conn.Write([]byte("0"))

	//receiving the cost of the ride from server
	cost, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("The cost:", (cost))
	//buff := make([]byte, 1024)
	//n, _ := conn.Read(buff)
	//log.Printf("Receive: %s", buff[:n])

	//reading the reply from the client and sending to server and receiving the driver details
	input1, _ := reader.ReadString('\n')
	conn.Write([]byte(input1))
	message1, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Reply: " + message1 + "\n")
	fmt.Println('\n')

	//receiving the ride history of the client from the server
	for {
		history, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Client left.")
			conn.Close()
			return
		}
		history = string(history)
		history = strings.TrimSpace(history)
		if history == "completed" {
			break
		}

		fmt.Println("History: ", (history))
	}

}
