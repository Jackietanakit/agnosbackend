package main

import (
	"github.com/JackieTanakit/agnosbackend/go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handlers.IndexHandler)
	r.POST("/api/strong_password_steps", handlers.StrongPasswordStepsHandler)
	r.Run(":8080")
}
