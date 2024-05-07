package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathKey struct {
	PathName string
	Original string
}

func (p PathKey) fileName() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}

func CASPathTransform(k string) PathKey {
	hash := sha1.Sum([]byte(k))
	// fmt.Println(hash)
	hashStr := hex.EncodeToString(hash[:])
	// fmt.Println(hashStr)
	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}
	return PathKey{
		PathName: strings.Join(paths, "/"),
		Original: hashStr,
	}
}

type PathTransformer func(string) PathKey

type StoreOpts struct {
	PathTransformer PathTransformer
}

type Store struct {
	StoreOpts
}

var DefaultPathTransformer = func(key string) string {
	return key
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pk := s.PathTransformer(key)
	fmt.Println(pk)

	if err := os.MkdirAll(pk.PathName, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(pk.fileName())
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("%d bytes written to disk: %s", n, pk.fileName())
	return nil
}
