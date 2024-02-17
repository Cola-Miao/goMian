package service

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func Message(c *gin.Context, message string, data ...any) {
	c.JSON(http.StatusOK, gin.H{
		"mess": message,
		"data": data,
	})
}

func ErrorMessage(c *gin.Context, err error, desc string, rcv bool, data ...any) {
	var errMess string
	if desc == "" {
		errMess = err.Error()
	} else {
		errMess = desc
	}
	if rcv {
		slog.Error(err.Error(), "ip", c.ClientIP(), "url", c.Request.URL.Path)
	}
	c.JSON(http.StatusOK, gin.H{
		"error": errMess,
		"data":  data,
	})
	c.Abort()
}
