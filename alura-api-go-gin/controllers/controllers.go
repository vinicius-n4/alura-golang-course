package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/database"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/models"
)

func ShowAllStudents(c *gin.Context) {
	var s []models.Student
	database.DB.Find(&s)
	c.JSON(http.StatusOK, s)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API says": "Hello, " + name + ", how are you doing?",
	})
}

func CreateStudent(c *gin.Context) {
	var s models.Student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&s)
	c.JSON(http.StatusOK, s)
}
