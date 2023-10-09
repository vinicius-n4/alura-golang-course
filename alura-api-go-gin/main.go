package main

import (
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/database"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
