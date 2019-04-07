package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":5000"
}

func main() {
	r := gin.New()

	flag.Parse()
	hub := newHub()
	go hub.run()

	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	port := getPort()
	_ = r.Run(port)
}
