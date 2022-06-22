package helpers

import (
	"context"
	"fmt"
	"jwt/database"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type signedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection = database.OpenCollection(database.Client, "user")
var SECRET_KEY = os.Getenv("secret_key")

func GenerateAllTokens(email, first_name, last_name, user_type, user_id string) (string, string, error) {
	claims := &signedDetails{
		Email:      email,
		First_Name: first_name,
		Last_Name:  last_name,
		User_type:  user_type,
		Uid:        user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaim := &signedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 168).Unix(),
		},
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	signedRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	return signedToken, signedRefreshToken, err
}

func UpdateToken(signedToken, signedRefreshToken, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})
	Update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"update_at", Update_at})
	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
	return
}

func ValidateToken(clientToken string) (claims *signedDetails, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &signedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*signedDetails)
	if !ok {
		msg = fmt.Sprintf("Token is invalid")
		msg = err.Error()
		return
	}
	msg = fmt.Sprintf("Token is valid")
	return
}
