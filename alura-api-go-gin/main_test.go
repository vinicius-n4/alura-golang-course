package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/controllers"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/database"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/models"
)

var ID int

func SetupTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CreateMockStudent() {
	s := models.Student{
		Name: "Student Test",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	database.DB.Create(&s)
	ID = int(s.ID)
}

func DeleteMockStudent() {
	var s models.Student
	database.DB.Delete(&s, ID)
}

func TestGreeting(t *testing.T) {
	r := SetupTest()
	r.GET("/:name", controllers.Greeting)

	req, _ := http.NewRequest("GET", "/studentName", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

	mockResponseBody := `{"API says":"Hello, studentName, how are you doing?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockResponseBody, string(responseBody))
}

func TestShowAllStudents(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTest()
	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByCpf(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTest()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByID(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTest()
	r.GET("/students/:id", controllers.GetStudentByID)

	req, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Student Test", studentMock.Name)
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()

	r := SetupTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := SetupTest()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	studentMock := models.Student{Name: "Student Test 2", CPF: "09876543210", RG: "098765432"}
	jsonBody, _ := json.Marshal(studentMock)

	req, _ := http.NewRequest("PATCH", "/students/"+strconv.Itoa(ID), bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var updatedStudentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedStudentMock)

	assert.Equal(t, "Student Test 2", updatedStudentMock.Name)
	assert.Equal(t, "09876543210", updatedStudentMock.CPF)
	assert.Equal(t, "098765432", updatedStudentMock.RG)
}
