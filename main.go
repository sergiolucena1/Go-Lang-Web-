package main

import (
	"html/template"
	"log"
	"net/http"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))
//template.Must (encapsula todas as nossas templates e devolve 2 retornos(o template e o erro))
//template.ParseGlob (Passa o caminho dos templates)"nome da pata /* tudo que for html"

func main() {
	http.HandleFunc("/", index)// toda vez que tiver uma "/" a função index vai atender
	log.Fatal(http.ListenAndServe(":3000", nil)) // porta do servidor
}

	//Parametros(quem vai escrever, quem vai exibir )
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}