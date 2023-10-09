package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
