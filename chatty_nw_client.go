package main

import (
	"fmt"
	"log"
	"net"

	"./ArchiveFileMessage"
	"code.google.com/p/goprotobuf/proto"
)

func main() {

	// TODO: have a channel and create a for loop where a function
	// keeps sending messages after a time interval on the channel
	// Have another for loop which listens on the channel and as and when messages
	// come, it just sends across to the socket (with length prefixed)
	// Also have a quit message where if it reads that, it closes the socket
	// after sending a quit message on the channel
	// connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	client_msg := new(message.ArchiveFileMessage)
	client_msg.Message = proto.String("yet another client")
	client_msg.MsgType = proto.Int32(int32(422))
	log.Printf("Sending msg from client %s\n", client_msg)

	raw_msg, err := proto.Marshal(client_msg)
	if err != nil {
		fmt.Println(err)
	}

	length := len(raw_msg)
	var buf = make([]byte, 1)
	buf[0] = byte(length)
	n, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Sent %d bytes\n", n)

	n, err = conn.Write(raw_msg)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Sent %d bytes\n", n)

	conn.Close()

}
