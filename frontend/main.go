package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Serve static files
    r.Static("/static", "./static")

    // Define your routes
    r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    // Endpoint to handle Wordle solving
    r.POST("/solve", solveWordle)

    // Run the server
    r.Run(":8080")
}

func solveWordle(c *gin.Context) {
    // Parse input
    var input struct {
        Guess string `json:"guess"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Call your Wordle solving function (implement this)
    result := solve(input.Guess)

    // Respond with the result
    c.JSON(http.StatusOK, gin.H{"result": result})
}

// Placeholder for your Wordle solving logic
func solve(guess string) string {
    // Implement your logic here
    return "solution"
}
