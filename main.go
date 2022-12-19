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
