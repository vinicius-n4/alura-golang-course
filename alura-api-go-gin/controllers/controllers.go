package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/models"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API says":"Hello, " + name + ", how are you doing?",
	})
}