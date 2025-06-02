package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // Optional: set release mode

	r := gin.Default()

	// ✅ Allow all IPs by setting to nil (and suppress warning safely)
	// err := r.SetTrustedProxies(nil)
	// if err != nil {
	// 	panic(err)
	// }

	chatRoom := NewChatRoom()

	r.GET("/join", JoinHandler(chatRoom))
	r.GET("/leave", LeaveHandler(chatRoom))
	r.GET("/send", SendHandler(chatRoom))
	r.GET("/messages", MessagesHandler(chatRoom))

	r.Run(":8080")
}
