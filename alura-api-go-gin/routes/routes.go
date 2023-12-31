package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/controllers"
	docs "github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/docs"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.EndpointNotFound)

	docs.SwaggerInfo.BasePath = "/"
	// To access Swagger on browser: http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
