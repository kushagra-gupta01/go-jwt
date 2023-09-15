package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/form3tech-oss/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kushagra-gupta01/go-jwt/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email				string
	First_Name	string
	Last_Name		string
	User_id			string
	User_type		string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user") 

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string)(signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Email: email,
		First_Name: firstName,
		Last_Name: lastName,
		User_id: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(199)).Unix(),
		},
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err !=nil{
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}