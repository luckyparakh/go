package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt"
)

var Secret = []byte(os.Getenv("SECRET_KEY"))

func main() {
	http.Handle("/", isAuth(homepage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func isAuth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				aud := "rp.jwt.io"
				iss := "jwt.io"
				if !token.Claims.(jwt.MapClaims).VerifyAudience(aud, false) {
					return nil, fmt.Errorf("invalid Aud")
				}
				if !token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false) {
					return nil, fmt.Errorf("invalid Iss")
				}
				return Secret, nil
			})
			if err != nil {
				fmt.Fprint(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprint(w, "Invalid token")
		}
	})

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret Message")
}
