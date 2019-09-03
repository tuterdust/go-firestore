package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingServiceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
