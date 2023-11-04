package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/database"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/models"
)

// ShowAllStudents godoc
//
//	@Summary		Show all students
//	@Description	get all students recorded on database
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} models.Student
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students [get]
func ShowAllStudents(c *gin.Context) {
	var s []models.Student
	database.DB.Find(&s)
	c.JSON(http.StatusOK, s)
}

// Greeting godoc
//
//	@Summary		Show a greeting
//	@Description	get a greeting for a given name
//	@Tags			greeting
//	@Accept			json
//	@Produce		json
//	@Param			name path string true "Student Name"
//	@Success		200	{object} string
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/{name} [get]
func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API says": "Hello, " + name + ", how are you doing?",
	})
}

// CreateStudent godoc
//
//	@Summary		Create a student
//	@Description	create a register for a student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			name body string true "Student Name"
//	@Param			cpf body int true "Student CPF"
//	@Param			rg body	 int true "Student RG"
//	@Success		200	{object} models.Student
//	@Failure		400	{object} httputil.HTTPError
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students [post]
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

// GetStudentByID godoc
//
//	@Summary		Show a student
//	@Description	get a student by ID
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id	path int true "Student ID"
//	@Success		200	{object} models.Student
//	@Failure		404	{object} httputil.HTTPError
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students/{id} [get]
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

// DeleteStudent godoc
//
//	@Summary		Delete a student
//	@Description	delete a student's register by ID
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id	path int true "Student ID"
//	@Success		200	{object} models.Student
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	var s models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&s, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Student successfuly deleted"})
}

// UpdateStudent godoc
//
//	@Summary		Update a student
//	@Description	update a student's register by ID
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id	path int true "Student ID"
//	@Success		200	{object} models.Student
//	@Failure		400	{object} httputil.HTTPError
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students/{id} [patch]
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

// GetStudentByCpf godoc
//
//	@Summary		Show a student
//	@Description	get a student by CPF
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			cpf	path int true "Student CPF"
//	@Success		200	{object} models.Student
//	@Failure		404	{object} httputil.HTTPError
//	@Failure		500	{object} httputil.HTTPError
//	@Router			/students/cpf/{cpf} [get]
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

func ShowIndexPage(c *gin.Context) {
	var s []models.Student
	database.DB.Find(&s)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": s,
	})
}

func EndpointNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
