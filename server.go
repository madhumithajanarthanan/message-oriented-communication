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

// A structure is created to store driver details
type Details struct {
	name            string
	currentLocation int
	age             int
	contact         int
}

var flag1 int = 0
var flag2 int = 0

//It is used to store the history of the client's ride
var test [100][]string

func main() {
	//creating a server and listening on port 8081
	//port can also be changed
	//test := [][]string{[]string{}, []string{}}
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	//closing the connections after handling all the clients
	defer l.Close()

	//loop to handle multiple clients
	for {

		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {

	//test := [][]string{[]string{}, []string{}}
	//details of the drivers who are registered
	driver1 := Details{"Ram", 50, 40, 9678012348}
	driver2 := Details{"Arun", 52, 44, 9645609808}

	pickup, err := bufio.NewReader(conn).ReadByte()
	drop, err := bufio.NewReader(conn).ReadByte()
	id, err := bufio.NewReader(conn).ReadByte()
	ID := string(id)
	log.Println(ID)
	if ID == "0" {
		test[0] = append(test[0], string(pickup), string(drop), " ")
		fmt.Println(test[0])
	}
	if ID == "1" {
		test[1] = append(test[1], string(pickup), string(drop), " ")
		fmt.Println(test[1])
	}

	//for debugging purpose
	//log.Println("Pickup location:", pickup)
	//log.Println("Drop location:", drop)

	//calculating the absolute distance between drivers and the clients
	t1 := int(pickup) - (driver1.currentLocation)
	t2 := int(pickup) - (driver2.currentLocation)
	//log.Println(t1)

	if t1 < 0 {
		t1 = -t1
	}
	if t2 < 0 {
		t2 = -t2
	}
	//calculating the cost by computing the distance between pickup and drop location
	dist := int(pickup) - int(drop)
	//log.Println(dist)
	if dist < 0 {
		dist = -dist
	}
	k := dist * 10
	cost := k
	log.Println(cost)               //for debugging
	fmt.Fprintf(conn, string(cost)) //sending the cost to the client

	//checking whether the client agreed to the cost and if he accepts, then we are sending the driver details to the client.
	buffer1, err := bufio.NewReader(conn).ReadBytes('\n')
	reply := string(buffer1)
	reply = strings.TrimSpace(reply)
	log.Println(reply)
	if reply == "ok" {

		if (t1 <= t2 && flag1 == 0) || (t1 > t2 && flag2 == 1) {

			fmt.Fprintf(conn, "your ride is accepted with ", driver1.name, "   contact: ", driver1.contact)
			flag1 = 1

		} else if (t1 > t2 && flag2 == 0) || (t1 <= t2 && flag1 == 1) {

			fmt.Fprintf(conn, "your ride is accepted with", driver2.name, "   contact: ", driver2.contact)
			flag2 = 1

		}
		log.Println("completed")
		//log.Println("History")
	}
	//conn.Close()

	//conn.Close()
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	//sending the ride history to client
	if ID == "0" {
		fmt.Fprintf(conn, "\n")
		for _, i := range test[0] {
			fmt.Fprintf(conn, i+"\n")
			log.Println(i)
		}
		// log.Println(test[0])
		fmt.Fprintf(conn, "completed"+"\n")
	}
	if ID == "1" {
		fmt.Fprintf(conn, "\n")
		for _, i := range test[1] {
			fmt.Fprintf(conn, i+"\n")
			log.Println(i)
		}
		//fmt.Fprintf(conn, (test[1]))
		fmt.Fprintf(conn, "completed"+"\n")
	}
	conn.Close()

	//conn.Close()
}
