package controllers

import (
	"net/http"
	"text/template"

	"github.com/sergio/Go-Lang-Web-/models"
)

//variavel com todos os nossos templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

//template.Must (encapsula todas as nossas templates e devolve 2 retornos(o template e o erro))
//template.ParseGlob (Passa o caminho dos templates)"nome da pata /* tudo que for html"

//Parametros(quem vai escrever, quem vai exibir )
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	//EXECUTA NO TEMPLATE
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}
