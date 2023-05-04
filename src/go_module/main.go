package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"dengwj.vip/module/services"
)

func main() {
	println("hello world")

	services.TestUser()
	services.TestLogin()
	fmt.Printf("%v\n", "你好")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
