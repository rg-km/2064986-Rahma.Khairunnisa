package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dari contoh yang telah diberikan, buatlah HTTP server sederhana dengan menggunakan Gin.
// Buatlah server yang memiliki address localhost dengan port 3000.
// Buatlah route "/hello" yang menampilkan JSON {"message": "hello"} dan "/world" yang menampilkan JSON {"message": "world"}

func GetGinRoute() *gin.Engine {
	//return &gin.Engine{} // TODO: replace this
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": 1,
			"content": "Hello World",
		})
	})
	r.Run(":3000")
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
