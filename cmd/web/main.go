package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/admin", func(c *gin.Context) {
	})
	r.Run()
}
