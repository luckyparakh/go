package server

import "fmt"

type Message struct {
	From    string
	Payload string
}

type Serve struct {
	MsgCh  chan Message
	QuitCh chan struct{}
}

func (s *Serve) Server() {
	for {
		select {
		case msg := <-s.MsgCh:
			fmt.Println(msg.From, msg.Payload)
		case <-s.QuitCh:
			fmt.Println("Closing server")
			return
		}
	}
}

func SendMessage(ch chan<- Message, m Message) {
	ch <- m
}
