package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a route and a handler function
    router.GET("/hello", func(c *gin.Context) {
        // Respond with JSON
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Another route with parameters
    router.GET("/greet/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello " + name,
        })
    })

    // Start the web server on port 8080
    router.Run(":8080")
}
