package handlers

import (
	"net/http"

	"github.com/JackieTanakit/agnosbackend/go/internal/repositories"
	"github.com/JackieTanakit/agnosbackend/go/internal/services"

	"github.com/gin-gonic/gin"
)

func StrongPasswordStepsHandler(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	password := req["init_password"]
	numOfSteps := services.CalculateStrongPasswordSteps(password)

	// Store log
	repositories.StoreLog(req, gin.H{"num_of_steps": numOfSteps})

	c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}