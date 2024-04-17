package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	server := newServer()
	server.Start()
	// wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go func(i int) {
			// defer wg.Done()
			server.addCh <- fmt.Sprintf("user_%d", i)
			// server.AddUser(fmt.Sprintf("user_%d", i))
		}(i)
	}
	// close(server.addCh)
	// wg.Wait()
	time.Sleep(2 * time.Second)
	server.quitCh <- struct{}{}
	fmt.Println(server.user)
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
