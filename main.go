package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.New()

	flag.Parse()
	hub := newHub()
	go hub.run()

	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	_ = r.Run(":5000")
}
