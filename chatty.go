package main

import (
	"log"
	"math/rand"
	"time"

	"./ArchiveFileMessage"
	"code.google.com/p/goprotobuf/proto"
)

func main() {

	c := make(chan *message.ArchiveFileMessage)
	go pingerfunc(c)
	go pongerfunc(c)

	// just an empty for loop for now
	for {
		time.Sleep(time.Second)
	}

}

func pingerfunc(c chan *message.ArchiveFileMessage) {
	i := 0
	for {
		pinger_msg := new(message.ArchiveFileMessage)
		pinger_msg.Message = proto.String("hello world from pinger")
		pinger_msg.MsgType = proto.Int32(int32(i))
		log.Printf("Sending msg from pinger\n")
		c <- pinger_msg
		// sleeping otherwise there is a chance this function itself gobbles it up
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)

		msg := <-c
		log.Printf("Received message - %s\n", msg)
		if *msg.MsgType == 5 {
			//sending a quit message over the chan
			ender_msg := new(message.ArchiveFileMessage)
			ender_msg.Message = proto.String("hello world from ender")
			ender_msg.MsgType = proto.Int32(int32(100))
			log.Printf("Sending quit msg from pinger\n")
			c <- ender_msg
			return
		}
		i++

	}
}

func pongerfunc(c chan *message.ArchiveFileMessage) {
	i := 0
	for {
		msg := <-c
		log.Printf("Received message - %s\n", msg)
		if *msg.MsgType == 100 {
			log.Printf("Ending message received\n")
			return
		}
		i++

		ponger_msg := new(message.ArchiveFileMessage)
		ponger_msg.Message = proto.String("hello world from ponger")
		ponger_msg.MsgType = proto.Int32(int32(i))
		log.Printf("Sending msg from ponger\n")
		c <- ponger_msg
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)

	}
}
