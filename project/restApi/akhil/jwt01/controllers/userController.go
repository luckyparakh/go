package controllers

import (
	"context"
	"fmt"
	"jwt/database"
	"jwt/helpers"
	"jwt/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": validationErr.Error()})
			return
		}
		countEmail, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "Error occured while checking email"})
		}
		countPhone, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "Error occured while checking phone number"})
		}
		if countEmail > 0 || countPhone > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "User with these email and phone number is already present"})
		}
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()
		token, refresh_token, _ = helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refresh_token
		resultInsertionNUmber, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"err": "User not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNUmber)
	}

}

func Login() {
	fmt.Println("")
}

func HashPassword() {

}

func VerifyPassword() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
	}
}
