package main

import (
	"alura-golang-course/alura-api-go-rest/api-go-rest/database"
	"alura-golang-course/alura-api-go-rest/api-go-rest/models"
	"alura-golang-course/alura-api-go-rest/api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
