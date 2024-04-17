// https://dog.ceo/api/breeds/list/all
// https://dog.ceo/api/breed/bulldog/english/images/random
//

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

//	type Dog struct{
//		breed []string
//	}
type DogList struct {
	Message map[string][]string
}

type Output struct {
	Message  string `json:"message"`
	Status   string `json:"status"`
	Error    error
	Breed    string
	subBreed string
}

const (
	siteUrl = "https://dog.ceo/api/breed"
)

func main() {
	ch := make(chan Output, 10)
	var wg sync.WaitGroup
	var dl DogList
	client := &http.Client{}
	listUrl := fmt.Sprintf("%ss/list/all", siteUrl)
	// fmt.Println(listUrl)
	resp, err := client.Get(listUrl)
	if err != nil {
		log.Fatal(err)
	}
	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(string(respB))
	}
	if err := json.Unmarshal(respB, &dl); err != nil {
		log.Fatal(err)
	}
	ll := DogList{
		Message: map[string][]string{
			"bulldog": {
				"boston",
				"english",
				"french",
			},
			"african": {},
		},
	}
	fmt.Println(dl.Message)
	// for breed, subBreed := range dl.Message {
	for breed, subBreed := range ll.Message {
		wg.Add(1)
		go getImage(breed, "", ch, &wg)

		for v := range subBreed {
			wg.Add(1)
			go getImage(breed, subBreed[v], ch, &wg)
		}

	}

	wg.Wait()
	fmt.Println("Channel Closure")
	close(ch)
	for out := range ch {
		if out.Error == nil {
			fmt.Println(out)
		} else {
			fmt.Println(out.Error.Error())
		}
	}

}

func getImage(breed, subBreed string, ch chan<- Output, wg *sync.WaitGroup) {
	defer wg.Done()
	var url string
	if subBreed == "" {
		url = fmt.Sprintf("%s/%s/images/random", siteUrl, breed)
	} else {
		url = fmt.Sprintf("%s/%s/%s/images/random", siteUrl, breed, subBreed)
	}
	fmt.Println(url)
	data := Output{
		subBreed: subBreed,
		Breed:    breed,
	}
	resp, err := http.Get(url)
	if err != nil {
		data.Error = err
		ch <- data
		return
	}
	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		data.Error = err
		ch <- data
		return
	}

	if resp.StatusCode != http.StatusOK {
		data.Error = fmt.Errorf(string(respB))
		ch <- data
		return
	}
	if err := json.Unmarshal(respB, &data); err != nil {
		data.Error = err
		ch <- data
		return
	}
	fmt.Println(string(respB))
	ch <- data
}
