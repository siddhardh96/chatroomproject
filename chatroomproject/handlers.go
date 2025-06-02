package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JoinHandler(cr *ChatRoom) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
			return
		}

		client := cr.Join(id)
		c.JSON(http.StatusOK, gin.H{"message": "Joined chat", "client_id": client.ID})
	}
}

func LeaveHandler(cr *ChatRoom) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
			return
		}
		cr.Leave(id)
		c.JSON(http.StatusOK, gin.H{"message": "Left chat"})
	}
}

func SendHandler(cr *ChatRoom) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		msg := c.Query("message")
		if id == "" || msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id or message"})
			return
		}
		if _, ok := cr.GetClient(id); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Client not found"})
			return
		}
		cr.SendMessage(id, msg)
		c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
	}
}

func MessagesHandler(cr *ChatRoom) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
			return
		}

		client, ok := cr.GetClient(id)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Client not found"})
			return
		}

		select {
		case msg := <-client.MsgChan:
			c.JSON(http.StatusOK, gin.H{"message": msg})
		case <-time.After(10 * time.Second): // timeout
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No new messages"})
		}
	}
}
