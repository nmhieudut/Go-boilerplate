package main

import (
	"fmt"
	"rest-api/config"
	"rest-api/database"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println(database.Client)
	r := gin.Default()
	fmt.Println(config.Setting)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}