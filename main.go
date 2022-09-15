package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUser = "terry"
const mongoPassword = "G2yskjJc5sGFYDSi"
const uri = "mongodb+srv://terry:G2yskjJc5sGFYDSi@simpchat.k1ulald.mongodb.net/?retryWrites=true&w=majority"

var ctx = context.TODO()

var router *gin.Engine

func homePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{},
	)
}

func accountSignUp(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"account-signup.html",
		gin.H{},
	)
}

func accountLogin(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"account-login.html",
		gin.H{},
	)
}

func mongoClient() *mongo.Client {

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return mongoClient
	
}

func main() {

	client := mongoClient()

	client.Database("SimpChat").CreateCollection(ctx, "user", )

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", homePage)
	router.GET("/signup", accountSignUp)
	router.GET("/login", accountLogin)
	router.Run(":8080")
	
}
