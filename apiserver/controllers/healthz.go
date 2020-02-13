package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// [get] healthz
func HealthzController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "ok",
	})
}