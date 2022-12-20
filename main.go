package main

import (
	"chat/twilio"
	"net/http"

	"github.com/gin-gonic/gin"
)

var twilioClient = twilio.NewClient()

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/conversation", createConversation)
	r.POST("/user", createUser)
	r.Run()
}

type createConversationArgs struct {
	FriendlyName string `json:"name"`
}

func createConversation(c *gin.Context) {
	args := createConversationArgs{}
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if args.FriendlyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "name key should have non-empty value"})
		return
	}

	sid, err := twilioClient.CreateConversation(args.FriendlyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "success", "sid": sid})
}

type createUserArgs struct {
	Identity     string `json:"username"`
	FriendlyName string `json:"name"`
}

func createUser(c *gin.Context) {
	args := createUserArgs{}

	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if args.Identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "username key should have non-empty value"})
		return
	}

	sid, err := twilioClient.CreateUser(args.Identity, args.FriendlyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "success", "sid": sid})
}
