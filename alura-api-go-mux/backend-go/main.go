package main

import (
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-mux/backend-go/database"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-mux/backend-go/models"
	"github.com/vinicius-n4/alura-golang-course/alura-api-go-mux/backend-go/routes"
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
