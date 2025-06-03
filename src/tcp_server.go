package src

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = ":8080"

func StartTCPServer() (err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", port); err != nil {
		err = fmt.Errorf("error on creation listener: %v", err)
		return
	}
	log.Printf(" TCP server has been started on %s port\n Accepting connections...", port)
	var connection net.Conn
	if connection, err = listener.Accept(); err != nil {
		err = fmt.Errorf("error on accepting connection: %s", err.Error())
		return
	}
	log.Printf(" Accepted connection by %s", connection.RemoteAddr().String())
	reader := bufio.NewReader(connection)
	testData := make([]byte, 5)
	if _, err = reader.Read(testData); err != nil {
		err = fmt.Errorf("error on receiving data: %v", err)
		return
	}
	log.Printf("Data successfully received: %v", testData)
	if err = connection.Close(); err != nil {
		err = fmt.Errorf("error on closing %s connection: %v", connection.RemoteAddr().String(), err)
		return
	}
	log.Printf("%s connection successfully closed", connection.RemoteAddr().String())
	if err = listener.Close(); err != nil {
		err = fmt.Errorf("error on closing listener: %v", err)
		return
	}
	log.Print("Listener successfully closed")
	return
}
