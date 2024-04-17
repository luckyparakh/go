// https://www.youtube.com/watch?v=P4tckkcyef0
package main

import (
	"concurrent/cc"
	"concurrent/server"
	"fmt"
	"sync"
	"time"
)

func getUser(id int) string {
	fmt.Printf("Get User %d\n", id)
	time.Sleep(80 * time.Millisecond)
	return "Got User"
}

func getUserLikes(id int) string {
	fmt.Printf("Get User Likes %d\n", id)
	time.Sleep(120 * time.Millisecond)
	return "Got likes"
}

func getUserFriends(id int) string {
	fmt.Printf("Get User Friends %d\n", id)
	time.Sleep(40 * time.Millisecond)
	return "Got friends"
}

func main() {
	// start := time.Now()
	// getUser(1)
	// getUserLikes(1)
	// getUserFriends(1)
	// fmt.Println(time.Since(start))

	start1 := time.Now()
	userCh := make(chan string, 10)
	wg := &sync.WaitGroup{}
	go cc.GetUser(1, userCh, wg)
	wg.Add(1)
	go cc.GetUserLikes(1, userCh, wg)
	wg.Add(1)
	go cc.GetUserFriends(1, userCh, wg)
	wg.Add(1)
	wg.Wait()
	close(userCh)
	for data := range userCh {
		fmt.Println(data)
	}
	fmt.Println(time.Since(start1))

	s := &server.Serve{
		MsgCh:  make(chan server.Message),
		QuitCh: make(chan struct{}),
	}
	go s.Server()
	for i := 0; i < 10; i++ {
		go func() {
			server.SendMessage(s.MsgCh, server.Message{
				From:    "Abc",
				Payload: "Hi",
			})
		}()
	}

	// time.Sleep(2 * time.Second)
	s.QuitCh <- struct{}{}
}
