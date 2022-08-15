package main

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

func content(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	out := &strings.Builder{}
	buff := make([]byte, 1024)
	defer resp.Body.Close()
	n, _ := resp.Body.Read(buff)
	if n <= 0 {
		return "", errors.New("no data to read")
	}
	out.Write(buff[:n])
	time.Sleep(2000 * time.Millisecond)
	return out.String(), nil
}

type Counts map[string]int

func countWords(content string) Counts {
	out := Counts{}
	words := strings.Split(content, ",")
	for _, w := range words {
		out[w]++
	}
	return out
}

func contentAsync(url string, out chan<- string) error {
	c, err := content(url)
	if err != nil {
		return err
	}
	out <- c
	return nil
}

func countAsync(input <-chan string, output chan<- Counts) {
	for c := range input {
		output <- countWords(c)
	}
	close(output)
}
