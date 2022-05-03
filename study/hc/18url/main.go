package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Url Parsing")
	myurl := "https://lco.dev/"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	databyte, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

	url1 := "https://lco.dev:3000/learn?coursename=reactjs"
	result, _ := url.Parse(url1)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	path := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
	}
	fmt.Println(path)
}
