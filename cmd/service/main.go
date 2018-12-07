package main

import "github.com/gin-gonic/gin"

func main() {
	// Http Server
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
	})
	r.Run()
}
