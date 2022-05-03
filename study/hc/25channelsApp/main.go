package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
)
const (
	logInfo = "INFO"
	logErr = "Error"
)
var wg sync.WaitGroup
type logger struct{
	time time.Time
	level string
	message string
}

func main(){
	ch := make(chan logger,50)
	//defer close(ch)
	wg.Add(1)
	go printLog(ch)
	for i:=0;i<4;i++{
		ch <- logger{time.Now(), logInfo,"abc"+strconv.Itoa(i)}
	}
	close(ch)
	wg.Wait()
}

func printLog(ch <- chan logger){
	for i := range ch{
		fmt.Println(i)
	}
	wg.Done()
}