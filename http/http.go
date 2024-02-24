package http

import (
	"fmt"
	"log"
	"net"
	"os"
)

func InitServer() {
	const HOST = "localhost"
	const PORT = "8080"
	const TYPE = "tcp"

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	defer conn.Close()
	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", buffer[:n])
	}
}
