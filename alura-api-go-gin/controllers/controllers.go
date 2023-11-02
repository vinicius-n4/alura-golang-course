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

	err = models.ValidateStudentData(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&s)
	c.JSON(http.StatusOK, s)
}

func GetStudentByID(c *gin.Context) {
	var s models.Student
	id := c.Params.ByName("id")
	database.DB.First(&s, id)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student with ID " + id + " was not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}

func DeleteStudent(c *gin.Context) {
	var s models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&s, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Student successfuly deleted"})
}

func UpdateStudent(c *gin.Context) {
	var s models.Student
	id := c.Params.ByName("id")
	database.DB.First(&s, id)

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := models.ValidateStudentData(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Save(&s)
	c.JSON(http.StatusOK, s)
}

func GetStudentByCpf(c *gin.Context) {
	var s models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&s)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student was not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}
