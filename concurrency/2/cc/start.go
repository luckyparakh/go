package cc

import (
	"fmt"
	"sync"
	"time"
)

func GetUser(id int, userCh chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Get User %d\n", id)
	time.Sleep(80 * time.Millisecond)
	userCh <- "Got User"
	wg.Done()
}

func GetUserLikes(id int, userCh chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Get User Likes %d\n", id)
	time.Sleep(120 * time.Millisecond)
	userCh <- "Got likes"
	wg.Done()
}

func GetUserFriends(id int, userCh chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Get User Friends %d\n", id)
	time.Sleep(40 * time.Millisecond)
	userCh <- "Got friends"
	wg.Done()
}
