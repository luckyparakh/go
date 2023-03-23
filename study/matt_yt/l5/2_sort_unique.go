package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Empty Map
	words := map[string]int{}
	scan := bufio.NewScanner(os.Stdin)

	// Split scan on words
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}

	// If there is error while scanning, exit
	if err := scan.Err(); err != nil {
		os.Exit(-1)
	}
	// fmt.Println(words)

	// Define key values pair struct
	type kv struct {
		k string
		v int
	}

	// Empty slice of KV
	kvs := []kv{}

	// Convert Map into KV struct
	for k, v := range words {
		kvs = append(kvs, kv{k: k, v: v})
	}

	// Sort a slice of kvs with custom sort function (closure or anonymous or lambda function)
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v > kvs[j].v
	})

	// Print Top four
	for _, kv := range kvs[:6] {
		fmt.Println(kv)
	}
}
