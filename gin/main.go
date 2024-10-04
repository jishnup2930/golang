package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skil map")

	server := gin.Default()
	server.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Skill map home page",
		})
	})
	server.Run(":8080")
}
