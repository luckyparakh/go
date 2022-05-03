package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	//getRequest()
	postRequest()
	postFormRequest()
}

func postFormRequest() {
	const myurl = "http://localhost:8000/postform"
	values := url.Values{}
	values.Add("state", "up")
	values.Add("city", "kanpur")
	resp, err := http.PostForm(myurl, values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

}

func postRequest() {
	const myurl = "http://localhost:8000/post"
	data := strings.NewReader(`
		{
			"city":"Pali",
			"country":"India"
		}
	`)
	resp, err := http.Post(myurl, "application/json", data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	databyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
	// var responseBuilder strings.Builder
	// responseBuilder.Write(databyte)
	// fmt.Println(responseBuilder.String())
}

func getRequest() {
	const myurl = "http://localhost:8000/get"
	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	databyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.ContentLength)
	fmt.Println(string(databyte))

	//One more way
	var responseBuilder strings.Builder
	responseBuilder.Write(databyte)
	fmt.Println(responseBuilder.String())

}
