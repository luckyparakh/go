package main

import (
	"time"
	"math/rand"
)

const alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateKey() string{
	key:=""
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5;i++{
		key += string(alphaNum[rand.Intn(len(alphaNum))])
	}
	return key
}