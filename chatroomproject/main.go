package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	chatRoom := NewChatRoom()

	r.GET("/join", JoinHandler(chatRoom))
	r.GET("/leave", LeaveHandler(chatRoom))
	r.GET("/send", SendHandler(chatRoom))
	r.GET("/messages", MessagesHandler(chatRoom))

	r.Run(":8080")
}
