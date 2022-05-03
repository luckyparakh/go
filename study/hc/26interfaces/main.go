package main

import "fmt"

func main(){
	var w Writer = ConsoleWriter{}
	w.Write("w")

	var wc WriterCloser = NewConsoleWriter()
	wc.Write("wc")
	wc.Close("wc")

}

type Writer interface{
	Write(string) (int, error)
}

type Closer interface{
	Close(string)
}

//Composite Interface
type WriterCloser interface{
	Writer
	Closer
}

type ConsoleWriter struct{

}

func (cw ConsoleWriter) Write(str string) (int, error){
	fmt.Println("Write", str)
	return 0,nil
}

func (cw ConsoleWriter) Close(s string){
	fmt.Println("Close", s)
}

func NewConsoleWriter() *ConsoleWriter{
	return &ConsoleWriter{}
}