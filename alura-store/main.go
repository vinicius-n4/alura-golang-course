package main

import (
	"alura-golang-course/alura-store/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	// método para subir um servidor na porta 8000 de localhost
	http.ListenAndServe(":8000", nil)
}
