package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func init() {
	fmt.Println("Hello world")
}

func trackIt(c *gin.Context) {
	fmt.Println("track it!")
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/login", trackIt)

	router.Run(":8080")
}
