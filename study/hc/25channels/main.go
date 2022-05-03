package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main(){
	fmt.Println("Channels")
	simple()
	// simpleMultiReadSingleWrite() //deadlock
	channelAsArg()
	simpleMultiReadSingleWriteBuf()
	channelLoop()
	wg.Wait()

}

func simple(){
	// There is no way other than to create channel always use make
	// make is can/always used for making channel, slice and maps.
	ch := make (chan int)
	wg.Add(2)
	go func(){
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func(){
		i:=42
		ch <- i // Passning copy of i
		i=27 // Manipulate i
		wg.Done()
	}()
}

func simpleMultiReadSingleWrite(){
	// There is no way other than to create channel always use make
	// make is can/always used for making channel, slice and maps.
	ch := make (chan int)
	// It will create deadlock as there are 5 channels that are sending data but only one to receive
	// Channel operation are blocking ie function wait for channel for someone to read/put data.
	// If that doesn't happen than fx waits and causes deadlock
	wg.Add(1)
	func(){
		i:=42
		ch <- i
		wg.Done()
	}()
	for i:=0;i<4;i++{
		wg.Add(1)
		func(){
			i:= <- ch
			fmt.Println(i)
			wg.Done()
		}()
	}
}

func channelAsArg(){
	fmt.Println("Running channels as arg")
	ch := make(chan int)
	wg.Add(2)
	// Read only channel
	go func(ch <- chan int){
		fmt.Println(<-ch) // read from ch
		wg.Done()
	// kind of polymorphism as ch (passed here) is direction 
	// but in func arg unidirectional ch was used (read only).
	// in case of channels it is ok and runtime does that for us
	}(ch) 
	// Send only channel
	go func(ch chan <- int){
		i:=42
		ch <- i //write in ch
		//i = <- ch // read //invalid operation: cannot receive from send-only channel ch
		wg.Done()
	}(ch)
}

func simpleMultiReadSingleWriteBuf(){
	fmt.Println("Running channels with buffer")
	// There is no way other than to create channel always use make
	// make is can/always used for making channel, slice and maps.
	ch := make (chan int,5)
	wg.Add(1)
	func(){
		ch <- 20
		ch <- 30
		ch <- 40
		ch <- 50
		wg.Done()
	}()
	for i:=0;i<4;i++{
		wg.Add(1)
		func(){
			i:= <- ch
			fmt.Println(i)
			wg.Done()
		}()
	}
}

func channelLoop(){
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <- chan int){
		for i := range ch{
			fmt.Println("From Channel loop",i)
		}
		// Another way to loop channel
		// for{
		// 	if i,ok := <-ch;ok{
		// 		fmt.Println("From Channel loop",i)
		// 	}else{
		// 		break
		// 	}
		// }
		wg.Done()
	}(ch)
	go func(ch chan <- int){
		ch <- 12
		ch <- 23
		// If channel is not closed that loop at reader will cause deadlock 
		// as it will try to read value more than 2
		close(ch)
		wg.Done()
	}(ch)
}