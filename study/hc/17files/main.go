package main

import (
	"encoding/json"
	"fmt"
	"os"

)

// UrlInfo is struct which contains info related to URLs
type UrlInfo struct {
	Id     uint64 `json:"id"`
	Link   string `json:"link"`
	Encode string `json:"encode"`
}
var urls []UrlInfo
func main() {
	fmt.Println("Create Files")

	// file, err := os.Create("./myFile.txt")
	// checkNilErr(err)
	// numLine, err := file.WriteString("Hello World")
	// checkNilErr(err)
	// fmt.Println("Number of line written ", numLine)
	// defer file.Close()
	// readFile("./myFile.txt")
	createJsonF()
	createJsonF()
	readJsonF()
}

func readFile(fn string) {
	databtye, err := os.ReadFile(fn)
	checkNilErr(err)
	fmt.Println("File contents:", string(databtye))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createJsonF(){
	urlinfo := UrlInfo{
		Id: 1,
		Link: "http://gogle.com",
		Encode: "AAAA",
	}
	urls = append(urls, urlinfo)
	f, err := os.Create("urlInfo.json")
	if err != nil {
		panic(err)
	}
	if b, err := json.Marshal(&urls); err !=nil{
		f.Close()
		panic(err)
	}else if _, err := f.Write(b);err !=nil{
		f.Close()
		panic(err)
	}
}

func readJsonF(){
	var urlinfo []UrlInfo
	if bytedata, err := os.ReadFile("urlInfo.json"); err != nil {
		panic(err)
	}else if err := json.Unmarshal(bytedata, &urlinfo); err != nil {
		panic(err)
	}
	for k,val := range urlinfo{
		fmt.Printf("%v,%v",k,val.Link)
	}
}
