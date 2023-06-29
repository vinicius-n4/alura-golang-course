package routes

import (
	"github.com/vinicius-n4/alura-golang-course/alura-store/controllers"
	"net/http"
)

func CarregaRotas() {
	// HandleFunc() eu configuro a rota da url, nesse caso "/"
	// index é a função que eu crio pra chamar um template específico
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
