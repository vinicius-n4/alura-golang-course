package main

import (
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/models"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-gin/routes"
)

func main() {
	models.Students = []models.Student{
	{Name: "Vinicius", CPF: "00000000000", RG: "11111111-1"},
	{Name: "Costa", CPF: "22222222222", RG: "33333333-3"},
	}
	routes.HandleRequests()
}
