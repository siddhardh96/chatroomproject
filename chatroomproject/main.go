package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//initialize gin
	r := gin.Default()

	//this create new classroom
	chatRoom := NewChatRoom()

	//route for handler functions
	r.GET("/join", JoinHandler(chatRoom))
	r.GET("/leave", LeaveHandler(chatRoom))
	r.GET("/send", SendHandler(chatRoom))
	r.GET("/messages", MessagesHandler(chatRoom))

	//run in the port 8080
	r.Run(":8080")
}
