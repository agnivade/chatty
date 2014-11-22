package main

import (
	"fmt"
	//"log"
	//"math/rand"
	"net"
	//"time"

	"./ArchiveFileMessage"
	"code.google.com/p/goprotobuf/proto"
)

func main() {

	msg_chan := make(chan *message.ArchiveFileMessage)

	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	go channelListener(msg_chan)

	for {
		fmt.Println("Waiting for connections")
		// accept a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleConn(conn, msg_chan)
	}

}

func handleConn(conn net.Conn, msg_chan chan *message.ArchiveFileMessage) {
	fmt.Println("Connection established")
	defer conn.Close()

	//Create a data buffer of type byte slice with capacity of 1
	data := make([]byte, 1)
	//Read the data waiting on the connection and put it in the data buffer
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg_length := data[0]
	fmt.Printf("Read %d bytes\n", n)
	if n == 1 {
		fmt.Printf("Got the length - %d\n", msg_length)
	}

	//Create a data buffer of type byte slice with capacity of the length
	data = make([]byte, msg_length)
	//Read the data waiting on the connection and put it in the data buffer
	n, err = conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read %d bytes\n", n)

	fmt.Println("Decoding Protobuf message")
	//Create an struct pointer of type ProtobufTest.TestMessage struct
	protodata := new(message.ArchiveFileMessage)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(data[0:n], protodata)
	if err != nil {
		fmt.Println(err)
		return
	}

	// pushing the message to the channel
	msg_chan <- protodata

}

func channelListener(msg_chan chan *message.ArchiveFileMessage) {
	for {
		message := <-msg_chan
		fmt.Printf("Got msg from client %s\n", message)
	}
}
