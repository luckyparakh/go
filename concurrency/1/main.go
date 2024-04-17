// https://www.youtube.com/watch?v=z75DKfOfDA4&t=91s
package main

import (
	"fmt"
	"time"
)

type Server struct {
	user   map[string]string
	addCh  chan string
	quitCh chan struct{}
}

func newServer() *Server {
	return &Server{
		user:   make(map[string]string),
		addCh:  make(chan string),
		quitCh: make(chan struct{}),
	}
}

func (s *Server) AddUser(user string) {
	s.user[user] = user
	fmt.Printf("User Added %s\n", user)
}

func (s *Server) Start() {
	fmt.Println("Start")
	// go s.loop()
	go s.loop1()
	time.Sleep(2 * time.Second)
	
}

func (s *Server) loop() {
	fmt.Println("Loop")
	for {
		u, ok := <-s.addCh
		if !ok {
			break
		}
		s.user[u] = u
		fmt.Printf("User Added %s\n", u)
	}
}

func (s *Server) loop1() {
	fmt.Println("Loop1")
	for {
		select {
		case u := <-s.addCh:
			s.user[u] = u
			fmt.Printf("User Added %s\n", u)
		case <-s.quitCh:
			fmt.Println("Quit")
			close(s.addCh)
			close(s.quitCh)
			return
		}
	}
}
