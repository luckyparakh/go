package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var Secret = []byte(os.Getenv("SECRET_KEY"))

func main() {
	handleRequest()
}
func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJwt()
	fmt.Println(validToken)
	if err != nil {
		fmt.Fprintln(w, "Failed to generate token")
	}
	fmt.Println(validToken)
	fmt.Fprintln(w, string(validToken))
}

func GetJwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authrized": "true",
		"client":    "rp",
		"aud":       "rp.jwt.io",
		"iss":       "jwt.io",
		"exp":       time.Now().Add(time.Minute * 10).Unix(),
	})
	tokenStr, err := token.SignedString(Secret)
	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}
	return tokenStr, nil
}

func handleRequest() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
