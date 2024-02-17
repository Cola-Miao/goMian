package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mess": "ok",
	})
}

func Private(c *gin.Context) {
	id := c.GetInt("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
