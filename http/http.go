package http

import (
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	Host string
	Port string
}

func (s Server) InitServer() {

	listen, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	fmt.Printf("Server is listening on port %v", s.Port)

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
